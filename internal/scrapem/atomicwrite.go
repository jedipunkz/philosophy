package scrapem

import (
	"os"
	"path/filepath"
)

// atomicWriteFile writes data to path atomically by writing to a temporary
// file in the same directory and renaming it into place. rename(2) is atomic
// on a POSIX filesystem, so a crash (SIGKILL/OOM/power loss) mid-write leaves
// either the previous file or the complete new one — never a truncated one.
//
// This matters most for .scrapem/seen-urls.json: it is rewritten after every
// new note, and a torn write there is unparseable JSON that would otherwise
// halt every subsequent run. The same guarantee protects the inbox notes.
//
// The temporary file is dot-prefixed and does not end in ".md", so the
// directory walkers (syncSeenFromInbox, findMarkdownBySource, Dedupe) ignore a
// leftover temp from an interrupted write.
func atomicWriteFile(path string, data []byte, perm os.FileMode) error {
	dir := filepath.Dir(path)
	tmp, err := os.CreateTemp(dir, "."+filepath.Base(path)+".tmp-*")
	if err != nil {
		return err
	}
	tmpName := tmp.Name()
	// Remove the temp file if we return before the rename succeeds.
	defer func() {
		if tmpName != "" {
			os.Remove(tmpName)
		}
	}()

	if _, err := tmp.Write(data); err != nil {
		tmp.Close()
		return err
	}
	if err := tmp.Sync(); err != nil {
		tmp.Close()
		return err
	}
	if err := tmp.Close(); err != nil {
		return err
	}
	if err := os.Chmod(tmpName, perm); err != nil {
		return err
	}
	if err := os.Rename(tmpName, path); err != nil {
		return err
	}
	tmpName = "" // rename succeeded; nothing to clean up
	return nil
}
