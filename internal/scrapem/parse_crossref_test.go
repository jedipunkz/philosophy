package scrapem

import (
	"strings"
	"testing"
)

func TestParseCrossrefResults(t *testing.T) {
	body := `{
		"message": {
			"items": [
				{
					"DOI": "10.1000/xyz123",
					"URL": "https://doi.org/10.1000/xyz123",
					"title": ["On the Forms"],
					"author": [{"given": "Ada", "family": "Example"}, {"given": "Bea", "family": "Sample"}],
					"container-title": ["Journal of Philosophy Examples"],
					"published": {"date-parts": [[2019, 3]]},
					"abstract": "<jats:p>This paper discusses <jats:italic>Plato</jats:italic>.</jats:p>"
				},
				{
					"DOI": "",
					"title": ["Should be skipped: no DOI"]
				}
			]
		}
	}`

	items, err := parseCrossrefResults(strings.NewReader(body))
	if err != nil {
		t.Fatalf("parseCrossrefResults returned error: %v", err)
	}
	if len(items) != 1 {
		t.Fatalf("got %d items, want 1 (item without DOI should be skipped): %+v", len(items), items)
	}
	item := items[0]
	if item.Title != "On the Forms" {
		t.Errorf("Title = %q", item.Title)
	}
	if item.Author != "Ada Example, Bea Sample" {
		t.Errorf("Author = %q", item.Author)
	}
	if item.Year != "2019" {
		t.Errorf("Year = %q", item.Year)
	}
	if item.URL != "https://doi.org/10.1000/xyz123" {
		t.Errorf("URL = %q", item.URL)
	}
	if item.Info != "Journal of Philosophy Examples" {
		t.Errorf("Info = %q", item.Info)
	}
	if item.Abstract != "This paper discusses Plato." {
		t.Errorf("Abstract = %q", item.Abstract)
	}
	if item.Citation != "DOI: 10.1000/xyz123" {
		t.Errorf("Citation = %q", item.Citation)
	}
}

func TestParseCrossrefResultsFallsBackToDOIURL(t *testing.T) {
	body := `{"message": {"items": [{"DOI": "10.1000/abc", "title": ["No URL Field"]}]}}`
	items, err := parseCrossrefResults(strings.NewReader(body))
	if err != nil {
		t.Fatalf("parseCrossrefResults returned error: %v", err)
	}
	if len(items) != 1 {
		t.Fatalf("got %d items, want 1", len(items))
	}
	if items[0].URL != "https://doi.org/10.1000/abc" {
		t.Errorf("URL = %q, want DOI-derived fallback", items[0].URL)
	}
}
