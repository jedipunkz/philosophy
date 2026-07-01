package scrapem

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/jedipunkz/philosophy/internal/config"
)

// Retry tuning for transient upstream failures (429/5xx). 403 is deliberately
// excluded: in practice it has meant a durable WAF/IP-reputation block on the
// CI runner's IP range rather than a transient condition, so retrying it just
// burns the run's time budget without changing the outcome.
const (
	maxRetryAttempts = 3
	baseRetryDelay   = 3 * time.Second
	maxRetryDelay    = 30 * time.Second
)

type sourceQuery struct {
	keyword config.KeywordConfig
	query   string
}

type Runner struct {
	cfg          config.Config
	client       *http.Client
	sources      map[string]config.SourceConfig
	requestDelay time.Duration
}

func New(cfg config.Config) *Runner {
	timeout, err := time.ParseDuration(cfg.Scrape.RequestTimeout)
	if err != nil {
		timeout = 20 * time.Second
	}
	requestDelay, err := time.ParseDuration(cfg.Scrape.RequestDelay)
	if err != nil {
		requestDelay = 2 * time.Second
	}
	sources := make(map[string]config.SourceConfig, len(cfg.Sources))
	for _, source := range cfg.Sources {
		sources[source.Name] = source
	}
	return &Runner{
		cfg: cfg,
		client: &http.Client{
			Timeout: timeout,
		},
		sources:      sources,
		requestDelay: requestDelay,
	}
}

func (r *Runner) Run(ctx context.Context) error {
	inbox := filepath.Join(r.cfg.Vault.Root, r.cfg.Vault.Inbox)
	if err := os.MkdirAll(inbox, 0o755); err != nil {
		return err
	}

	statePath := filepath.Join(r.cfg.Vault.Root, ".scrapem", r.cfg.Vault.SeenFile)
	seen, err := loadSeen(statePath)
	if err != nil {
		return err
	}
	added, err := syncSeenFromInbox(inbox, seen)
	if err != nil {
		return err
	}
	if added > 0 {
		log.Printf("seen index restored from inbox: +%d urls", added)
		if err := seen.Save(statePath); err != nil {
			return err
		}
	}

	jobs := map[string][]sourceQuery{}
	for _, keyword := range r.cfg.Keywords {
		for _, sourceName := range keyword.Sources {
			for _, query := range keyword.Queries {
				jobs[sourceName] = append(jobs[sourceName], sourceQuery{
					keyword: keyword,
					query:   query,
				})
			}
		}
	}

	var (
		wg    sync.WaitGroup
		mu    sync.Mutex
		total int
	)

	for sourceName, queries := range jobs {
		source, ok := r.sources[sourceName]
		if !ok {
			log.Printf("source %q is not defined; skipping", sourceName)
			continue
		}
		if !source.Enabled {
			log.Printf("source %q is disabled; skipping", sourceName)
			continue
		}

		wg.Add(1)
		go func(source config.SourceConfig, queries []sourceQuery) {
			defer wg.Done()
			r.runSource(ctx, source, queries, inbox, statePath, seen, &mu, &total)
		}(source, queries)
	}

	wg.Wait()

	if err := seen.Save(statePath); err != nil {
		log.Printf("seen save failed: %v", err)
	}
	log.Printf("scrape completed: wrote %d new notes", total)
	return nil
}

