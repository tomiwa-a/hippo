package ingestion

import (
	"context"
	"fmt"
	"os"
)

type TextExtractor struct{}

func (e *TextExtractor) Extract(ctx context.Context, path string) (*Document, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}
	return &Document{
		Path:    path,
		Content: string(content),
	}, nil
}
