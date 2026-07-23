package scrapem

import (
	"os"
	"path/filepath"
	"testing"
)

// TestLoadSeenRecoversFromCorruptFile verifies that a torn/corrupt seen index
// does not halt the run: loadSeen quarantines it and returns an empty index so
// syncSeenFromInbox can rebuild from the inbox notes on disk.
func TestLoadSeenRecoversFromCorruptFile(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "seen-urls.json")
	// Simulate a torn write: valid JSON prefix cut off mid-object.
	if err := os.WriteFile(path, []byte(`{"urls":{"https://example.com/a":`), 0o644); err != nil {
		t.Fatalf("seed corrupt file: %v", err)
	}

	seen, err := loadSeen(path)
	if err != nil {
		t.Fatalf("loadSeen returned error on corrupt file, want graceful recovery: %v", err)
	}
	if seen == nil || seen.URLs == nil {
		t.Fatal("loadSeen returned nil index after recovery")
	}
	if len(seen.URLs) != 0 {
		t.Fatalf("expected empty index after recovery, got %d entries", len(seen.URLs))
	}
	if _, err := os.Stat(path + ".corrupt"); err != nil {
		t.Fatalf("corrupt file was not quarantined to %s.corrupt: %v", path, err)
	}
	if _, err := os.Stat(path); !os.IsNotExist(err) {
		t.Fatalf("corrupt file should have been moved aside, but %s still exists", path)
	}
}

// TestAtomicWriteFileReplacesContent verifies atomicWriteFile writes the exact
// bytes with the requested mode and overwrites an existing file in place,
// leaving no temp files behind.
func TestAtomicWriteFileReplacesContent(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "note.md")
	if err := os.WriteFile(path, []byte("old"), 0o644); err != nil {
		t.Fatalf("seed: %v", err)
	}
	if err := atomicWriteFile(path, []byte("new content"), 0o644); err != nil {
		t.Fatalf("atomicWriteFile: %v", err)
	}
	got, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("read back: %v", err)
	}
	if string(got) != "new content" {
		t.Fatalf("content = %q, want %q", got, "new content")
	}
	entries, err := os.ReadDir(dir)
	if err != nil {
		t.Fatalf("readdir: %v", err)
	}
	if len(entries) != 1 {
		t.Fatalf("expected only the target file to remain, got %d entries: %v", len(entries), entries)
	}
}
