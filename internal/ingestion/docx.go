package ingestion

import (
	"context"
	"fmt"

	"github.com/nguyenthenguyen/docx"
)

type DocxExtractor struct{}

func (e *DocxExtractor) Extract(ctx context.Context, path string) (*Document, error) {
	r, err := docx.ReadDocxFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read docx: %w", err)
	}
	defer r.Close()

	content := r.Editable().GetContent()
	doc := &Document{
		Path:    path,
		Content: content,
	}

	// Simple heuristic: lines with fewer than 50 chars followed by newline might be headers
	// But docx XML is messy. For now, we'll keep it simple as requested "clean".
	return doc, nil
}
