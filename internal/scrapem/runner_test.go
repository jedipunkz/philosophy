package scrapem

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/jedipunkz/philosophy/internal/config"
)

func TestIsRetryableStatus(t *testing.T) {
	cases := map[int]bool{
		http.StatusOK:                  false,
		http.StatusForbidden:           false, // durable WAF/IP block; not worth retrying
		http.StatusNotFound:            false,
		http.StatusTooManyRequests:     true,
		http.StatusInternalServerError: true,
		http.StatusBadGateway:          true,
	}
	for status, want := range cases {
		if got := isRetryableStatus(status); got != want {
			t.Errorf("isRetryableStatus(%d) = %v, want %v", status, got, want)
		}
	}
}

func TestRetryDelayHonorsRetryAfterSeconds(t *testing.T) {
	resp := &http.Response{Header: http.Header{"Retry-After": []string{"5"}}}
	if got := retryDelay(resp, 0); got != 5*time.Second {
		t.Fatalf("retryDelay = %v, want 5s", got)
	}
}

func TestRetryDelayCapsAtMax(t *testing.T) {
	resp := &http.Response{Header: http.Header{"Retry-After": []string{"3600"}}}
	if got := retryDelay(resp, 0); got != maxRetryDelay {
		t.Fatalf("retryDelay = %v, want capped %v", got, maxRetryDelay)
	}
}

func TestRetryDelayFallsBackToExponentialBackoff(t *testing.T) {
	resp := &http.Response{Header: http.Header{}}
	if got := retryDelay(resp, 0); got != baseRetryDelay {
		t.Fatalf("retryDelay attempt 0 = %v, want %v", got, baseRetryDelay)
	}
	if got := retryDelay(resp, 1); got != 2*baseRetryDelay {
		t.Fatalf("retryDelay attempt 1 = %v, want %v", got, 2*baseRetryDelay)
	}
}

func newTestRunner(t *testing.T) *Runner {
	t.Helper()
	return New(config.Config{
		Scrape: config.ScrapeConfig{
			UserAgent:      "test-agent",
			RequestTimeout: "5s",
			RequestDelay:   "0s",
		},
	})
}

func TestDoRequestRetriesOnTooManyRequestsThenSucceeds(t *testing.T) {
	attempts := 0
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		attempts++
		if attempts == 1 {
			w.Header().Set("Retry-After", "0")
			w.WriteHeader(http.StatusTooManyRequests)
			return
		}
		w.WriteHeader(http.StatusOK)
	}))
	defer srv.Close()

	r := newTestRunner(t)
	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, srv.URL, nil)
	if err != nil {
		t.Fatalf("NewRequestWithContext: %v", err)
	}
	resp, err := r.doRequest(context.Background(), req)
	if err != nil {
		t.Fatalf("doRequest returned error: %v", err)
	}
	resp.Body.Close()
	if attempts != 2 {
		t.Fatalf("attempts = %d, want 2", attempts)
	}
}

func TestDoRequestDoesNotRetryOnForbidden(t *testing.T) {
	attempts := 0
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		attempts++
		w.WriteHeader(http.StatusForbidden)
	}))
	defer srv.Close()

	r := newTestRunner(t)
	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, srv.URL, nil)
	if err != nil {
		t.Fatalf("NewRequestWithContext: %v", err)
	}
	if _, err := r.doRequest(context.Background(), req); err == nil {
		t.Fatal("expected error for 403 response, got nil")
	}
	if attempts != 1 {
		t.Fatalf("attempts = %d, want 1 (no retry on 403)", attempts)
	}
}

func TestDoRequestGivesUpAfterMaxRetryAttempts(t *testing.T) {
	attempts := 0
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		attempts++
		w.Header().Set("Retry-After", "0")
		w.WriteHeader(http.StatusTooManyRequests)
	}))
	defer srv.Close()

	r := newTestRunner(t)
	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, srv.URL, nil)
	if err != nil {
		t.Fatalf("NewRequestWithContext: %v", err)
	}
	if _, err := r.doRequest(context.Background(), req); err == nil {
		t.Fatal("expected error after exhausting retries, got nil")
	}
	if attempts != maxRetryAttempts+1 {
		t.Fatalf("attempts = %d, want %d", attempts, maxRetryAttempts+1)
	}
}

