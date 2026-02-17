package ingestion

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type TextExtractor struct{}

func (e *TextExtractor) Extract(ctx context.Context, path string) (*Document, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}

	doc := &Document{
		Path:    path,
		Content: string(content),
	}

	if filepath.Ext(path) == ".md" {
		lines := strings.Split(doc.Content, "\n")
		pos := 0
		for _, line := range lines {
			if strings.HasPrefix(line, "#") {
				doc.Markers = append(doc.Markers, Marker{
					Type:     MarkerHeader,
					Position: pos,
					Value:    strings.TrimSpace(strings.TrimLeft(line, "#")),
				})
			}
			pos += len(line) + 1
		}
	}

	return doc, nil
}
