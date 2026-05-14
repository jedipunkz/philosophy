package scrapem

import (
	"fmt"
	"strings"
	"time"
)

func renderBookMarkdown(now time.Time, item Item) string {
	var b strings.Builder
	b.WriteString("---\n")
	writeYAMLString(&b, "source", item.URL)
	writeYAMLString(&b, "title", item.Title)
	writeYAMLString(&b, "author", item.Author)
	writeYAMLString(&b, "year", item.Year)
	writeYAMLString(&b, "captured_at", now.Format(time.RFC3339))
	writeYAMLString(&b, "updated_at", time.Now().Format(time.RFC3339))
	writeYAMLString(&b, "capture_tool", "scrapem-book")
	writeYAMLString(&b, "source_name", item.SourceName)
	writeYAMLString(&b, "keyword", item.Keyword)
	writeYAMLString(&b, "query", item.Query)
	if item.GutenbergURL != "" {
		writeYAMLString(&b, "gutenberg_url", item.GutenbergURL)
	}
	if item.PublicDomain {
		b.WriteString("public_domain: true\n")
	}
	b.WriteString("subjects:\n")
	for _, s := range item.Subjects {
		fmt.Fprintf(&b, "  - %q\n", s)
	}
	b.WriteString("tags:\n")
	for _, tag := range item.Tags {
		fmt.Fprintf(&b, "  - %q\n", tag)
	}
	b.WriteString("status: raw\n")
	b.WriteString("---\n\n")

	fmt.Fprintf(&b, "# %s\n\n", item.Title)
	fmt.Fprintf(&b, "- 著者: %s\n", emptyDash(item.Author))
	if item.Year != "" {
		fmt.Fprintf(&b, "- 初版: %s\n", item.Year)
	}
	fmt.Fprintf(&b, "- 情報源: [%s](%s)\n", item.SourceName, item.URL)
	if item.Download != "" {
		fmt.Fprintf(&b, "- ダウンロード: %s\n", item.Download)
	}
	if item.PublicDomain {
		b.WriteString("- パブリックドメイン: ✓\n")
	}
	writeBookObsidianLinks(&b, item)

	if len(item.Subjects) > 0 {
		b.WriteString("\n## 主題・分野\n\n")
		for _, s := range item.Subjects {
			fmt.Fprintf(&b, "- %s\n", s)
		}
	}

	if item.Abstract != "" {
		b.WriteString("\n## Abstract\n\n")
		b.WriteString(item.Abstract)
		b.WriteString("\n")
	}

	b.WriteString("\n## Notes\n\n")
	b.WriteString("- 自動収集された未処理ノート。書籍/ フォルダへの統合前に内容と出典を確認する。\n")
	return b.String()
}

func writeBookObsidianLinks(b *strings.Builder, item Item) {
	keyword := strings.TrimSpace(item.Keyword)
	if keyword == "" {
		return
	}
	b.WriteString("\n## Obsidian Links\n\n")
	fmt.Fprintf(b, "- キーワード: [[%s]]\n", keyword)
	fmt.Fprintf(b, "- 書籍: [[書籍/%s]]\n", keyword)
}
