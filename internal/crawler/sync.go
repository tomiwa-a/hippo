package crawler

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	gitignore "github.com/sabhiram/go-gitignore"
	"github.com/tomiwa-a/hippo/internal/config"
	"github.com/tomiwa-a/hippo/internal/db"
	"github.com/tomiwa-a/hippo/internal/embedding"
	"github.com/tomiwa-a/hippo/internal/ingestion"
)

type Crawler struct {
	DB        *db.DB
	Config    *config.Config
	IgnoreMap *gitignore.GitIgnore
	work      chan string
	registry  *ingestion.Registry
	chunker   *ingestion.Chunker
	embedder  embedding.Embedder
}

func New(database *db.DB, cfg *config.Config) *Crawler {
	gi := gitignore.CompileIgnoreLines(cfg.Ignore...)

	registry := ingestion.NewRegistry()
	chunker := ingestion.NewChunker(1000, 200) // 1000 chars, 200 overlap

	var embedder embedding.Embedder
	if cfg.Embedding.Provider == "ollama" {
		embedder = embedding.NewOllamaEmbedder(cfg.Embedding.BaseURL, cfg.Embedding.Model)
	}

	return &Crawler{
		DB:        database,
		Config:    cfg,
		IgnoreMap: gi,
		work:      make(chan string, 1000),
		registry:  registry,
		chunker:   chunker,
		embedder:  embedder,
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
	// Wait a bit or use a more robust way to signal end of work if workers are concurrent.
	// For simplicity in this local-first tool, we'll just run a resolution at the end.
	if err := c.DB.UpdateResolvedLinks(ctx); err != nil {
		log.Printf("Link resolution failed: %v", err)
	}
	return nil
}

func (c *Crawler) handleFileChange(ctx context.Context, path string) {

	// Calculate relative path for logging and DB
	relPath := path
	for _, root := range c.Config.WatchPaths {
		if strings.HasPrefix(path, root) {
			if rel, err := filepath.Rel(root, path); err == nil {
				relPath = rel
				break
			}
		}
	}

	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		// File deleted or renamed (old path)
		if err := c.DB.DeleteFile(ctx, path); err != nil {
			log.Printf("Failed to delete file %s: %v", relPath, err)
		} else {
			log.Printf("Deleted/Pruned: %s", relPath)
		}
		return
	}
	if err != nil {
		log.Printf("Error accessing file %s: %v", relPath, err)
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

	existing, err := c.DB.GetFile(ctx, path)
	if err != nil {
		log.Printf("Error getting file %s: %v\n", relPath, err)
	}

	if existing != nil && existing.LastModified == mtime && existing.Size == size {
		return
	}

	log.Printf("Processing: %s", relPath)
	doc, err := c.registry.Extract(ctx, path)
	if err != nil {
		log.Printf("Extraction failed for %s: %v", relPath, err)
		return
	}

	// Resolve relative links
	for i, l := range doc.Links {
		if strings.HasPrefix(l.Target, ".") {
			absTarget := filepath.Join(filepath.Dir(path), l.Target)
			doc.Links[i].Target = absTarget
		}
	}
	//

	// Compute SHA256 hash of content
	hash := sha256.Sum256([]byte(doc.Content))
	hashStr := hex.EncodeToString(hash[:])

	// Hash check optimization
	if existing != nil && existing.Hash == hashStr {
		f := &db.File{
			ID:           existing.ID,
			Path:         path,
			RelativePath: relPath, // Use calculated relative path
			Hash:         hashStr,
			LastModified: mtime,
			Size:         size,
			IndexedAt:    time.Now().Unix(),
		}

		if err := c.DB.UpsertFile(ctx, f); err != nil {
			log.Printf("Failed to update file metadata %s: %v", relPath, err)
		}
		return
	}

	f := &db.File{
		Path:         path,
		RelativePath: relPath,
		Hash:         hashStr,
		LastModified: mtime,
		Size:         size,
		IndexedAt:    time.Now().Unix(),
	}

	if existing != nil {
		f.ID = existing.ID
	}

	// First upsert the file to get/ensure ID
	if err := c.DB.UpsertFile(ctx, f); err != nil {
		log.Printf("Failed to upsert file %s: %v", path, err)
		return
	}

	if f.ID == 0 {
		savedFile, _ := c.DB.GetFile(ctx, path)
		f.ID = savedFile.ID
	}

	// Save Knowledge Graph Links
	if len(doc.Links) > 0 {
		if err := c.DB.SaveLinks(ctx, f.ID, doc.Links); err != nil {
			log.Printf("Failed to save links for %s: %v", relPath, err)
		}
	}

	// Then process chunks
	chunks := c.chunker.Chunk(doc)

	for _, chunk := range chunks {
		chunk.FileID = f.ID

		hasEmbedding, _ := c.DB.HasEmbedding(ctx, chunk.ID)

		var vec []float32
		if !hasEmbedding && c.embedder != nil {
			v, err := c.embedder.Embed(ctx, chunk.Content)
			if err != nil {
				log.Printf("Embedding failed for chunk %s: %v", chunk.ID, err)
				continue
			}
			vec = v
		}

		if err := c.DB.SaveChunk(ctx, chunk, vec); err != nil {
			log.Printf("Failed to save chunk %s: %v", chunk.ID, err)
		}
	}
	log.Printf("Indexed: %s", relPath)
}
