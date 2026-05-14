package scrapem

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strings"
)

var gutenbergStartRe = regexp.MustCompile(`(?m)^\s*\*\*\*\s*START OF (?:THIS|THE) PROJECT GUTENBERG EBOOK[^*]*\*\*\*\s*$`)
var gutenbergEndRe = regexp.MustCompile(`(?m)^\s*\*\*\*\s*END OF (?:THIS|THE) PROJECT GUTENBERG EBOOK[^*]*\*\*\*\s*$`)

// enrichBookText fetches the plain-text body of a Project Gutenberg book
// and stores a cleaned, length-limited version on the item.
func (r *Runner) enrichBookText(ctx context.Context, item *Item) error {
	if item.PlainTextURL == "" {
		return nil
	}
	if err := r.wait(ctx); err != nil {
		return err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, item.PlainTextURL, nil)
	if err != nil {
		return err
	}
	req.Header.Set("User-Agent", r.cfg.Scrape.UserAgent)
	req.Header.Set("Accept", "text/plain,*/*;q=0.8")

	resp, err := r.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Errorf("unexpected status %s", resp.Status)
	}
	if r.cfg.Scrape.MaxBookBytes > 0 && resp.ContentLength > r.cfg.Scrape.MaxBookBytes {
		return fmt.Errorf("book text too large: %d bytes", resp.ContentLength)
	}

	limit := r.cfg.Scrape.MaxBookBytes + 1
	body, err := io.ReadAll(io.LimitReader(resp.Body, limit))
	if err != nil {
		return err
	}
	if r.cfg.Scrape.MaxBookBytes > 0 && int64(len(body)) > r.cfg.Scrape.MaxBookBytes {
		return fmt.Errorf("book text too large: exceeded %d bytes", r.cfg.Scrape.MaxBookBytes)
	}

	cleaned := cleanGutenbergText(string(body))
	item.BookText, item.BookTextTruncated = truncateRunes(cleaned, r.cfg.Scrape.MaxBookChars)
	return nil
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

// enrichOpenLibrary fetches the work-level description from Open Library
// (the search endpoint does not include description).
func (r *Runner) enrichOpenLibrary(ctx context.Context, item *Item) error {
	if item.URL == "" {
		return nil
	}
	if err := r.wait(ctx); err != nil {
		return err
	}

	apiURL := item.URL + ".json"
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, apiURL, nil)
	if err != nil {
		return err
	}
	req.Header.Set("User-Agent", r.cfg.Scrape.UserAgent)
	req.Header.Set("Accept", "application/json")

	resp, err := r.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusNotFound {
		return nil
	}
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Errorf("unexpected status %s", resp.Status)
	}

	var work struct {
		Description json.RawMessage `json:"description"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&work); err != nil {
		return err
	}
	item.Description = parseOpenLibraryDescription(work.Description)
	return nil
}

// parseOpenLibraryDescription handles both string and {"value": "..."} forms.
func parseOpenLibraryDescription(raw json.RawMessage) string {
	if len(raw) == 0 {
		return ""
	}
	var asString string
	if err := json.Unmarshal(raw, &asString); err == nil {
		return strings.TrimSpace(asString)
	}
	var asObj struct {
		Value string `json:"value"`
	}
	if err := json.Unmarshal(raw, &asObj); err == nil {
		return strings.TrimSpace(asObj.Value)
	}
	return ""
}
