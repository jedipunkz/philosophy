package scrapem

import (
	"encoding/json"
	"fmt"
	"io"
	"net/url"
	"strings"
)

type wikipediaResponse struct {
	Query *struct {
		Pages []wikipediaPage `json:"pages"`
	} `json:"query"`
}

type wikipediaPage struct {
	PageID  int    `json:"pageid"`
	Title   string `json:"title"`
	Extract string `json:"extract"`
	Index   int    `json:"index"`
}

// parseWikipediaResults parses the MediaWiki action=query response that combines
// list=search (via generator=search) with prop=extracts&explaintext=1. The
// resulting Items already contain the page body in BookText; no per-page enrich
// fetch is needed afterwards.
func parseWikipediaResults(body io.Reader, baseURL string) ([]Item, error) {
	data, err := io.ReadAll(body)
	if err != nil {
		return nil, err
	}
	var resp wikipediaResponse
	if err := json.Unmarshal(data, &resp); err != nil {
		return nil, err
	}
	if resp.Query == nil {
		return nil, nil
	}
	items := make([]Item, 0, len(resp.Query.Pages))
	for _, page := range resp.Query.Pages {
		if page.Title == "" || page.PageID == 0 {
			continue
		}
		items = append(items, Item{
			ItemType: "book",
			Title:    page.Title,
			URL:      wikipediaPageURL(baseURL, page.Title),
			BookText: strings.TrimSpace(page.Extract),
		})
	}
	return items, nil
}

func wikipediaPageURL(baseURL, title string) string {
	encoded := strings.ReplaceAll(url.PathEscape(title), "%20", "_")
	return fmt.Sprintf("%s/wiki/%s", strings.TrimRight(baseURL, "/"), encoded)
}
