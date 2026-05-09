package scrapem

import (
	"encoding/json"
	"os"
	"path/filepath"
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
