package scrapem

import (
	"io"
	"regexp"
	"strings"
)

var sepEntryHrefRe = regexp.MustCompile(`(?i)href=['"](?:https?://plato\.stanford\.edu)?/entries/([a-z0-9][a-z0-9\-]*)/?['"#]`)

// parseSEPSearchResults extracts /entries/<slug>/ URLs from a SEP search HTML
// page. The body of each entry is fetched in enrichSEP.
func parseSEPSearchResults(body io.Reader, baseURL string) ([]Item, error) {
	data, err := io.ReadAll(body)
	if err != nil {
		return nil, err
	}
	text := string(data)
	matches := sepEntryHrefRe.FindAllStringSubmatch(text, -1)
	seen := map[string]bool{}
	base := strings.TrimRight(baseURL, "/")
	items := make([]Item, 0, len(matches))
	for _, m := range matches {
		slug := m[1]
		if seen[slug] {
			continue
		}
		seen[slug] = true
		items = append(items, Item{
			ItemType: "book",
			Title:    slug, // replaced by extracted <h1> in enrichSEP
			URL:      base + "/entries/" + slug + "/",
		})
	}
	return items, nil
}
