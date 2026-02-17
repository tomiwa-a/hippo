package ingestion

import "context"

type MarkerType string

const (
	MarkerHeader   MarkerType = "header"
	MarkerPage     MarkerType = "page"
	MarkerFunction MarkerType = "function"
)

type Marker struct {
	Type     MarkerType
	Position int
	Value    string
}

type Document struct {
	Path    string
	Content string
	Markers []Marker
}

type Chunk struct {
	ID          string
	FileID      int64
	Content     string
	SourcePath  string
	MarkerValue string
	MarkerType  MarkerType
	StartIndex  int
}

type Extractor interface {
	Extract(ctx context.Context, path string) (*Document, error)
}
