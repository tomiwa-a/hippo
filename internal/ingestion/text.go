package ingestion

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
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

		// Extract Markdown Links
		e.extractMarkdownLinks(doc)
	}

	ext := filepath.Ext(path)
	if ext == ".go" || ext == ".ts" || ext == ".tsx" {
		e.extractCodeLinks(doc)
	}

	return doc, nil
}

func (e *TextExtractor) extractMarkdownLinks(doc *Document) {
	// [[Wikilinks]]
	reWiki := regexp.MustCompile(`\[\[(.*?)\]\]`)
	matches := reWiki.FindAllStringSubmatch(doc.Content, -1)
	for _, m := range matches {
		doc.Links = append(doc.Links, Link{Target: m[1], Type: "wikilink"})
	}

	// Standard links [Text](Target)
	reLink := regexp.MustCompile(`\[.*?\]\((.*?)\)`)
	matches = reLink.FindAllStringSubmatch(doc.Content, -1)
	for _, m := range matches {
		target := m[1]
		if !strings.HasPrefix(target, "http") && !strings.HasPrefix(target, "#") {
			doc.Links = append(doc.Links, Link{Target: target, Type: "link"})
		}
	}
}

func (e *TextExtractor) extractCodeLinks(doc *Document) {
	// Simple import parsing for Go/TS
	reImport := regexp.MustCompile(`import\s+["'](.*?)["']`)
	matches := reImport.FindAllStringSubmatch(doc.Content, -1)
	for _, m := range matches {
		doc.Links = append(doc.Links, Link{Target: m[1], Type: "import"})
	}
}
