package scrapem

import (
	"context"
	"fmt"
	"io/fs"
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

	groups := map[string][]string{}
	err := filepath.WalkDir(inbox, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
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
		if url == "" {
			return nil
		}
		groups[url] = append(groups[url], path)
		return nil
	})
	if err != nil {
		return err
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