func (r *Runner) runSource(ctx context.Context, source config.SourceConfig, queries []sourceQuery, inbox, statePath string, seen *Seen, mu *sync.Mutex, total *int) {
	for _, sq := range queries {
		if ctx.Err() != nil {
			return
		}
		if r.maxNewItemsReached(mu, total) {
			return
		}
		items, err := r.search(ctx, source, sq.query)
		if err != nil {
			log.Printf("search failed source=%s query=%q: %v", source.Name, sq.query, err)
			continue
		}
		if len(items) > r.cfg.Scrape.MaxResults {
			items = items[:r.cfg.Scrape.MaxResults]
		}
		isBookSource := isBookSourceType(source.Type)
		for _, item := range items {
			if ctx.Err() != nil {
				return
			}
			if r.maxNewItemsReached(mu, total) {
				return
			}
			// Skip expensive fetches early when the item is already known and
			// refresh is disabled. We re-check under the lock after enrichment
			// to handle the case where a sibling goroutine just added the URL.
			mu.Lock()
			alreadySeen := seen.Has(item.URL)
			mu.Unlock()
			if alreadySeen && !r.cfg.Scrape.RefreshExisting {
				continue
			}

			switch source.Type {
			case "gutenberg_api":
				if err := r.enrichBookText(ctx, &item, source); err != nil {
					log.Printf("gutenberg text fetch failed url=%s text=%s: %v", item.URL, item.PlainTextURL, err)
				}
			case "archive_api":
				if err := r.enrichArchiveText(ctx, &item, source); err != nil {
					log.Printf("archive text fetch failed url=%s: %v", item.URL, err)
				}
			case "sep_html":
				if err := r.enrichSEP(ctx, &item, source); err != nil {
					log.Printf("sep entry fetch failed url=%s: %v", item.URL, err)
				}
			case "wikipedia_api":
				if err := r.enrichWikipediaText(ctx, &item, source); err != nil {
					log.Printf("wikipedia text fetch failed url=%s: %v", item.URL, err)
				}
			case "api", "crossref_api":
				// arxiv / crossref — search response already has full metadata, no detail fetch
			default:
				if err := r.enrich(ctx, &item, source); err != nil {
					log.Printf("detail fetch failed url=%s: %v", item.URL, err)
				}
			}
			if !isBookSource {
				if err := r.enrichPDF(ctx, &item, source); err != nil {
					log.Printf("pdf fetch failed url=%s pdf=%s: %v", item.URL, item.PDF, err)
				}
			}
			if item.BookText != "" {
				item.BookText, item.BookTextTruncated = truncateRunes(item.BookText, r.cfg.Scrape.MaxBookChars)
			}
			// Skip book items that yielded no extractable text: Wikipedia
			// disambiguation pages/stubs, and Internet Archive items that are
			// either access-restricted (lending-only) or have no OCR text
			// derivative. A book note with no body is not useful, and
			// archive.org's full-text search often matches irrelevant items
			// (court filings, magazines, unrelated works) that also tend to
			// lack usable text.
			if (source.Type == "wikipedia_api" || source.Type == "archive_api") && item.BookText == "" {
				log.Printf("skipping %s item with no text url=%s", source.Type, item.URL)
				continue
			}
			// Skip paper items (crossref / arXiv) that carry no usable body:
			// no abstract, no extractable PDF text, and no detail text. These
			// otherwise produce inbox notes with only bibliographic metadata
			// and a "要旨は含まれていない" placeholder, which are not worth
			// reviewing.
			if !isBookSource &&
				strings.TrimSpace(item.Abstract) == "" &&
				strings.TrimSpace(item.PDFText) == "" &&
				strings.TrimSpace(item.DetailText) == "" {
				log.Printf("skipping %s item with no abstract, pdf, or detail text url=%s", source.Type, item.URL)
				continue
			}

			item.Keyword = sq.keyword.Name
			item.Query = sq.query
			item.Tags = sq.keyword.Tags
			item.SourceName = source.Name

			mu.Lock()
			if r.cfg.Scrape.MaxNewItems > 0 && *total >= r.cfg.Scrape.MaxNewItems {
				mu.Unlock()
				return
			}
			alreadySeen = seen.Has(item.URL)
			if alreadySeen && !r.cfg.Scrape.RefreshExisting {
				mu.Unlock()
				continue
			}
			var werr error
			if alreadySeen {
				werr = updateMarkdownBySource(inbox, item)
			} else {
				_, werr = writeMarkdown(inbox, item)
			}
			if werr != nil {
				mu.Unlock()
				log.Printf("write failed url=%s: %v", item.URL, werr)
				continue
			}
			if !alreadySeen {
				seen.Add(item.URL)
			}
			*total++
			if err := seen.Save(statePath); err != nil {
				log.Printf("seen save failed: %v", err)
			}
			mu.Unlock()
		}
	}
}

func (r *Runner) maxNewItemsReached(mu *sync.Mutex, total *int) bool {
	if r.cfg.Scrape.MaxNewItems <= 0 {
		return false
	}
	mu.Lock()
	defer mu.Unlock()
	return *total >= r.cfg.Scrape.MaxNewItems
}

// userAgentFor returns the source's user_agent override when set, falling
// back to the global scrape.user_agent otherwise.
func (r *Runner) userAgentFor(source config.SourceConfig) string {
	if source.UserAgent != "" {
		return source.UserAgent
	}
	return r.cfg.Scrape.UserAgent
}

