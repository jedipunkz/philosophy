package scrapem

import (
	"strings"
	"testing"
)

func TestParseArchiveResults(t *testing.T) {
	body := `{
		"response": {
			"docs": [
				{
					"identifier": "republicofplato00plat",
					"title": "The Republic of Plato",
					"creator": "Plato",
					"year": "1902",
					"access-restricted-item": false
				},
				{
					"identifier": "multiauthorbook",
					"title": "Essays on Ethics",
					"creator": ["First Author", "Second Author"],
					"year": ["1999"],
					"access-restricted-item": true
				},
				{
					"identifier": "",
					"title": "Missing identifier should be skipped"
				}
			]
		}
	}`

	items, err := parseArchiveResults(strings.NewReader(body), "https://archive.org")
	if err != nil {
		t.Fatalf("parseArchiveResults returned error: %v", err)
	}
	if len(items) != 2 {
		t.Fatalf("got %d items, want 2: %+v", len(items), items)
	}

	first := items[0]
	if first.Title != "The Republic of Plato" {
		t.Errorf("Title = %q", first.Title)
	}
	if first.URL != "https://archive.org/details/republicofplato00plat" {
		t.Errorf("URL = %q", first.URL)
	}
	if first.Author != "Plato" {
		t.Errorf("Author = %q", first.Author)
	}
	if first.Year != "1902" {
		t.Errorf("Year = %q", first.Year)
	}
	if !first.PublicDomain {
		t.Errorf("PublicDomain = false, want true for access-restricted-item=false")
	}
	if first.ItemType != "book" {
		t.Errorf("ItemType = %q, want book", first.ItemType)
	}

	second := items[1]
	if second.Author != "First Author, Second Author" {
		t.Errorf("Author = %q", second.Author)
	}
	if second.PublicDomain {
		t.Errorf("PublicDomain = true, want false for access-restricted-item=true")
	}
}

// TestParseArchiveResultsHandlesNumericYear reproduces a production failure:
// archive.org's advancedsearch.php sometimes sends "year" as a bare JSON
// number (e.g. 1902) rather than a string, which previously made the whole
// response fail to decode with "cannot unmarshal number into ... []string".
func TestParseArchiveResultsHandlesNumericYear(t *testing.T) {
	body := `{
		"response": {
			"docs": [
				{
					"identifier": "numericyearbook",
					"title": "A Book With Numeric Year",
					"creator": "Some Author",
					"year": 1902,
					"access-restricted-item": false
				},
				{
					"identifier": "numericyeararray",
					"title": "A Book With Numeric Year Array",
					"creator": "Another Author",
					"year": [1999, 2001],
					"access-restricted-item": false
				}
			]
		}
	}`

	items, err := parseArchiveResults(strings.NewReader(body), "https://archive.org")
	if err != nil {
		t.Fatalf("parseArchiveResults returned error: %v", err)
	}
	if len(items) != 2 {
		t.Fatalf("got %d items, want 2: %+v", len(items), items)
	}
	if items[0].Year != "1902" {
		t.Errorf("Year = %q, want 1902", items[0].Year)
	}
	if items[1].Year != "1999" {
		t.Errorf("Year = %q, want 1999 (first element of numeric array)", items[1].Year)
	}
}

func TestPickArchiveTextFilePrefersDjvuTextAndLargestFile(t *testing.T) {
	files := []archiveFile{
		{Name: "book_meta.txt", Format: "Metadata", Size: "500"},
		{Name: "book_djvu.txt", Format: "DjVuTXT", Size: "12000"},
		{Name: "book_abbyy.gz", Format: "Abbyy GZ", Size: "999999"},
	}
	got := pickArchiveTextFile(files)
	if got != "book_djvu.txt" {
		t.Errorf("pickArchiveTextFile = %q, want book_djvu.txt", got)
	}
}

func TestPickArchiveTextFileReturnsEmptyWhenNoneMatch(t *testing.T) {
	files := []archiveFile{
		{Name: "book.pdf", Format: "PDF", Size: "1000"},
		{Name: "book_meta.txt", Format: "Metadata", Size: "500"},
	}
	if got := pickArchiveTextFile(files); got != "" {
		t.Errorf("pickArchiveTextFile = %q, want empty", got)
	}
}
