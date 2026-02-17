package ingestion

import (
	"crypto/sha256"
	"encoding/hex"
	"sort"
	"strings"
)

type Chunker struct {
	Size    int
	Overlap int
}

func NewChunker(size, overlap int) *Chunker {
	return &Chunker{
		Size:    size,
		Overlap: overlap,
	}
}

func (c *Chunker) Chunk(doc *Document) []Chunk {
	if len(doc.Content) == 0 {
		return nil
	}

	var chunks []Chunk
	markers := doc.Markers
	sort.Slice(markers, func(i, j int) bool {
		return markers[i].Position < markers[j].Position
	})

	start := 0
	for start < len(doc.Content) {
		end := start + c.Size
		if end > len(doc.Content) {
			end = len(doc.Content)
		}

		bestBreak := end
		var activeMarker *Marker

		for i := range markers {
			if markers[i].Position > start && markers[i].Position <= end {
				bestBreak = markers[i].Position
				activeMarker = &markers[i]
			}
		}

		if bestBreak == end && end < len(doc.Content) {
			if idx := strings.LastIndex(doc.Content[start:end], "\n\n"); idx != -1 {
				bestBreak = start + idx + 2
			} else if idx := strings.LastIndex(doc.Content[start:end], "\n"); idx != -1 {
				bestBreak = start + idx + 1
			} else if idx := strings.LastIndex(doc.Content[start:end], ". "); idx != -1 {
				bestBreak = start + idx + 2
			}
		}

		content := strings.TrimSpace(doc.Content[start:bestBreak])
		if len(content) > 0 {
			chunk := Chunk{
				Content:    content,
				SourcePath: doc.Path,
				StartIndex: start,
			}
			if activeMarker != nil {
				chunk.MarkerValue = activeMarker.Value
				chunk.MarkerType = activeMarker.Type
			}
			h := sha256.Sum256([]byte(content + doc.Path))
			chunk.ID = hex.EncodeToString(h[:])
			chunks = append(chunks, chunk)
		}

		nextStart := bestBreak - c.Overlap
		if nextStart <= start {
			nextStart = bestBreak
		}
		start = nextStart
	}

	return chunks
}
