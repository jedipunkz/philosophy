package scrapem

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"
)

var unsafeFileRe = regexp.MustCompile(`[/:*?"<>|\\]+`)
var sourceLineRe = regexp.MustCompile(`(?m)^source:\s*"?([^"\n]+)"?\s*$`)

func writeMarkdown(inbox string, item Item) (string, error) {
	now := time.Now()
	base := slug(fmt.Sprintf("%s-%s", now.Format("2006-01-02"), item.Title))
	if base == "" {
		base = now.Format("2006-01-02-150405")
	}
	path := filepath.Join(inbox, base+".md")
	for i := 2; ; i++ {
		if _, err := os.Stat(path); os.IsNotExist(err) {
			break
		}
		path = filepath.Join(inbox, fmt.Sprintf("%s-%d.md", base, i))
	}

	content := chooseRenderer(now, item)
	return path, os.WriteFile(path, []byte(content), 0o644)
}

func updateMarkdownBySource(inbox string, item Item) error {
	path, capturedAt, err := findMarkdownBySource(inbox, item.URL)
	if err != nil {
		return err
	}
	if path == "" {
		_, err := writeMarkdown(inbox, item)
		return err
	}
	when := time.Now()
	if capturedAt != "" {
		if parsed, err := time.Parse(time.RFC3339, capturedAt); err == nil {
			when = parsed
		}
	}
	return os.WriteFile(path, []byte(chooseRenderer(when, item)), 0o644)
}

func findMarkdownBySource(inbox, source string) (string, string, error) {
	entries, err := os.ReadDir(inbox)
	if err != nil {
		return "", "", err
	}
	for _, entry := range entries {
		if entry.IsDir() || filepath.Ext(entry.Name()) != ".md" {
			continue
		}
		path := filepath.Join(inbox, entry.Name())
		data, err := os.ReadFile(path)
		if err != nil {
			return "", "", err
		}
		text := string(data)
		match := sourceLineRe.FindStringSubmatch(text)
		if len(match) < 2 || match[1] != source {
			continue
		}
		return path, frontMatterValue(text, "captured_at"), nil
	}
	return "", "", nil
}

func renderMarkdown(now time.Time, item Item) string {
	var b strings.Builder
	b.WriteString("---\n")
	writeYAMLString(&b, "source", item.URL)
	writeYAMLString(&b, "title", item.Title)
	writeYAMLString(&b, "author", item.Author)
	writeYAMLString(&b, "year", item.Year)
	writeYAMLString(&b, "publication", item.Info)
	writeYAMLString(&b, "download", item.Download)
	writeYAMLString(&b, "pdf", item.PDF)
	writeYAMLString(&b, "captured_at", now.Format(time.RFC3339))
	writeYAMLString(&b, "updated_at", time.Now().Format(time.RFC3339))
	writeYAMLString(&b, "capture_tool", "scrapem")
	writeYAMLString(&b, "source_name", item.SourceName)
	writeYAMLString(&b, "keyword", item.Keyword)
	writeYAMLString(&b, "query", item.Query)
	b.WriteString("tags:\n")
	for _, tag := range item.Tags {
		fmt.Fprintf(&b, "  - %q\n", tag)
	}
	b.WriteString("status: raw\n")
	b.WriteString("---\n\n")

	fmt.Fprintf(&b, "# %s\n\n", item.Title)
	if item.Author != "" || item.Year != "" {
		fmt.Fprintf(&b, "- 著者: %s\n", emptyDash(item.Author))
		fmt.Fprintf(&b, "- 年: %s\n", emptyDash(item.Year))
	}
	if item.Info != "" {
		fmt.Fprintf(&b, "- 掲載情報: %s\n", item.Info)
	}
	fmt.Fprintf(&b, "- 情報源: [%s](%s)\n", item.SourceName, item.URL)
	if item.Download != "" {
		fmt.Fprintf(&b, "- ダウンロード: %s\n", item.Download)
	}
	if item.PDF != "" {
		fmt.Fprintf(&b, "- PDF: %s\n", item.PDF)
	}
	writeObsidianLinks(&b, item)
	b.WriteString("\n## Abstract\n\n")
	if item.Abstract == "" {
		b.WriteString("取得したページに要旨は含まれていない。\n")
	} else {
		b.WriteString(item.Abstract)
		b.WriteString("\n")
	}
	if item.Citation != "" {
		b.WriteString("\n## Citation\n\n")
		b.WriteString(item.Citation)
		b.WriteString("\n")
	}
	if item.BibTeX != "" {
		b.WriteString("\n## BibTeX\n\n")
		b.WriteString("```bibtex\n")
		b.WriteString(strings.TrimSpace(item.BibTeX))
		b.WriteString("\n```\n")
	}
	if item.DetailText != "" {
		b.WriteString("\n## Source Details\n\n")
		b.WriteString(item.DetailText)
		b.WriteString("\n")
	}
	if item.PDFText != "" {
		b.WriteString("\n## PDF Text\n\n")
		b.WriteString(item.PDFText)
		b.WriteString("\n")
		if item.PDFTextTruncated {
			b.WriteString("\n<!-- PDF text truncated by scrapem max_pdf_chars. -->\n")
		}
	}
	b.WriteString("\n## Notes\n\n")
	b.WriteString("- 自動収集された未処理ノート。正式ノート化する前に内容と出典を確認する。\n")
	return b.String()
}

