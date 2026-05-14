package scrapem

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"
)

type openLibraryResponse struct {
	Docs []openLibraryDoc `json:"docs"`
}

type openLibraryDoc struct {
	Key              string   `json:"key"`
	Title            string   `json:"title"`
	AuthorName       []string `json:"author_name"`
	FirstPublishYear int      `json:"first_publish_year"`
	Subject          []string `json:"subject"`
}

func parseOpenLibraryResults(body io.Reader, baseURL string) ([]Item, error) {
	data, err := io.ReadAll(body)
	if err != nil {
		return nil, err
	}
	var resp openLibraryResponse
	if err := json.Unmarshal(data, &resp); err != nil {
		return nil, err
	}
	items := make([]Item, 0, len(resp.Docs))
	for _, doc := range resp.Docs {
		if doc.Title == "" || doc.Key == "" {
			continue
		}
		year := ""
		if doc.FirstPublishYear > 0 {
			year = fmt.Sprintf("%d", doc.FirstPublishYear)
		}
		pageURL := strings.TrimRight(baseURL, "/") + doc.Key
		items = append(items, Item{
			ItemType: "book",
			Title:    doc.Title,
			URL:      pageURL,
			Author:   strings.Join(doc.AuthorName, ", "),
			Year:     year,
			Subjects: doc.Subject,
		})
	}
	return items, nil
}
