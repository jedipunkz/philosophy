package scrapem

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type Seen struct {
	URLs map[string]string `json:"urls"`
}

func loadSeen(path string) (*Seen, error) {
	data, err := os.ReadFile(path)
	if os.IsNotExist(err) {
		return &Seen{URLs: map[string]string{}}, nil
	}
	if err != nil {
		return nil, err
	}
	var seen Seen
	if err := json.Unmarshal(data, &seen); err != nil {
		return nil, err
	}
	if seen.URLs == nil {
		seen.URLs = map[string]string{}
	}
	return &seen, nil
}

func (s *Seen) Has(rawURL string) bool {
	_, ok := s.URLs[rawURL]
	return ok
}

func (s *Seen) Add(rawURL string) {
	s.URLs[rawURL] = time.Now().Format(time.RFC3339)
}

// syncSeenFromInbox rebuilds the seen index from existing inbox markdown files.
// This makes the seen state self-healing when seen-*.json is missing or stale
// (e.g. when a previous run crashed before Save). Returns the number of URLs added.
func syncSeenFromInbox(inbox string, seen *Seen) (int, error) {
	entries, err := os.ReadDir(inbox)
	if err != nil {
		if os.IsNotExist(err) {
			return 0, nil
		}
		return 0, err
	}
	added := 0
	for _, entry := range entries {
		if entry.IsDir() || filepath.Ext(entry.Name()) != ".md" {
			continue
		}
		data, err := os.ReadFile(filepath.Join(inbox, entry.Name()))
		if err != nil {
			return added, err
		}
		match := sourceLineRe.FindStringSubmatch(string(data))
		if len(match) < 2 {
			continue
		}
		url := strings.TrimSpace(match[1])
		if url == "" || seen.Has(url) {
			continue
		}
		seen.URLs[url] = ""
		added++
	}
	return added, nil
}

func (s *Seen) Save(path string) error {
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return err
	}
	data, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, append(data, '\n'), 0o644)
}