func writeObsidianLinks(b *strings.Builder, item Item) {
	links := obsidianLinks(item)
	tags := obsidianTags(item.Tags)
	if len(links) == 0 && len(tags) == 0 {
		return
	}

	b.WriteString("\n## Obsidian Links\n\n")
	for _, link := range links {
		fmt.Fprintf(b, "- %s\n", link)
	}
	if len(tags) > 0 {
		fmt.Fprintf(b, "- 関連タグ: %s\n", strings.Join(tags, " "))
	}
}

func obsidianLinks(item Item) []string {
	seen := map[string]bool{}
	var links []string
	add := func(label, target, alias string) {
		target = strings.TrimSpace(target)
		if target == "" || seen[target] {
			return
		}
		seen[target] = true
		if alias == "" || alias == target {
			links = append(links, fmt.Sprintf("%s: [[%s]]", label, target))
			return
		}
		links = append(links, fmt.Sprintf("%s: [[%s|%s]]", label, target, alias))
	}

	keyword := strings.TrimSpace(item.Keyword)
	if keyword != "" {
		add("研究動向", fmt.Sprintf("研究動向/%s-現代研究動向", keyword), fmt.Sprintf("%s-現代研究動向", keyword))
		add("キーワード", keyword, "")
	}
	for _, tag := range item.Tags {
		tag = strings.TrimSpace(tag)
		if tag == "" {
			continue
		}
		add("関連分野", tag, "")
	}
	return links
}

func obsidianTags(tags []string) []string {
	seen := map[string]bool{}
	var out []string
	for _, tag := range tags {
		tag = strings.TrimSpace(tag)
		if tag == "" {
			continue
		}
		tag = strings.Join(strings.Fields(tag), "-")
		if tag == "" || seen[tag] {
			continue
		}
		seen[tag] = true
		out = append(out, "#"+tag)
	}
	return out
}

func writeYAMLString(b *strings.Builder, key, value string) {
	if value == "" {
		fmt.Fprintf(b, "%s: \"\"\n", key)
		return
	}
	fmt.Fprintf(b, "%s: %q\n", key, value)
}

func slug(input string) string {
	s := unsafeFileRe.ReplaceAllString(input, "-")
	s = strings.Trim(s, " .-")
	if len([]rune(s)) <= 90 {
		return s
	}
	runes := []rune(s)
	return strings.Trim(string(runes[:90]), " .-")
}

func emptyDash(input string) string {
	if input == "" {
		return "-"
	}
	return input
}

func chooseRenderer(now time.Time, item Item) string {
	if item.ItemType == "book" {
		return renderBookMarkdown(now, item)
	}
	return renderMarkdown(now, item)
}

func frontMatterValue(text, key string) string {
	re := regexp.MustCompile(`(?m)^` + regexp.QuoteMeta(key) + `:\s*"?([^"\n]+)"?\s*$`)
	match := re.FindStringSubmatch(text)
	if len(match) < 2 {
		return ""
	}
	return match[1]
}
