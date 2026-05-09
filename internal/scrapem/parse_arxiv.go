package scrapem

import (
	"encoding/xml"
	"io"
	"strings"
)

type arxivFeed struct {
	XMLName xml.Name     `xml:"feed"`
	Entries []arxivEntry `xml:"entry"`
}

type arxivEntry struct {
	ID        string       `xml:"id"`
	Title     string       `xml:"title"`
	Summary   string       `xml:"summary"`
	Published string       `xml:"published"`
	Authors   []arxivName  `xml:"author"`
	Links     []arxivLink  `xml:"link"`
	Category  arxivCategry `xml:"primary_category"`
}

type arxivName struct {
	Name string `xml:"name"`
}

type arxivLink struct {
	Href  string `xml:"href,attr"`
	Rel   string `xml:"rel,attr"`
	Type  string `xml:"type,attr"`
	Title string `xml:"title,attr"`
}

type arxivCategry struct {
	Term string `xml:"term,attr"`
}

func parseArxivFeed(body io.Reader) ([]Item, error) {
	var feed arxivFeed
	if err := xml.NewDecoder(body).Decode(&feed); err != nil {
		return nil, err
	}
	items := make([]Item, 0, len(feed.Entries))
	for _, entry := range feed.Entries {
		title := normalizeArxivText(entry.Title)
		if title == "" {
			continue
		}
		pageURL, pdfURL := extractArxivLinks(entry)
		if pageURL == "" {
			pageURL = strings.TrimSpace(entry.ID)
		}
		if pageURL == "" {
			continue
		}
		item := Item{
			Title:    title,
			URL:      pageURL,
			Author:   joinArxivAuthors(entry.Authors),
			Year:     arxivYear(entry.Published),
			Info:     arxivInfo(entry),
			Abstract: normalizeArxivText(entry.Summary),
			PDF:      pdfURL,
			Download: pdfURL,
		}
		items = append(items, item)
	}
	return items, nil
}

func extractArxivLinks(entry arxivEntry) (string, string) {
	var pageURL, pdfURL string
	for _, link := range entry.Links {
		switch {
		case link.Rel == "alternate" && (link.Type == "text/html" || link.Type == ""):
			if pageURL == "" {
				pageURL = strings.TrimSpace(link.Href)
			}
		case link.Type == "application/pdf" || link.Title == "pdf":
			if pdfURL == "" {
				pdfURL = strings.TrimSpace(link.Href)
			}
		}
	}
	return pageURL, pdfURL
}

func joinArxivAuthors(authors []arxivName) string {
	names := make([]string, 0, len(authors))
	for _, a := range authors {
		name := strings.TrimSpace(a.Name)
		if name != "" {
			names = append(names, name)
		}
	}
	return strings.Join(names, ", ")
}

func arxivYear(published string) string {
	published = strings.TrimSpace(published)
	if len(published) >= 4 {
		return published[:4]
	}
	return published
}

func arxivInfo(entry arxivEntry) string {
	parts := []string{"arXiv preprint"}
	if cat := strings.TrimSpace(entry.Category.Term); cat != "" {
		parts = append(parts, cat)
	}
	return strings.Join(parts, " / ")
}

func normalizeArxivText(input string) string {
	return strings.TrimSpace(spaceRe.ReplaceAllString(input, " "))
}
