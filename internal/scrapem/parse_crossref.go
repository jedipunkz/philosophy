package scrapem

import (
	"encoding/json"
	"fmt"
	"io"
	"regexp"
	"strconv"
	"strings"
)

type crossrefResponse struct {
	Message struct {
		Items []crossrefItem `json:"items"`
	} `json:"message"`
}

type crossrefItem struct {
	DOI            string             `json:"DOI"`
	URL            string             `json:"URL"`
	Title          []string           `json:"title"`
	Author         []crossrefAuthor   `json:"author"`
	Abstract       string             `json:"abstract"`
	ContainerTitle []string           `json:"container-title"`
	Published      *crossrefDateParts `json:"published"`
}

type crossrefAuthor struct {
	Given  string `json:"given"`
	Family string `json:"family"`
}

type crossrefDateParts struct {
	DateParts [][]int `json:"date-parts"`
}

var jatsTagRe = regexp.MustCompile(`(?is)</?jats:[a-z0-9]+[^>]*>`)

// parseCrossrefResults parses a Crossref /works search response. Unlike
// philarchive's HTML search + detail-page pair, Crossref returns enough
// bibliographic metadata (title, authors, year, container, abstract, DOI) in
// a single call, so no secondary detail fetch is needed.
func parseCrossrefResults(body io.Reader) ([]Item, error) {
	var resp crossrefResponse
	if err := json.NewDecoder(body).Decode(&resp); err != nil {
		return nil, err
	}
	items := make([]Item, 0, len(resp.Message.Items))
	for _, c := range resp.Message.Items {
		title := strings.TrimSpace(strings.Join(c.Title, " "))
		if title == "" || c.DOI == "" {
			continue
		}
		pageURL := strings.TrimSpace(c.URL)
		if pageURL == "" {
			pageURL = "https://doi.org/" + c.DOI
		}
		items = append(items, Item{
			Title:    title,
			URL:      pageURL,
			Author:   joinCrossrefAuthors(c.Author),
			Year:     crossrefYear(c.Published),
			Info:     strings.Join(c.ContainerTitle, ", "),
			Abstract: cleanJATS(c.Abstract),
			Citation: fmt.Sprintf("DOI: %s", c.DOI),
		})
	}
	return items, nil
}

func joinCrossrefAuthors(authors []crossrefAuthor) string {
	names := make([]string, 0, len(authors))
	for _, a := range authors {
		name := strings.TrimSpace(strings.TrimSpace(a.Given) + " " + strings.TrimSpace(a.Family))
		if name != "" {
			names = append(names, name)
		}
	}
	return strings.Join(names, ", ")
}

func crossrefYear(d *crossrefDateParts) string {
	if d == nil || len(d.DateParts) == 0 || len(d.DateParts[0]) == 0 {
		return ""
	}
	return strconv.Itoa(d.DateParts[0][0])
}

// cleanJATS strips Crossref's JATS XML markup (e.g. <jats:p>...</jats:p>)
// from abstract fields, leaving plain text.
func cleanJATS(input string) string {
	if input == "" {
		return ""
	}
	return strings.TrimSpace(jatsTagRe.ReplaceAllString(input, ""))
}
