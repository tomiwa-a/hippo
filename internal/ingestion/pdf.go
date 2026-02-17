package ingestion

import (
	"bytes"
	"context"
	"fmt"

	"github.com/ledongthuc/pdf"
)

type PdfExtractor struct{}

func (e *PdfExtractor) Extract(ctx context.Context, path string) (string, error) {
	f, r, err := pdf.Open(path)
	if err != nil {
		return "", fmt.Errorf("failed to open pdf: %w", err)
	}
	defer f.Close()

	var buf bytes.Buffer
	b, err := r.GetPlainText()
	if err != nil {
		return "", fmt.Errorf("failed to get pdf text: %w", err)
	}

	_, err = buf.ReadFrom(b)
	if err != nil {
		return "", fmt.Errorf("failed to read pdf buffer: %w", err)
	}

	return buf.String(), nil
}
