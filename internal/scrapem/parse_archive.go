package scrapem

import (
	"encoding/json"
	"io"
	"strconv"
	"strings"
)

type archiveSearchResponse struct {
	Response struct {
		Docs []archiveDoc `json:"docs"`
	} `json:"response"`
}

type archiveDoc struct {
	Identifier           string        `json:"identifier"`
	Title                stringOrSlice `json:"title"`
	Creator              stringOrSlice `json:"creator"`
	Year                 stringOrSlice `json:"year"`
	AccessRestrictedItem looseBool     `json:"access-restricted-item"`
}

type archiveMetadataResponse struct {
	Files []archiveFile `json:"files"`
}

type archiveFile struct {
	Name   string `json:"name"`
	Format string `json:"format"`
	Size   string `json:"size"`
}

// stringOrSlice unmarshals a JSON field that Internet Archive's search API
// inconsistently returns as either a single string or an array of strings
// (e.g. "creator" is a string for single-author items, an array otherwise).
type stringOrSlice []string

func (s *stringOrSlice) UnmarshalJSON(data []byte) error {
	var single string
	if err := json.Unmarshal(data, &single); err == nil {
		if single != "" {
			*s = []string{single}
		}
		return nil
	}
	var multi []string
	if err := json.Unmarshal(data, &multi); err != nil {
		return err
	}
	*s = multi
	return nil
}

func (s stringOrSlice) first() string {
	if len(s) == 0 {
		return ""
	}
	return s[0]
}

func (s stringOrSlice) join(sep string) string {
	return strings.Join([]string(s), sep)
}

// looseBool unmarshals a JSON field that may arrive as a bool, a "true"/"false"
// string, or a 0/1 number, all of which Internet Archive's API has been
// observed to send for "access-restricted-item" depending on endpoint.
type looseBool bool

func (b *looseBool) UnmarshalJSON(data []byte) error {
	var v bool
	if err := json.Unmarshal(data, &v); err == nil {
		*b = looseBool(v)
		return nil
	}
	var s string
	if err := json.Unmarshal(data, &s); err == nil {
		*b = looseBool(s == "true" || s == "1")
		return nil
	}
	var n int
	if err := json.Unmarshal(data, &n); err == nil {
		*b = looseBool(n != 0)
		return nil
	}
	return nil
}

// parseArchiveResults parses an archive.org advancedsearch.php response into
// book Items. Full text is not fetched here — enrichArchiveText does a
// separate /metadata/{identifier} lookup to find a downloadable plain-text
// rendition, mirroring the search+detail split used for gutenberg/SEP.
func parseArchiveResults(body io.Reader) ([]Item, error) {
	var resp archiveSearchResponse
	if err := json.NewDecoder(body).Decode(&resp); err != nil {
		return nil, err
	}
	items := make([]Item, 0, len(resp.Response.Docs))
	for _, d := range resp.Response.Docs {
		title := strings.TrimSpace(d.Title.first())
		if title == "" || d.Identifier == "" {
			continue
		}
		items = append(items, Item{
			ItemType:     "book",
			Title:        title,
			URL:          "https://archive.org/details/" + d.Identifier,
			Author:       d.Creator.join(", "),
			Year:         d.Year.first(),
			PublicDomain: !bool(d.AccessRestrictedItem),
		})
	}
	return items, nil
}

// pickArchiveTextFile returns the download filename of the best plain-text
// rendition among an item's files, preferring the largest match (a fuller
// OCR pass) when more than one candidate exists. Returns "" if none found.
func pickArchiveTextFile(files []archiveFile) string {
	var best archiveFile
	bestSize := int64(-1)
	for _, f := range files {
		if !isArchiveTextFormat(f.Format, f.Name) {
			continue
		}
		size, _ := strconv.ParseInt(f.Size, 10, 64)
		if size > bestSize {
			bestSize = size
			best = f
		}
	}
	return best.Name
}

func isArchiveTextFormat(format, name string) bool {
	lower := strings.ToLower(name)
	switch {
	case strings.HasSuffix(lower, "_djvu.txt"):
		return true
	case format == "DjVuTXT":
		return true
	case format == "Text" && strings.HasSuffix(lower, ".txt") && !strings.HasSuffix(lower, "_meta.txt"):
		return true
	}
	return false
}
