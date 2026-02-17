package ingestion

import (
	"context"
	"fmt"
	"path/filepath"
)

type Extractor interface {
	Extract(ctx context.Context, path string) (string, error)
}

type Registry struct {
	handlers map[string]Extractor
}

func NewRegistry() *Registry {
	r := &Registry{
		handlers: make(map[string]Extractor),
	}

	textHandler := &TextExtractor{}

	r.Register(".txt", textHandler)
	r.Register(".md", textHandler)
	r.Register(".go", textHandler)
	r.Register(".ts", textHandler)
	r.Register(".tsx", textHandler)

	return r
}

func (r *Registry) Register(ext string, handler Extractor) {
	r.handlers[ext] = handler
}

func (r *Registry) Extract(ctx context.Context, path string) (string, error) {
	ext := filepath.Ext(path)
	handler, ok := r.handlers[ext]
	if !ok {
		return "", fmt.Errorf("no extractor registered for extension: %s", ext)
	}
	return handler.Extract(ctx, path)
}
