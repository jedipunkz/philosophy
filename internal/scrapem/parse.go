package scrapem

import (
	"html"
	"io"
	"net/url"
	"regexp"
	"strings"
)

var (
	entryRe    = regexp.MustCompile(`(?is)<li\b[^>]*class=['"][^'"]*entry[^'"]*['"][^>]*>(.*?)</li>`)
	recLinkRe  = regexp.MustCompile(`(?is)<a\b[^>]*href=['"]([^'"]*/rec/[^'"]+)['"][^>]*>(.*?)</a>`)
	nameRe     = regexp.MustCompile(`(?is)<a\b[^>]*class=['"][^'"]*discreet[^'"]*['"][^>]*>\s*<span\b[^>]*class=['"]name['"][^>]*>(.*?)</span>`)
	yearRe     = regexp.MustCompile(`(?is)<span\b[^>]*class=['"][^'"]*pubYear[^'"]*['"][^>]*>(.*?)</span>`)
	infoRe     = regexp.MustCompile(`(?is)<span\b[^>]*class=['"][^'"]*pubInfo[^'"]*['"][^>]*>(.*?)</span>`)
	abstractRe = regexp.MustCompile(`(?is)<div\b[^>]*class=['"][^'"]*abstract[^'"]*['"][^>]*>(.*?)</div>`)
	downloadRe = regexp.MustCompile(`(?is)<a\b[^>]*href=['"]([^'"]+)['"][^>]*>\s*(?:<i\b[^>]*></i>\s*)?Download\s*</a>`)
	metaRe     = regexp.MustCompile(`(?is)<meta\b[^>]*(?:name|property)=['"]([^'"]+)['"][^>]*content=['"]([^'"]*)['"][^>]*>`)
	recTitleRe = regexp.MustCompile(`(?is)<h1\b[^>]*class=['"][^'"]*recTitle[^'"]*['"][^>]*>(.*?)</h1>`)
	authorsRe  = regexp.MustCompile(`(?is)<div\b[^>]*class=['"][^'"]*recAuthors[^'"]*['"][^>]*>(.*?)</div>`)
	pubInfoRe  = regexp.MustCompile(`(?is)<div\b[^>]*class=['"][^'"]*recPubInfo[^'"]*['"][^>]*>(.*?)<!-- for clipboard copy feature -->`)
	descRe     = regexp.MustCompile(`(?is)<div\b[^>]*itemprop=['"]description['"][^>]*>(.*?)</div>`)
	textareaRe = regexp.MustCompile(`(?is)<textarea\b[^>]*id=['"]([^'"]+)['"][^>]*>(.*?)</textarea>`)
	tagRe      = regexp.MustCompile(`(?is)<[^>]+>`)
	spaceRe    = regexp.MustCompile(`\s+`)
)

func parseSearchResults(body io.Reader, base *url.URL) ([]Item, error) {
	data, err := io.ReadAll(body)
	if err != nil {
		return nil, err
	}

	htmlText := string(data)
	matches := entryRe.FindAllStringSubmatch(htmlText, -1)
	items := make([]Item, 0, len(matches))
	for _, match := range matches {
		entry := match[1]
		recURL, title := firstRecLink(entry)
		item := Item{
			Title:    cleanHTML(title),
			URL:      absolutize(base, recURL),
			Author:   cleanHTML(first(nameRe, entry)),
			Year:     cleanHTML(first(yearRe, entry)),
			Info:     cleanHTML(first(infoRe, entry)),
			Abstract: cleanHTML(first(abstractRe, entry)),
			Download: absolutize(base, first(downloadRe, entry)),
		}
		if item.Title == "" || item.URL == "" {
			continue
		}
		items = append(items, item)
	}
	return items, nil
}

func parseDetail(body io.Reader, base *url.URL) (Item, error) {
	data, err := io.ReadAll(body)
	if err != nil {
		return Item{}, err
	}
	htmlText := string(data)
	metas := parseMetas(htmlText)
	textareas := parseTextareas(htmlText)

	item := Item{
		Title:      cleanHTML(first(recTitleRe, htmlText)),
		Author:     cleanHTML(first(authorsRe, htmlText)),
		Info:       cleanHTML(first(pubInfoRe, htmlText)),
		Abstract:   cleanAbstract(first(descRe, htmlText)),
		Download:   absolutize(base, first(downloadRe, htmlText)),
		PDF:        metas["citation_pdf_url"],
		Citation:   cleanHTML(textareas["plaintext"]),
		BibTeX:     strings.TrimSpace(html.UnescapeString(textareas["bibtex"])),
		DetailText: collectDetailText(htmlText),
	}
	if item.Title == "" {
		item.Title = metas["citation_title"]
	}
	if item.Author == "" {
		item.Author = metas["citation_author"]
	}
	if item.Year == "" {
		item.Year = metas["citation_publication_date"]
	}
	if item.Citation == "" {
		item.Citation = metas["citation_format_txt"]
	}
	if item.BibTeX == "" {
		item.BibTeX = metas["citation_format_bib"]
	}
	return item, nil
}

func parseMetas(input string) map[string]string {
	metas := map[string]string{}
	for _, match := range metaRe.FindAllStringSubmatch(input, -1) {
		if len(match) < 3 {
			continue
		}
		metas[match[1]] = strings.TrimSpace(html.UnescapeString(match[2]))
	}
	return metas
}

func parseTextareas(input string) map[string]string {
	textareas := map[string]string{}
	for _, match := range textareaRe.FindAllStringSubmatch(input, -1) {
		if len(match) < 3 {
			continue
		}
		textareas[match[1]] = strings.TrimSpace(html.UnescapeString(match[2]))
	}
	return textareas
}

func cleanAbstract(input string) string {
	s := cleanHTML(input)
	s = strings.TrimPrefix(s, "Abstract ")
	return strings.TrimSpace(s)
}

func collectDetailText(input string) string {
	parts := []string{}
	for _, heading := range []string{"Archival history", "Categories", "Keywords", "DOI"} {
		re := regexp.MustCompile(`(?is)<h1\b[^>]*>` + regexp.QuoteMeta(heading) + `</h1>(.*?)(?:<h1\b|</div>\s*</div>\s*<div class="row|<div class='entry-info-partial'>)`)
		if match := re.FindStringSubmatch(input); len(match) >= 2 {
			text := cleanHTML(match[1])
			if text != "" && !strings.HasPrefix(text, "Add ") {
				parts = append(parts, heading+": "+text)
			}
		}
	}
	return strings.Join(parts, "\n")
}

func firstRecLink(input string) (string, string) {
	matches := recLinkRe.FindAllStringSubmatch(input, -1)
	for _, match := range matches {
		if len(match) < 3 {
			continue
		}
		if strings.Contains(match[1], "#analytics") {
			continue
		}
		return match[1], match[2]
	}
	return "", ""
}

func first(re *regexp.Regexp, input string) string {
	match := re.FindStringSubmatch(input)
	if len(match) < 2 {
		return ""
	}
	return match[1]
}

func cleanHTML(input string) string {
	noTags := tagRe.ReplaceAllString(input, " ")
	decoded := html.UnescapeString(noTags)
	return strings.TrimSpace(spaceRe.ReplaceAllString(decoded, " "))
}

func absolutize(base *url.URL, raw string) string {
	if raw == "" {
		return ""
	}
	parsed, err := url.Parse(html.UnescapeString(raw))
	if err != nil {
		return ""
	}
	return base.ResolveReference(parsed).String()
}
