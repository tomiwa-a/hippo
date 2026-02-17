package ingestion_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/tomiwa-a/hippo/internal/ingestion"
)

func TestIngestionPipeline(t *testing.T) {
	registry := ingestion.NewRegistry()
	chunker := ingestion.NewChunker(500, 100)
	ctx := context.Background()

	testFiles := []string{
		// "../../../dummy/test.md",
		// "../../../dummy/test.txt",
		"../../../dummy/test1.pdf",
	}

	for _, path := range testFiles {
		t.Run(path, func(t *testing.T) {
			doc, err := registry.Extract(ctx, path)
			if err != nil {
				t.Errorf("Failed to extract %s: %v", path, err)
				return
			}

			fmt.Printf("\n--- Document: %s ---\n", doc.Path)
			fmt.Printf("Markers found: %d\n", len(doc.Markers))
			for _, m := range doc.Markers {
				fmt.Printf("  [%s] at %d: %s\n", m.Type, m.Position, m.Value)
			}

			chunks := chunker.Chunk(doc)
			fmt.Printf("Chunks created: %d\n", len(chunks))
			for i, c := range chunks {
				fmt.Printf("\n[Chunk %d] (Meta: %v)\n%s\n", i, c.Meta, c.Content)
			}
		})
	}
}
