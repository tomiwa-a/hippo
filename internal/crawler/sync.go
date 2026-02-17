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
	work      chan string
}

func New(database *db.DB, cfg *config.Config) *Crawler {
	gi := gitignore.CompileIgnoreLines(cfg.Ignore...)
	return &Crawler{
		DB:        database,
		Config:    cfg,
		IgnoreMap: gi,
		work:      make(chan string, 1000),
	}
}

func (c *Crawler) Start(ctx context.Context) {
	for i := 0; i < c.Config.Workers; i++ {
		go c.worker(ctx)
	}
}

func (c *Crawler) worker(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		case path := <-c.work:
			c.handleFileChange(ctx, path)
		}
	}
}
func (c *Crawler) Sync(ctx context.Context) error {
	fileChan := Walk(c.Config.WatchPaths, c.IgnoreMap)
	for path := range fileChan {
		c.work <- path
	}
	return nil
}

func (c *Crawler) handleFileChange(ctx context.Context, path string) {
	info, err := os.Stat(path)
	if err != nil {
		return
	}
	if info.IsDir() {
		return
	}

	if info.Size() > c.Config.MaxSize {
		return
	}

	mtime := info.ModTime().Unix()
	size := info.Size()

	existing, _ := c.DB.GetFile(ctx, path)
	if existing != nil && existing.LastModified == mtime && existing.Size == size {
		return
	}

	hash, err := hashFile(path)
	if err != nil {
		return
	}

	if existing != nil && existing.Hash == hash {
		existing.LastModified = mtime
		existing.Size = size
		c.DB.UpsertFile(ctx, existing)
		return
	}

	log.Printf("Detected change: %s", path)

	f := &db.File{
		Path:         path,
		Hash:         hash,
		LastModified: mtime,
		Size:         size,
		IndexedAt:    time.Now().Unix(),
	}
	c.DB.UpsertFile(ctx, f)
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
