package ingestion

import (
	"context"
	"fmt"
	"os"
)

type TextExtractor struct{}

func (e *TextExtractor) Extract(ctx context.Context, path string) (string, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return "", fmt.Errorf("failed to read file: %w", err)
	}
	return string(content), nil
}