func (r *Runner) search(ctx context.Context, source config.SourceConfig, query string) ([]Item, error) {
	if err := r.wait(ctx); err != nil {
		return nil, err
	}
	searchURL := strings.ReplaceAll(source.SearchURL, "{query}", url.QueryEscape(query))
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, searchURL, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", r.userAgentFor(source))
	switch source.Type {
	case "api":
		req.Header.Set("Accept", "application/atom+xml,application/xml;q=0.9,*/*;q=0.8")
	case "archive_api", "crossref_api":
		req.Header.Set("Accept", "application/json,*/*;q=0.8")
	default:
		req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
	}

	resp, err := r.doRequest(ctx, req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	switch source.Type {
	case "api":
		return parseArxivFeed(resp.Body)
	case "crossref_api":
		return parseCrossrefResults(resp.Body)
	case "archive_api":
		return parseArchiveResults(resp.Body, source.BaseURL)
	case "gutenberg_api":
		return parseGutenbergResults(resp.Body)
	case "wikipedia_api":
		return parseWikipediaResults(resp.Body, source.BaseURL)
	case "sep_html":
		return parseSEPSearchResults(resp.Body, source.BaseURL)
	default:
		return parseSearchResults(resp.Body, resp.Request.URL)
	}
}

func isBookSourceType(t string) bool {
	switch t {
	case "gutenberg_api", "archive_api", "wikipedia_api", "sep_html":
		return true
	}
	return false
}

func (r *Runner) enrich(ctx context.Context, item *Item, source config.SourceConfig) error {
	if err := r.wait(ctx); err != nil {
		return err
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, item.URL, nil)
	if err != nil {
		return err
	}
	req.Header.Set("User-Agent", r.userAgentFor(source))
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")

	resp, err := r.doRequest(ctx, req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	detail, err := parseDetail(resp.Body, resp.Request.URL)
	if err != nil {
		return err
	}
	mergeDetail(item, detail)
	return nil
}

func (r *Runner) wait(ctx context.Context) error {
	return r.sleepFor(ctx, r.requestDelay)
}

func (r *Runner) sleepFor(ctx context.Context, d time.Duration) error {
	if d <= 0 {
		return nil
	}
	timer := time.NewTimer(d)
	defer timer.Stop()
	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-timer.C:
		return nil
	}
}

// doRequest executes req, retrying on 429 and 5xx responses (honoring a
// Retry-After header when the upstream sends one) up to maxRetryAttempts
// times. Any other non-2xx status, including 403, is returned immediately
// without retry. req.Body must be nil (true for every caller in this
// package, which only issue GET requests) since a retried request is reused
// as-is.
func (r *Runner) doRequest(ctx context.Context, req *http.Request) (*http.Response, error) {
	for attempt := 0; ; attempt++ {
		resp, err := r.client.Do(req)
		if err != nil {
			return nil, err
		}
		if resp.StatusCode >= 200 && resp.StatusCode < 300 {
			return resp, nil
		}
		if !isRetryableStatus(resp.StatusCode) || attempt >= maxRetryAttempts {
			status := resp.Status
			resp.Body.Close()
			return nil, fmt.Errorf("unexpected status %s", status)
		}
		delay := retryDelay(resp, attempt)
		status := resp.Status
		resp.Body.Close()
		log.Printf("retrying request url=%s status=%s attempt=%d/%d delay=%s", req.URL, status, attempt+1, maxRetryAttempts, delay)
		if err := r.sleepFor(ctx, delay); err != nil {
			return nil, err
		}
	}
}

func isRetryableStatus(code int) bool {
	return code == http.StatusTooManyRequests || (code >= 500 && code < 600)
}

// retryDelay honors the Retry-After response header (either delta-seconds or
// an HTTP-date) when present, falling back to exponential backoff from
// baseRetryDelay. The result is capped at maxRetryDelay so a misbehaving
// upstream can't stall the run past its --max-duration budget.
func retryDelay(resp *http.Response, attempt int) time.Duration {
	if ra := strings.TrimSpace(resp.Header.Get("Retry-After")); ra != "" {
		if secs, err := strconv.Atoi(ra); err == nil && secs >= 0 {
			return capDelay(time.Duration(secs) * time.Second)
		}
		if when, err := http.ParseTime(ra); err == nil {
			if d := time.Until(when); d > 0 {
				return capDelay(d)
			}
		}
	}
	return capDelay(baseRetryDelay * time.Duration(1<<attempt))
}

func capDelay(d time.Duration) time.Duration {
	if d > maxRetryDelay {
		return maxRetryDelay
	}
	return d
}

func mergeDetail(item *Item, detail Item) {
	if detail.Title != "" {
		item.Title = detail.Title
	}
	if detail.Author != "" {
		item.Author = detail.Author
	}
	if detail.Year != "" {
		item.Year = detail.Year
	}
	if detail.Info != "" {
		item.Info = detail.Info
	}
	if detail.Abstract != "" {
		item.Abstract = detail.Abstract
	}
	if detail.Download != "" {
		item.Download = detail.Download
	}
	if detail.PDF != "" {
		item.PDF = detail.PDF
	}
	if detail.Citation != "" {
		item.Citation = detail.Citation
	}
	if detail.BibTeX != "" {
		item.BibTeX = detail.BibTeX
	}
	if detail.DetailText != "" {
		item.DetailText = detail.DetailText
	}
	if detail.PDFText != "" {
		item.PDFText = detail.PDFText
		item.PDFTextTruncated = detail.PDFTextTruncated
	}
	if detail.BookText != "" {
		item.BookText = detail.BookText
		item.BookTextTruncated = detail.BookTextTruncated
	}
}
