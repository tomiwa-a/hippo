package ingestion

import (
	"bytes"
	"context"
	"fmt"

	"github.com/ledongthuc/pdf"
)

type PdfExtractor struct{}

func (e *PdfExtractor) Extract(ctx context.Context, path string) (*Document, error) {
	f, r, err := pdf.Open(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open pdf: %w", err)
	}
	defer f.Close()

	var buf bytes.Buffer
	b, err := r.GetPlainText()
	if err != nil {
		return nil, fmt.Errorf("failed to get pdf text: %w", err)
	}

	_, err = buf.ReadFrom(b)
	if err != nil {
		return nil, fmt.Errorf("failed to read pdf buffer: %w", err)
	}

	return &Document{
		Path:    path,
		Content: buf.String(),
	}, nil
}
