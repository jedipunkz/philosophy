package scrapem

type Item struct {
	Title            string
	URL              string
	Author           string
	Year             string
	Info             string
	Abstract         string
	Download         string
	PDF              string
	Citation         string
	BibTeX           string
	DetailText       string
	PDFText          string
	PDFTextTruncated bool
	SourceName       string
	Keyword          string
	Query            string
	Tags             []string
	// book-specific fields
	ItemType     string   // "book" or "" (paper)
	Subjects     []string
	PublicDomain bool
	GutenbergURL string
}
