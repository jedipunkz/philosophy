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
