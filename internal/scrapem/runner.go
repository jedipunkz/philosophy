package scrapem

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/jedipunkz/philosophy/internal/config"
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
			r.runSource(ctx, source, queries, inbox, seen, &mu, &total)
		}(source, queries)
	}

	wg.Wait()

	if err := seen.Save(statePath); err != nil {
		return err
	}
	log.Printf("scrape completed: wrote %d new notes", total)
	return nil
}

func (r *Runner) runSource(ctx context.Context, source config.SourceConfig, queries []sourceQuery, inbox string, seen *Seen, mu *sync.Mutex, total *int) {
	for _, sq := range queries {
		items, err := r.search(ctx, source, sq.query)
		if err != nil {
			log.Printf("search failed source=%s query=%q: %v", source.Name, sq.query, err)
			continue
		}
		if len(items) > r.cfg.Scrape.MaxResults {
			items = items[:r.cfg.Scrape.MaxResults]
		}
		isBookSource := source.Type == "openlibrary_api" || source.Type == "gutenberg_api"
		for _, item := range items {
			switch source.Type {
			case "openlibrary_api":
				if err := r.enrichOpenLibrary(ctx, &item); err != nil {
					log.Printf("openlibrary detail fetch failed url=%s: %v", item.URL, err)
				}
			case "gutenberg_api":
				if err := r.enrichBookText(ctx, &item); err != nil {
					log.Printf("gutenberg text fetch failed url=%s text=%s: %v", item.URL, item.PlainTextURL, err)
				}
			case "api":
				// arxiv etc — handled below
			default:
				if err := r.enrich(ctx, &item); err != nil {
					log.Printf("detail fetch failed url=%s: %v", item.URL, err)
				}
			}
			if !isBookSource {
				if err := r.enrichPDF(ctx, &item); err != nil {
					log.Printf("pdf fetch failed url=%s pdf=%s: %v", item.URL, item.PDF, err)
				}
			}

			item.Keyword = sq.keyword.Name
			item.Query = sq.query
			item.Tags = sq.keyword.Tags
			item.SourceName = source.Name

			mu.Lock()
			alreadySeen := seen.Has(item.URL)
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
			mu.Unlock()
		}
	}
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
	req.Header.Set("User-Agent", r.cfg.Scrape.UserAgent)
	if source.Type == "api" {
		req.Header.Set("Accept", "application/atom+xml,application/xml;q=0.9,*/*;q=0.8")
	} else {
		req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
	}

	resp, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("unexpected status %s", resp.Status)
	}

	switch source.Type {
	case "api":
		return parseArxivFeed(resp.Body)
	case "openlibrary_api":
		return parseOpenLibraryResults(resp.Body, source.BaseURL)
	case "gutenberg_api":
		return parseGutenbergResults(resp.Body)
	default:
		return parseSearchResults(resp.Body, resp.Request.URL)
	}
}

func (r *Runner) enrich(ctx context.Context, item *Item) error {
	if err := r.wait(ctx); err != nil {
		return err
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, item.URL, nil)
	if err != nil {
		return err
	}
	req.Header.Set("User-Agent", r.cfg.Scrape.UserAgent)
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")

	resp, err := r.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Errorf("unexpected status %s", resp.Status)
	}

	detail, err := parseDetail(resp.Body, resp.Request.URL)
	if err != nil {
		return err
	}
	mergeDetail(item, detail)
	return nil
}

func (r *Runner) wait(ctx context.Context) error {
	if r.requestDelay <= 0 {
		return nil
	}
	timer := time.NewTimer(r.requestDelay)
	defer timer.Stop()
	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-timer.C:
		return nil
	}
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
	if detail.Description != "" {
		item.Description = detail.Description
	}
	if detail.BookText != "" {
		item.BookText = detail.BookText
		item.BookTextTruncated = detail.BookTextTruncated
	}
}
