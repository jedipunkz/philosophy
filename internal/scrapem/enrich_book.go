package scrapem

import (
	"context"
	"fmt"
	"html"
	"io"
	"net/http"
	"regexp"
	"strings"
)

var (
	gutenbergStartRe = regexp.MustCompile(`(?m)^\s*\*\*\*\s*START OF (?:THIS|THE) PROJECT GUTENBERG EBOOK[^*]*\*\*\*\s*$`)
	gutenbergEndRe   = regexp.MustCompile(`(?m)^\s*\*\*\*\s*END OF (?:THIS|THE) PROJECT GUTENBERG EBOOK[^*]*\*\*\*\s*$`)
	sepTitleRe       = regexp.MustCompile(`(?is)<h1[^>]*>(.*?)</h1>`)
	sepBodyRe        = regexp.MustCompile(`(?is)<h1[^>]*>.*?</h1>(.*?)(?:<div\s+id=['"]bibliography['"]|<h2[^>]*>\s*Bibliography\s*</h2>)`)
	htmlBlockTagRe   = regexp.MustCompile(`(?i)</?(p|div|h[1-6]|li|ul|ol|blockquote|tr|td|th|pre|section|article|nav|aside|header|footer|table)[^>]*>`)
	htmlBrRe         = regexp.MustCompile(`(?i)<br\s*/?>`)
	htmlHrRe         = regexp.MustCompile(`(?i)<hr\s*/?>`)
	htmlScriptRe     = regexp.MustCompile(`(?is)<script[^>]*>.*?</script>`)
	htmlStyleRe      = regexp.MustCompile(`(?is)<style[^>]*>.*?</style>`)
	htmlInlineWSRe   = regexp.MustCompile(`[ \t]+`)
	htmlMultiLineRe  = regexp.MustCompile(`\n{3,}`)
)

// enrichBookText fetches the plain-text body of a Project Gutenberg book and
// stores it on the item. Truncation by MaxBookChars is applied later in the
// runner so it is uniform across sources.
func (r *Runner) enrichBookText(ctx context.Context, item *Item) error {
	if item.PlainTextURL == "" {
		return nil
	}
	body, err := r.fetchBodyWithLimit(ctx, item.PlainTextURL, "text/plain,*/*;q=0.8")
	if err != nil {
		return err
	}
	item.BookText = cleanGutenbergText(string(body))
	return nil
}

// enrichSEP fetches an SEP entry page, sets the title from <h1>, and stores the
// main body text on the item.
func (r *Runner) enrichSEP(ctx context.Context, item *Item) error {
	if item.URL == "" {
		return nil
	}
	body, err := r.fetchBodyWithLimit(ctx, item.URL, "text/html,*/*;q=0.8")
	if err != nil {
		return err
	}
	htmlText := string(body)
	if t := extractSEPTitle(htmlText); t != "" {
		item.Title = t
	}
	item.BookText = extractSEPBody(htmlText)
	return nil
}

// fetchBodyWithLimit performs a GET with a body-size cap derived from
// MaxBookBytes. The cap is enforced both by Content-Length (when present) and
// by a LimitReader on the body itself.
func (r *Runner) fetchBodyWithLimit(ctx context.Context, target, accept string) ([]byte, error) {
	if err := r.wait(ctx); err != nil {
		return nil, err
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, target, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", r.cfg.Scrape.UserAgent)
	req.Header.Set("Accept", accept)

	resp, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("unexpected status %s", resp.Status)
	}
	if r.cfg.Scrape.MaxBookBytes > 0 && resp.ContentLength > r.cfg.Scrape.MaxBookBytes {
		return nil, fmt.Errorf("body too large: %d bytes", resp.ContentLength)
	}
	limit := r.cfg.Scrape.MaxBookBytes + 1
	body, err := io.ReadAll(io.LimitReader(resp.Body, limit))
	if err != nil {
		return nil, err
	}
	if r.cfg.Scrape.MaxBookBytes > 0 && int64(len(body)) > r.cfg.Scrape.MaxBookBytes {
		return nil, fmt.Errorf("body too large: exceeded %d bytes", r.cfg.Scrape.MaxBookBytes)
	}
	return body, nil
}

// cleanGutenbergText strips the Project Gutenberg license header/footer and
// normalizes whitespace.
func cleanGutenbergText(input string) string {
	input = strings.ReplaceAll(input, "\r\n", "\n")
	if loc := gutenbergStartRe.FindStringIndex(input); loc != nil {
		input = input[loc[1]:]
	}
	if loc := gutenbergEndRe.FindStringIndex(input); loc != nil {
		input = input[:loc[0]]
	}
	input = pdfHyphenBreakRe.ReplaceAllString(input, "$1$2")
	lines := strings.Split(input, "\n")
	out := make([]string, 0, len(lines))
	blank := false
	for _, line := range lines {
		line = strings.TrimRight(line, " \t")
		if line == "" {
			if !blank {
				out = append(out, "")
			}
			blank = true
			continue
		}
		out = append(out, line)
		blank = false
	}
	return strings.TrimSpace(strings.Join(out, "\n"))
}

func extractSEPTitle(htmlText string) string {
	if m := sepTitleRe.FindStringSubmatch(htmlText); len(m) >= 2 {
		return cleanHTML(m[1])
	}
	return ""
}

func extractSEPBody(htmlText string) string {
	if m := sepBodyRe.FindStringSubmatch(htmlText); len(m) >= 2 {
		return htmlToText(m[1])
	}
	return ""
}

// htmlToText converts an HTML fragment to plain text while preserving paragraph
// boundaries. Block-level tags become double newlines; <br> becomes a single
// newline; other tags are stripped.
func htmlToText(input string) string {
	input = htmlScriptRe.ReplaceAllString(input, "")
	input = htmlStyleRe.ReplaceAllString(input, "")
	input = htmlBrRe.ReplaceAllString(input, "\n")
	input = htmlHrRe.ReplaceAllString(input, "\n\n")
	input = htmlBlockTagRe.ReplaceAllString(input, "\n\n")
	input = tagRe.ReplaceAllString(input, "")
	input = html.UnescapeString(input)

	lines := strings.Split(input, "\n")
	for i, line := range lines {
		line = htmlInlineWSRe.ReplaceAllString(line, " ")
		lines[i] = strings.TrimSpace(line)
	}
	text := strings.Join(lines, "\n")
	text = htmlMultiLineRe.ReplaceAllString(text, "\n\n")
	return strings.TrimSpace(text)
}
