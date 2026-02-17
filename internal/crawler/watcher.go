package crawler

import (
	"context"
	"log"
	"os"
	"path/filepath"

	"github.com/fsnotify/fsnotify"
	"github.com/tomiwa-a/hippo/internal/db"
)

func (c *Crawler) Watch(ctx context.Context) error {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return err
	}
	defer watcher.Close()

	for _, root := range c.Config.WatchPaths {
		err = filepath.WalkDir(root, func(path string, d os.DirEntry, err error) error {
			if err != nil {
				return nil
			}
			if d.IsDir() {
				if isIgnored(path, c.Config.Ignore) {
					return filepath.SkipDir
				}
				return watcher.Add(path)
			}
			return nil
		})
		if err != nil {
			log.Printf("Error setting up watch for %s: %v", root, err)
		}
	}

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case event, ok := <-watcher.Events:
			if !ok {
				return nil
			}

			if isIgnored(event.Name, c.Config.Ignore) {
				continue
			}

			if event.Has(fsnotify.Write) || event.Has(fsnotify.Create) {
				c.handleFileChange(ctx, event.Name)
			}

			if event.Has(fsnotify.Create) {
				info, err := os.Stat(event.Name)
				if err == nil && info.IsDir() {
					watcher.Add(event.Name)
				}
			}

		case err, ok := <-watcher.Errors:
			if !ok {
				return nil
			}
			log.Printf("Watcher error: %v", err)
		}
	}
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

	log.Printf("Detected change via watcher: %s", path)

	f := &db.File{
		Path:         path,
		Hash:         hash,
		LastModified: mtime,
		Size:         size,
		IndexedAt:    info.ModTime().Unix(),
	}
	c.DB.UpsertFile(ctx, f)
}
