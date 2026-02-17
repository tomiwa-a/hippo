package ingestion

import (
	"bytes"
	"context"
	"fmt"
	"strings"

	"github.com/ledongthuc/pdf"
)

type PdfExtractor struct{}

func (e *PdfExtractor) Extract(ctx context.Context, path string) (*Document, error) {
	f, r, err := pdf.Open(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open pdf: %w", err)
	}
	defer f.Close()

	doc := &Document{
		Path: path,
	}

	var buf bytes.Buffer
	totalChars := 0
	numPages := r.NumPage()

	for i := 1; i <= numPages; i++ {
		p := r.Page(i)
		if p.V.IsNull() {
			continue
		}

		doc.Markers = append(doc.Markers, Marker{
			Type:     MarkerPage,
			Position: totalChars,
			Value:    fmt.Sprintf("Page %d", i),
		})

		content := p.Content()
		texts := content.Text
		if len(texts) == 0 {
			continue
		}

		// Heuristic: identify headers by larger font size
		var pageContent bytes.Buffer
		for j, text := range texts {
			if text.FontSize > 12 { // Assuming body is usually 10-12pt
				// Check if it's the start of a line/section
				isNewSection := j == 0 || (texts[j-1].Y != text.Y)
				if isNewSection && len(text.S) > 3 {
					doc.Markers = append(doc.Markers, Marker{
						Type:     MarkerHeader,
						Position: totalChars + pageContent.Len(),
						Value:    strings.TrimSpace(text.S),
					})
				}
			}
			pageContent.WriteString(text.S)
		}

		buf.Write(pageContent.Bytes())
		totalChars += pageContent.Len()
	}

	doc.Content = buf.String()
	return doc, nil
}
