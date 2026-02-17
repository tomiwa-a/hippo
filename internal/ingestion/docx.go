package ingestion

import (
	"context"
	"fmt"

	"github.com/nguyenthenguyen/docx"
)

type DocxExtractor struct{}

func (e *DocxExtractor) Extract(ctx context.Context, path string) (string, error) {
	r, err := docx.ReadDocxFile(path)
	if err != nil {
		return "", fmt.Errorf("failed to read docx: %w", err)
	}
	defer r.Close()

	return r.Editable().GetContent(), nil
}