// TestRunSkipsArchiveItemsWithNoExtractableText exercises the full Run()
// pipeline against a fake archive.org-shaped server. It reproduces the
// production case where advancedsearch.php matches an item (public domain,
// not access-restricted) that nonetheless has no OCR text derivative — this
// previously produced an inbox note with an empty body (see e.g. the
// "Carlton Plato Washington" court filing that matched a bare "Plato" query).
func TestRunSkipsArchiveItemsWithNoExtractableText(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/advancedsearch.php", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"response":{"docs":[
			{"identifier":"no-text-item","title":"An Item With No OCR Text","creator":"Someone","year":1900,"access-restricted-item":false}
		]}}`))
	})
	mux.HandleFunc("/metadata/no-text-item", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"files":[{"name":"no-text-item_meta.xml","format":"Metadata","size":"100"}]}`))
	})
	srv := httptest.NewServer(mux)
	defer srv.Close()

	vaultRoot := t.TempDir()
	cfg := config.Config{
		Vault: config.VaultConfig{Root: vaultRoot, Inbox: "inbox", SeenFile: "seen-urls.json"},
		Scrape: config.ScrapeConfig{
			UserAgent:      "test-agent",
			RequestTimeout: "5s",
			RequestDelay:   "0s",
			MaxResults:     5,
			MaxBookBytes:   10 << 20,
			MaxBookChars:   100000,
		},
		Sources: []config.SourceConfig{
			{
				Name:      "archive",
				Enabled:   true,
				Type:      "archive_api",
				BaseURL:   srv.URL,
				SearchURL: srv.URL + "/advancedsearch.php?q={query}",
			},
		},
		Keywords: []config.KeywordConfig{
			{
				Name:    "テスト",
				Queries: []string{"Plato"},
				Sources: []string{"archive"},
			},
		},
	}

	if err := New(cfg).Run(context.Background()); err != nil {
		t.Fatalf("Run returned error: %v", err)
	}

	inbox := filepath.Join(vaultRoot, "inbox")
	for _, f := range inboxMarkdownFiles(t, inbox) {
		t.Errorf("expected no notes to be written for a textless item, found %s", f)
	}
}

