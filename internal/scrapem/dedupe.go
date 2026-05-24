package scrapem

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

// Dedupe removes inbox markdown files that share the same `source:` URL,
// keeping a single canonical file per URL. The canonical file is chosen as
// the shortest base name (so `foo.md` beats `foo-2.md`); ties are broken
// alphabetically. After cleanup, seen-*.json is rewritten from the
// surviving files. Pass dryRun=true to preview without deleting.
func (r *Runner) Dedupe(_ context.Context, dryRun bool) error {
	inbox := filepath.Join(r.cfg.Vault.Root, r.cfg.Vault.Inbox)
	entries, err := os.ReadDir(inbox)
	if err != nil {
		return err
	}

	groups := map[string][]string{}
	for _, entry := range entries {
		if entry.IsDir() || filepath.Ext(entry.Name()) != ".md" {
			continue
		}
		path := filepath.Join(inbox, entry.Name())
		data, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		match := sourceLineRe.FindStringSubmatch(string(data))
		if len(match) < 2 {
			continue
		}
		url := strings.TrimSpace(match[1])
		if url == "" {
			continue
		}
		groups[url] = append(groups[url], path)
	}

	statePath := filepath.Join(r.cfg.Vault.Root, ".scrapem", r.cfg.Vault.SeenFile)
	seen, err := loadSeen(statePath)
	if err != nil {
		return err
	}

	var removed int
	for url, paths := range groups {
		if len(paths) <= 1 {
			if _, ok := seen.URLs[url]; !ok {
				seen.URLs[url] = ""
			}
			continue
		}
		sort.Slice(paths, func(i, j int) bool {
			ni, nj := filepath.Base(paths[i]), filepath.Base(paths[j])
			if len(ni) != len(nj) {
				return len(ni) < len(nj)
			}
			return ni < nj
		})
		keep := paths[0]
		for _, dup := range paths[1:] {
			if dryRun {
				log.Printf("[dry-run] would remove %s (dup of %s)", dup, filepath.Base(keep))
			} else {
				if err := os.Remove(dup); err != nil {
					return fmt.Errorf("remove %s: %w", dup, err)
				}
				log.Printf("removed %s (dup of %s)", dup, filepath.Base(keep))
			}
			removed++
		}
		if _, ok := seen.URLs[url]; !ok {
			seen.URLs[url] = ""
		}
	}

	if !dryRun {
		if err := seen.Save(statePath); err != nil {
			return err
		}
	}
	log.Printf("dedupe summary: %d files removed across %d urls", removed, len(groups))
	return nil
}
