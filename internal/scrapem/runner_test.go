package scrapem

import (
	"context"
	"net/http"
	"net/http/httptest"
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
