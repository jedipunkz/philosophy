package scrapem

import (
	"encoding/json"
	"io"
	"strconv"
	"strings"
)

type semanticScholarResponse struct {
	Data []semanticScholarPaper `json:"data"`
}

type semanticScholarPaper struct {
	PaperID       string                     `json:"paperId"`
	Title         string                     `json:"title"`
	Abstract      string                     `json:"abstract"`
	Year          int                        `json:"year"`
	Venue         string                     `json:"venue"`
	URL           string                     `json:"url"`
	Authors       []semanticScholarAuthor    `json:"authors"`
	ExternalIDs   semanticScholarExternalIDs `json:"externalIds"`
	OpenAccessPDF *semanticScholarOpenPDF    `json:"openAccessPdf"`
}

type semanticScholarAuthor struct {
	Name string `json:"name"`
}

// semanticScholarExternalIDs only reads the string identifiers we use. Fields
// like CorpusId are numeric, so we deliberately omit them to avoid decode
// errors and because we do not surface them in notes.
type semanticScholarExternalIDs struct {
	DOI   string `json:"DOI"`
	ArXiv string `json:"ArXiv"`
}

type semanticScholarOpenPDF struct {
	URL string `json:"url"`
}

// parseSemanticScholarResults parses a Semantic Scholar Graph API
// /paper/search response. Like Crossref, the search response already carries
// the abstract and bibliographic metadata (title, authors, year, venue, DOI)
// plus an open-access PDF link when available, so no secondary detail fetch is
// needed. When openAccessPdf.url is present it is set as the item's PDF so the
// shared PDF-extraction pipeline can pull the full text.
func parseSemanticScholarResults(body io.Reader) ([]Item, error) {
	var resp semanticScholarResponse
	if err := json.NewDecoder(body).Decode(&resp); err != nil {
		return nil, err
	}
	items := make([]Item, 0, len(resp.Data))
	for _, p := range resp.Data {
		title := strings.TrimSpace(p.Title)
		if title == "" {
			continue
		}
		pageURL := strings.TrimSpace(p.URL)
		if pageURL == "" && p.PaperID != "" {
			pageURL = "https://www.semanticscholar.org/paper/" + p.PaperID
		}
		if pageURL == "" {
			continue
		}
		item := Item{
			Title:    title,
			URL:      pageURL,
			Author:   joinSemanticScholarAuthors(p.Authors),
			Year:     semanticScholarYear(p.Year),
			Info:     strings.TrimSpace(p.Venue),
			Abstract: strings.TrimSpace(p.Abstract),
			Citation: semanticScholarCitation(p.ExternalIDs),
		}
		if p.OpenAccessPDF != nil {
			if pdf := strings.TrimSpace(p.OpenAccessPDF.URL); pdf != "" {
				item.PDF = pdf
				item.Download = pdf
			}
		}
		items = append(items, item)
	}
	return items, nil
}

func joinSemanticScholarAuthors(authors []semanticScholarAuthor) string {
	names := make([]string, 0, len(authors))
	for _, a := range authors {
		name := strings.TrimSpace(a.Name)
		if name != "" {
			names = append(names, name)
		}
	}
	return strings.Join(names, ", ")
}

func semanticScholarYear(year int) string {
	if year <= 0 {
		return ""
	}
	return strconv.Itoa(year)
}

func semanticScholarCitation(ids semanticScholarExternalIDs) string {
	if doi := strings.TrimSpace(ids.DOI); doi != "" {
		return "DOI: " + doi
	}
	if arxiv := strings.TrimSpace(ids.ArXiv); arxiv != "" {
		return "arXiv: " + arxiv
	}
	return ""
}
