package scrapem

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"strings"

	"github.com/ledongthuc/pdf"
)

var pdfSpaceRe = regexp.MustCompile(`[ \t]+`)
var pdfHyphenBreakRe = regexp.MustCompile(`([A-Za-z])\s*-\s*\n\s*([A-Za-z])`)
var pdfSoftBreakRe = regexp.MustCompile(`([a-z,;:])\n([a-z])`)

func (r *Runner) enrichPDF(ctx context.Context, item *Item) error {
	if item.PDF == "" {
		return nil
	}
	if err := r.wait(ctx); err != nil {
		return err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, item.PDF, nil)
	if err != nil {
		return err
	}
	req.Header.Set("User-Agent", r.cfg.Scrape.UserAgent)
	req.Header.Set("Accept", "application/pdf,*/*;q=0.8")

	resp, err := r.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Errorf("unexpected status %s", resp.Status)
	}
	if resp.ContentLength > r.cfg.Scrape.MaxPDFBytes {
		return fmt.Errorf("pdf too large: %d bytes", resp.ContentLength)
	}

	tmp, err := os.CreateTemp("", "scrapem-*.pdf")
	if err != nil {
		return err
	}
	tmpPath := tmp.Name()
	defer os.Remove(tmpPath)
	defer tmp.Close()

	limit := r.cfg.Scrape.MaxPDFBytes + 1
	written, err := io.Copy(tmp, io.LimitReader(resp.Body, limit))
	if err != nil {
		return err
	}
	if written > r.cfg.Scrape.MaxPDFBytes {
		return fmt.Errorf("pdf too large: exceeded %d bytes", r.cfg.Scrape.MaxPDFBytes)
	}
	if err := tmp.Close(); err != nil {
		return err
	}

	text, err := extractPDFText(tmpPath)
	if err != nil {
		return err
	}
	item.PDFText, item.PDFTextTruncated = truncateRunes(cleanPDFText(text), r.cfg.Scrape.MaxPDFChars)
	return nil
}

func extractPDFText(path string) (string, error) {
	f, reader, err := pdf.Open(path)
	if err != nil {
		return "", err
	}
	defer f.Close()

	plain, err := reader.GetPlainText()
	if err != nil {
		return "", err
	}
	data, err := io.ReadAll(plain)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func cleanPDFText(input string) string {
	input = pdfHyphenBreakRe.ReplaceAllString(input, "$1$2")
	input = pdfSoftBreakRe.ReplaceAllString(input, "$1 $2")
	lines := strings.Split(strings.ReplaceAll(input, "\r\n", "\n"), "\n")
	out := make([]string, 0, len(lines))
	blank := false
	for _, line := range lines {
		line = strings.TrimSpace(pdfSpaceRe.ReplaceAllString(line, " "))
		if line == "" {
			if !blank {
				out = append(out, "")
			}
			blank = true
			continue
		}
		out = append(out, line)
		blank = false
	}
	return strings.TrimSpace(strings.Join(out, "\n"))
}

func truncateRunes(input string, max int) (string, bool) {
	if max <= 0 {
		return input, false
	}
	runes := []rune(input)
	if len(runes) <= max {
		return input, false
	}
	return strings.TrimSpace(string(runes[:max])), true
}