// TestRunSkipsCrossrefItemsWithNoBody exercises Run() against a fake
// Crossref-shaped server that returns an item with no abstract (and, being a
// paper source, no PDF/detail fetch). This reproduces the production case
// where crossref matched works with only bibliographic metadata, producing
// inbox notes whose only body was the "要旨は含まれていない" placeholder.
func TestRunSkipsCrossrefItemsWithNoBody(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"message":{"items":[
			{"title":["A Work With No Abstract"],"DOI":"10.1000/no-abstract","container-title":["Some Journal"]}
		]}}`))
	}))
	defer srv.Close()

	vaultRoot := t.TempDir()
	cfg := config.Config{
		Vault: config.VaultConfig{Root: vaultRoot, Inbox: "inbox", SeenFile: "seen-urls.json"},
		Scrape: config.ScrapeConfig{
			UserAgent:      "test-agent",
			RequestTimeout: "5s",
			RequestDelay:   "0s",
			MaxResults:     5,
		},
		Sources: []config.SourceConfig{
			{
				Name:      "crossref",
				Enabled:   true,
				Type:      "crossref_api",
				BaseURL:   srv.URL,
				SearchURL: srv.URL + "/works?query={query}",
			},
		},
		Keywords: []config.KeywordConfig{
			{
				Name:    "テスト",
				Queries: []string{"nietzsche"},
				Sources: []string{"crossref"},
			},
		},
	}

	if err := New(cfg).Run(context.Background()); err != nil {
		t.Fatalf("Run returned error: %v", err)
	}

	inbox := filepath.Join(vaultRoot, "inbox")
	for _, f := range inboxMarkdownFiles(t, inbox) {
		t.Errorf("expected no notes to be written for a bodyless crossref item, found %s", f)
	}
}

// TestRunSkipsWikipediaPersonBiography exercises Run() against a fake
// MediaWiki-shaped server whose search generator returns a person biography
// (as ja.wikipedia does when a work query like "Critique of Pure Reason"
// matches the article on the Kant scholar P.F. Strawson). Such an article is
// not a 著作 and must not become an inbox note.
func TestRunSkipsWikipediaPersonBiography(t *testing.T) {
	bio := `ピーター・フレデリック・ストローソン （Peter Frederick Strawson、1919年11月23日 - 2006年2月13日）は、イギリスの哲学者。`
	mux := http.NewServeMux()
	mux.HandleFunc("/w/api.php", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		// Both the generator=search call and the per-page extract call hit
		// this same endpoint; returning the bio for both is fine.
		w.Write([]byte(`{"query":{"pages":[{"pageid":123,"title":"ピーター・フレデリック・ストローソン","extract":` +
			mustJSONString(bio) + `}]}}`))
	})
	srv := httptest.NewServer(mux)
	defer srv.Close()

	vaultRoot := t.TempDir()
	cfg := config.Config{
		Vault: config.VaultConfig{Root: vaultRoot, Inbox: "inbox", SeenFile: "seen-urls.json"},
		Scrape: config.ScrapeConfig{
			UserAgent:      "test-agent",
			RequestTimeout: "5s",
			RequestDelay:   "0s",
			MaxResults:     5,
			MaxBookBytes:   10 << 20,
			MaxBookChars:   100000,
		},
		Sources: []config.SourceConfig{
			{
				Name:      "wikipedia_ja",
				Enabled:   true,
				Type:      "wikipedia_api",
				BaseURL:   srv.URL,
				SearchURL: srv.URL + "/w/api.php?gsrsearch={query}",
			},
		},
		Keywords: []config.KeywordConfig{
			{
				Name:    "カント",
				Queries: []string{"Critique of Pure Reason"},
				Sources: []string{"wikipedia_ja"},
			},
		},
	}

	if err := New(cfg).Run(context.Background()); err != nil {
		t.Fatalf("Run returned error: %v", err)
	}

	inbox := filepath.Join(vaultRoot, "inbox")
	for _, f := range inboxMarkdownFiles(t, inbox) {
		t.Errorf("expected no notes for a person biography, found %s", f)
	}
}

func mustJSONString(s string) string {
	b, err := json.Marshal(s)
	if err != nil {
		panic(err)
	}
	return string(b)
}

// inboxMarkdownFiles returns every .md file under inbox (recursively), as paths
// relative to inbox. Notes are grouped under inbox/<著者>/ subdirectories, so a
// flat os.ReadDir is not enough to find them.
func inboxMarkdownFiles(t *testing.T, inbox string) []string {
	t.Helper()
	var files []string
	err := filepath.WalkDir(inbox, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			if os.IsNotExist(err) {
				return nil
			}
			return err
		}
		if d.IsDir() || !strings.HasSuffix(d.Name(), ".md") {
			return nil
		}
		rel, err := filepath.Rel(inbox, path)
		if err != nil {
			return err
		}
		files = append(files, rel)
		return nil
	})
	if err != nil {
		t.Fatalf("walk inbox: %v", err)
	}
	return files
}

// TestRunWritesNoteIntoKeywordSubdir verifies the happy path: a note with a
// usable body is written under inbox/<著者>/ where <著者> is the scrape keyword,
// rather than flat in inbox/.
func TestRunWritesNoteIntoKeywordSubdir(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"message":{"items":[
			{"title":["Nietzsche on Tragedy"],"DOI":"10.1000/has-abstract","container-title":["Some Journal"],"abstract":"An abstract about tragedy."}
		]}}`))
	}))
	defer srv.Close()

	vaultRoot := t.TempDir()
	cfg := config.Config{
		Vault: config.VaultConfig{Root: vaultRoot, Inbox: "inbox", SeenFile: "seen-urls.json"},
		Scrape: config.ScrapeConfig{
			UserAgent:      "test-agent",
			RequestTimeout: "5s",
			RequestDelay:   "0s",
			MaxResults:     5,
		},
		Sources: []config.SourceConfig{
			{
				Name:      "crossref",
				Enabled:   true,
				Type:      "crossref_api",
				BaseURL:   srv.URL,
				SearchURL: srv.URL + "/works?query={query}",
			},
		},
		Keywords: []config.KeywordConfig{
			{
				Name:    "ニーチェ",
				Queries: []string{"nietzsche"},
				Sources: []string{"crossref"},
			},
		},
	}

	if err := New(cfg).Run(context.Background()); err != nil {
		t.Fatalf("Run returned error: %v", err)
	}

	inbox := filepath.Join(vaultRoot, "inbox")
	files := inboxMarkdownFiles(t, inbox)
	if len(files) != 1 {
		t.Fatalf("expected exactly 1 note, got %d: %v", len(files), files)
	}
	if dir := filepath.Dir(files[0]); dir != "ニーチェ" {
		t.Errorf("note written to %q, want it under keyword subdir %q", files[0], "ニーチェ")
	}
}
