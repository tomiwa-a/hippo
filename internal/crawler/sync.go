package crawler

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"io"
	"log"
	"os"
	"time"

	gitignore "github.com/sabhiram/go-gitignore"
	"github.com/tomiwa-a/hippo/internal/config"
	"github.com/tomiwa-a/hippo/internal/db"
)

type Crawler struct {
	DB        *db.DB
	Config    *config.Config
	IgnoreMap *gitignore.GitIgnore
}

func New(database *db.DB, cfg *config.Config) *Crawler {
	gi := gitignore.CompileIgnoreLines(cfg.Ignore...)
	return &Crawler{
		DB:        database,
		Config:    cfg,
		IgnoreMap: gi,
	}
}

func (c *Crawler) Sync(ctx context.Context) error {
	fileChan := Walk(c.Config.WatchPaths, c.IgnoreMap)

	for path := range fileChan {
		info, err := os.Stat(path)
		if err != nil {
			log.Printf("Failed to stat %s: %v", path, err)
			continue
		}

		if info.Size() > c.Config.MaxSize {
			continue
		}

		mtime := info.ModTime().Unix()
		size := info.Size()

		existing, err := c.DB.GetFile(ctx, path)
		if err != nil {
			log.Printf("DB error for %s: %v", path, err)
			continue
		}

		if existing != nil && existing.LastModified == mtime && existing.Size == size {
			continue
		}

		hash, err := hashFile(path)
		if err != nil {
			log.Printf("Failed to hash %s: %v", path, err)
			continue
		}

		if existing != nil && existing.Hash == hash {
			existing.LastModified = mtime
			existing.Size = size
			if err := c.DB.UpsertFile(ctx, existing); err != nil {
				log.Printf("Failed to update mtime for %s: %v", path, err)
			}
			continue
		}

		log.Printf("Change detected: %s", path)

		f := &db.File{
			Path:         path,
			Hash:         hash,
			LastModified: mtime,
			Size:         size,
			IndexedAt:    time.Now().Unix(),
		}

		if err := c.DB.UpsertFile(ctx, f); err != nil {
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
