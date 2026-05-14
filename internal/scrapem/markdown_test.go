package scrapem

import (
	"strings"
	"testing"
	"time"
)

func TestRenderMarkdownAddsObsidianLinks(t *testing.T) {
	item := Item{
		Title:      "Jung and Existentialism",
		URL:        "https://example.com/jung",
		SourceName: "philarchive",
		Keyword:    "ユング",
		Query:      "Carl Gustav Jung",
		Tags:       []string{"現代思想", "分析 心理学", "現代思想"},
	}

	got := renderMarkdown(time.Date(2026, 5, 9, 0, 0, 0, 0, time.UTC), item)

	want := []string{
		"## Obsidian Links",
		"- 研究動向: [[研究動向/ユング-現代研究動向|ユング-現代研究動向]]",
		"- キーワード: [[ユング]]",
		"- 関連分野: [[現代思想]]",
		"- 関連分野: [[分析 心理学]]",
		"- 関連タグ: #現代思想 #分析-心理学",
	}
	for _, s := range want {
		if !strings.Contains(got, s) {
			t.Fatalf("rendered markdown does not contain %q:\n%s", s, got)
		}
	}
	if strings.Count(got, "- 関連分野: [[現代思想]]") != 1 {
		t.Fatalf("expected duplicate tags to be linked once:\n%s", got)
	}
}

func TestRenderBookMarkdownIncludesFullText(t *testing.T) {
	item := Item{
		ItemType:          "book",
		Title:             "The Republic",
		URL:               "https://www.gutenberg.org/ebooks/1497",
		GutenbergURL:      "https://www.gutenberg.org/ebooks/1497",
		PlainTextURL:      "https://www.gutenberg.org/cache/epub/1497/pg1497.txt",
		Author:            "Plato",
		Year:              "347",
		SourceName:        "gutenberg",
		Keyword:           "プラトン",
		PublicDomain:      true,
		Description:       "Open Library 由来の概要文。",
		BookText:          "本文の冒頭部分。",
		BookTextTruncated: true,
	}

	got := renderBookMarkdown(time.Date(2026, 5, 14, 0, 0, 0, 0, time.UTC), item)

	want := []string{
		"plain_text_url: \"https://www.gutenberg.org/cache/epub/1497/pg1497.txt\"",
		"public_domain: true",
		"book_text_truncated: true",
		"## 概要",
		"Open Library 由来の概要文。",
		"## Full Text",
		"本文の冒頭部分。",
		"<!-- Book text truncated by scrapem max_book_chars. -->",
	}
	for _, s := range want {
		if !strings.Contains(got, s) {
			t.Fatalf("rendered book markdown does not contain %q:\n%s", s, got)
		}
	}
}

func TestCleanGutenbergTextStripsLicenseMarkers(t *testing.T) {
	raw := "license header line\n\n*** START OF THIS PROJECT GUTENBERG EBOOK THE REPUBLIC ***\n\n本文ライン1\n本文ライン2\n\n*** END OF THIS PROJECT GUTENBERG EBOOK THE REPUBLIC ***\n\nlicense footer line\n"
	got := cleanGutenbergText(raw)
	if strings.Contains(got, "license header") || strings.Contains(got, "license footer") {
		t.Fatalf("license text not stripped:\n%s", got)
	}
	if !strings.Contains(got, "本文ライン1") || !strings.Contains(got, "本文ライン2") {
		t.Fatalf("body lost during clean:\n%s", got)
	}
}

func TestParseWikipediaResults(t *testing.T) {
	body := strings.NewReader(`{
  "batchcomplete": true,
  "query": {
    "pages": [
      {"pageid": 12345, "title": "ハンナ・アーレント", "extract": "ドイツ出身の哲学者。"},
      {"pageid": 67890, "title": "人間の条件", "extract": "アーレント著の代表作。"}
    ]
  }
}`)
	items, err := parseWikipediaResults(body, "https://ja.wikipedia.org")
	if err != nil {
		t.Fatalf("parse error: %v", err)
	}
	if len(items) != 2 {
		t.Fatalf("expected 2 items, got %d", len(items))
	}
	if items[0].URL != "https://ja.wikipedia.org/wiki/%E3%83%8F%E3%83%B3%E3%83%8A%E3%83%BB%E3%82%A2%E3%83%BC%E3%83%AC%E3%83%B3%E3%83%88" {
		t.Fatalf("unexpected URL: %s", items[0].URL)
	}
	if items[0].BookText != "ドイツ出身の哲学者。" {
		t.Fatalf("unexpected extract: %q", items[0].BookText)
	}
}

func TestParseSEPSearchResultsExtractsEntrySlugs(t *testing.T) {
	body := strings.NewReader(`
<html><body>
<a href="/entries/arendt/">Hannah Arendt</a>
<a href="https://plato.stanford.edu/entries/kant/">Kant</a>
<a href="/entries/arendt/#bio">duplicate</a>
<a href="/about/">other</a>
</body></html>`)
	items, err := parseSEPSearchResults(body, "https://plato.stanford.edu")
	if err != nil {
		t.Fatalf("parse error: %v", err)
	}
	if len(items) != 2 {
		t.Fatalf("expected 2 unique entries, got %d", len(items))
	}
	if items[0].URL != "https://plato.stanford.edu/entries/arendt/" {
		t.Fatalf("unexpected first URL: %s", items[0].URL)
	}
	if items[1].URL != "https://plato.stanford.edu/entries/kant/" {
		t.Fatalf("unexpected second URL: %s", items[1].URL)
	}
}

func TestExtractSEPBodyStopsAtBibliography(t *testing.T) {
	html := `<html><body>
<h1>Hannah Arendt</h1>
<p>First paragraph.</p>
<h2>1. Background</h2>
<p>Second paragraph.</p>
<div id="bibliography"><p>Should not appear.</p></div>
</body></html>`
	title := extractSEPTitle(html)
	if title != "Hannah Arendt" {
		t.Fatalf("title: got %q", title)
	}
	body := extractSEPBody(html)
	if !strings.Contains(body, "First paragraph.") || !strings.Contains(body, "Second paragraph.") {
		t.Fatalf("body missing paragraphs:\n%s", body)
	}
	if strings.Contains(body, "Should not appear") {
		t.Fatalf("body included bibliography content:\n%s", body)
	}
}

func TestHTMLToTextPreservesParagraphs(t *testing.T) {
	in := `<p>First</p><p>Second</p><div>Third<br>line break</div>`
	got := htmlToText(in)
	if !strings.Contains(got, "First\n\nSecond") {
		t.Fatalf("paragraph break lost:\n%s", got)
	}
	if !strings.Contains(got, "Third\nline break") {
		t.Fatalf("br lost:\n%s", got)
	}
}

func TestRenderMarkdownOmitsObsidianLinksWhenNoRelations(t *testing.T) {
	item := Item{
		Title:      "Untitled",
		URL:        "https://example.com/untitled",
		SourceName: "philarchive",
	}

	got := renderMarkdown(time.Date(2026, 5, 9, 0, 0, 0, 0, time.UTC), item)

	if strings.Contains(got, "## Obsidian Links") {
		t.Fatalf("expected Obsidian Links section to be omitted:\n%s", got)
	}
}
