package crawler

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"io"
	"log"
	"os"
	"time"

	"github.com/tomiwa-a/hippo/internal/db"
)

func Sync(ctx context.Context, database *db.DB, roots []string, ignores []string) error {
	fileChan := Walk(roots, ignores)

	for path := range fileChan {
		info, err := os.Stat(path)
		if err != nil {
			log.Printf("Failed to stat %s: %v", path, err)
			continue
		}

		mtime := info.ModTime().Unix()

		existing, err := database.GetFile(ctx, path)
		if err != nil {
			log.Printf("DB error for %s: %v", path, err)
			continue
		}

		if existing != nil && existing.LastModified == mtime {
			continue
		}

		hash, err := hashFile(path)
		if err != nil {
			log.Printf("Failed to hash %s: %v", path, err)
			continue
		}

		if existing != nil && existing.Hash == hash {
			existing.LastModified = mtime
			if err := database.UpsertFile(ctx, existing); err != nil {
				log.Printf("Failed to update mtime for %s: %v", path, err)
			}
			continue
		}

		log.Printf("Change detected: %s", path)

		f := &db.File{
			Path:         path,
			Hash:         hash,
			LastModified: mtime,
			IndexedAt:    time.Now().Unix(),
		}

		if err := database.UpsertFile(ctx, f); err != nil {
			log.Printf("Failed to upsert %s: %v", path, err)
		}
	}

	return nil
}

func hashFile(path string) (string, error) {
	f, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer f.Close()

	h := sha256.New()
	if _, err := io.Copy(h, f); err != nil {
		return "", err
	}

	return hex.EncodeToString(h.Sum(nil)), nil
}
