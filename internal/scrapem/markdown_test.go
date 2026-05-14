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

func TestParseOpenLibraryDescriptionHandlesBothShapes(t *testing.T) {
	if got := parseOpenLibraryDescription([]byte(`"plain string"`)); got != "plain string" {
		t.Fatalf("string form: got %q", got)
	}
	if got := parseOpenLibraryDescription([]byte(`{"value": "object form", "type": "/type/text"}`)); got != "object form" {
		t.Fatalf("object form: got %q", got)
	}
	if got := parseOpenLibraryDescription(nil); got != "" {
		t.Fatalf("nil: got %q", got)
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
