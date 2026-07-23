package scrapem

import (
	"encoding/json"
	"io/fs"
	"log"
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
		// A corrupt seen index (e.g. a torn write left by an earlier crash)
		// must not halt every future run. Quarantine it and start empty;
		// syncSeenFromInbox then rebuilds the index from the inbox notes that
		// already exist on disk, so nothing is re-scraped.
		log.Printf("seen index at %s is unparseable (%v); quarantining to %s.corrupt and rebuilding from inbox", path, err, path)
		if rerr := os.Rename(path, path+".corrupt"); rerr != nil {
			log.Printf("could not quarantine corrupt seen index: %v", rerr)
		}
		return &Seen{URLs: map[string]string{}}, nil
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
	added := 0
	err := filepath.WalkDir(inbox, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			if os.IsNotExist(err) {
				return nil
			}
			return err
		}
		if d.IsDir() || filepath.Ext(d.Name()) != ".md" {
			return nil
		}
		data, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		match := sourceLineRe.FindStringSubmatch(string(data))
		if len(match) < 2 {
			return nil
		}
		url := strings.TrimSpace(match[1])
		if url == "" || seen.Has(url) {
			return nil
		}
		seen.URLs[url] = ""
		added++
		return nil
	})
	if err != nil {
		return added, err
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
	return atomicWriteFile(path, append(data, '\n'), 0o644)
}
