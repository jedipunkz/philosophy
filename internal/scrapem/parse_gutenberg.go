package scrapem

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"
)

type gutenbergResponse struct {
	Results []gutenbergBook `json:"results"`
}

type gutenbergBook struct {
	ID       int               `json:"id"`
	Title    string            `json:"title"`
	Authors  []gutenbergAuthor `json:"authors"`
	Subjects []string          `json:"subjects"`
	Formats  map[string]string `json:"formats"`
	Copyright bool             `json:"copyright"`
}

type gutenbergAuthor struct {
	Name      string `json:"name"`
	BirthYear *int   `json:"birth_year"`
	DeathYear *int   `json:"death_year"`
}

func parseGutenbergResults(body io.Reader) ([]Item, error) {
	data, err := io.ReadAll(body)
	if err != nil {
		return nil, err
	}
	var resp gutenbergResponse
	if err := json.Unmarshal(data, &resp); err != nil {
		return nil, err
	}
	items := make([]Item, 0, len(resp.Results))
	for _, book := range resp.Results {
		if book.Title == "" || book.ID == 0 {
			continue
		}
		pageURL := fmt.Sprintf("https://www.gutenberg.org/ebooks/%d", book.ID)
		download := gutenbergDownloadURL(book.Formats)
		plainText := gutenbergPlainTextURL(book.Formats)
		authors := make([]string, 0, len(book.Authors))
		for _, a := range book.Authors {
			if name := gutenbergAuthorName(a.Name); name != "" {
				authors = append(authors, name)
			}
		}
		year := ""
		if len(book.Authors) > 0 && book.Authors[0].DeathYear != nil {
			year = fmt.Sprintf("%d", *book.Authors[0].DeathYear)
		}
		items = append(items, Item{
			ItemType:     "book",
			Title:        book.Title,
			URL:          pageURL,
			GutenbergURL: pageURL,
			Author:       strings.Join(authors, ", "),
			Year:         year,
			Download:     download,
			PlainTextURL: plainText,
			Subjects:     book.Subjects,
			PublicDomain: !book.Copyright,
		})
	}
	return items, nil
}

// gutenbergDownloadURL picks a human-readable download URL from the formats map,
// preferring HTML over epub over plain text.
func gutenbergDownloadURL(formats map[string]string) string {
	for _, mime := range []string{"text/html", "application/epub+zip", "text/plain"} {
		for k, v := range formats {
			if strings.HasPrefix(k, mime) {
				return v
			}
		}
	}
	return ""
}

// gutenbergPlainTextURL picks the plain text URL for full body extraction,
// preferring UTF-8 over us-ascii.
func gutenbergPlainTextURL(formats map[string]string) string {
	for k, v := range formats {
		if strings.HasPrefix(k, "text/plain") && strings.Contains(k, "utf-8") {
			return v
		}
	}
	for k, v := range formats {
		if strings.HasPrefix(k, "text/plain") {
			return v
		}
	}
	return ""
}

// gutenbergAuthorName converts "Lastname, Firstname" to "Firstname Lastname".
func gutenbergAuthorName(raw string) string {
	parts := strings.SplitN(raw, ", ", 2)
	if len(parts) == 2 {
		return strings.TrimSpace(parts[1]) + " " + strings.TrimSpace(parts[0])
	}
	return strings.TrimSpace(raw)
}
