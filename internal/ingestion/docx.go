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

	return &Document{
		Path:    path,
		Content: r.Editable().GetContent(),
	}, nil
}
