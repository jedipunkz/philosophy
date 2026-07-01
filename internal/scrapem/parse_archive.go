package scrapem

import (
	"encoding/json"
	"io"
	"math"
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
// inconsistently represents: a bare string or number for a single value
// (e.g. "creator": "Plato", "year": 1902), an array for multiple values, or
// a mix of strings and numbers within that array.
type stringOrSlice []string

func (s *stringOrSlice) UnmarshalJSON(data []byte) error {
	var raw interface{}
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}
	*s = flattenToStrings(raw)
	return nil
}

func flattenToStrings(raw interface{}) stringOrSlice {
	switch v := raw.(type) {
	case nil:
		return nil
	case string:
		if v == "" {
			return nil
		}
		return stringOrSlice{v}
	case float64:
		return stringOrSlice{formatJSONNumber(v)}
	case []interface{}:
		out := make(stringOrSlice, 0, len(v))
		for _, elem := range v {
			out = append(out, flattenToStrings(elem)...)
		}
		return out
	default:
		return nil
	}
}

// formatJSONNumber renders a JSON number (decoded as float64) the way it
// most likely appeared in the source document: as an integer when it has no
// fractional part (true for identifiers like publication years).
func formatJSONNumber(f float64) string {
	if f == math.Trunc(f) {
		return strconv.FormatInt(int64(f), 10)
	}
	return strconv.FormatFloat(f, 'f', -1, 64)
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
