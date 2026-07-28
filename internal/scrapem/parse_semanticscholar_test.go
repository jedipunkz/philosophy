package scrapem

import (
	"strings"
	"testing"
)

func TestParseSemanticScholarResults(t *testing.T) {
	body := `{
		"total": 2,
		"offset": 0,
		"data": [
			{
				"paperId": "abc123",
				"externalIds": {"DOI": "10.1000/xyz", "CorpusId": 42},
				"url": "https://www.semanticscholar.org/paper/abc123",
				"title": "On Michel Foucault: Power and Discourse",
				"venue": "Journal of Philosophy Examples",
				"year": 2023,
				"openAccessPdf": {"url": "https://example.org/paper.pdf", "status": "GOLD"},
				"authors": [{"authorId": "1", "name": "Ada Example"}, {"authorId": "2", "name": "Bea Sample"}],
				"abstract": "This paper examines power and discourse."
			},
			{
				"paperId": "no-title",
				"title": "",
				"year": 2020
			}
		]
	}`

	items, err := parseSemanticScholarResults(strings.NewReader(body))
	if err != nil {
		t.Fatalf("parseSemanticScholarResults returned error: %v", err)
	}
	if len(items) != 1 {
		t.Fatalf("got %d items, want 1 (empty-title item should be skipped): %+v", len(items), items)
	}
	item := items[0]
	if item.Title != "On Michel Foucault: Power and Discourse" {
		t.Errorf("Title = %q", item.Title)
	}
	if item.Author != "Ada Example, Bea Sample" {
		t.Errorf("Author = %q", item.Author)
	}
	if item.Year != "2023" {
		t.Errorf("Year = %q", item.Year)
	}
	if item.URL != "https://www.semanticscholar.org/paper/abc123" {
		t.Errorf("URL = %q", item.URL)
	}
	if item.Info != "Journal of Philosophy Examples" {
		t.Errorf("Info = %q", item.Info)
	}
	if item.Abstract != "This paper examines power and discourse." {
		t.Errorf("Abstract = %q", item.Abstract)
	}
	if item.Citation != "DOI: 10.1000/xyz" {
		t.Errorf("Citation = %q", item.Citation)
	}
	if item.PDF != "https://example.org/paper.pdf" {
		t.Errorf("PDF = %q", item.PDF)
	}
	if item.Download != "https://example.org/paper.pdf" {
		t.Errorf("Download = %q", item.Download)
	}
}

func TestParseSemanticScholarResultsFallbacks(t *testing.T) {
	// No url field -> derive from paperId; no DOI but arXiv -> arXiv citation;
	// year 0 -> empty; openAccessPdf null / empty url -> no PDF.
	body := `{
		"data": [
			{
				"paperId": "xyz789",
				"externalIds": {"ArXiv": "2101.00001"},
				"title": "A Preprint on Nihilism",
				"year": 0,
				"openAccessPdf": {"url": ""},
				"authors": []
			}
		]
	}`

	items, err := parseSemanticScholarResults(strings.NewReader(body))
	if err != nil {
		t.Fatalf("parseSemanticScholarResults returned error: %v", err)
	}
	if len(items) != 1 {
		t.Fatalf("got %d items, want 1", len(items))
	}
	item := items[0]
	if item.URL != "https://www.semanticscholar.org/paper/xyz789" {
		t.Errorf("URL = %q, want paperId-derived fallback", item.URL)
	}
	if item.Citation != "arXiv: 2101.00001" {
		t.Errorf("Citation = %q, want arXiv fallback", item.Citation)
	}
	if item.Year != "" {
		t.Errorf("Year = %q, want empty for year 0", item.Year)
	}
	if item.PDF != "" {
		t.Errorf("PDF = %q, want empty for blank openAccessPdf url", item.PDF)
	}
}
