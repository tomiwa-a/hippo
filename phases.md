# Hippo Phase 2: Core Functionality Implementation

This document outlines the 20-step plan to build the local-first semantic knowledge engine backend, structured into 3 key phases: **The Crawler**, **Ingestion & Mapping**, and **Querying**.

## Phase 2.1: The Crawler (Background Service)

- [x] 1. **Project Skeleton & Config**: Set up `cmd/hippo` with Viper/Cobra. Implement priority loading for a local `hippo.yml` to define inclusion/exclusion rules and focused index paths.
- [x] 2. **File System Database**: Create the `files` table in SQLite to track file paths, modification times (`mtime`), and hashes to detect changes.
- [x] 3. **Recursive Walker**: Implement a performant directory walker that respects `.gitignore` and user-defined exclusion patterns.
- [x] 4. **Change Detection Logic**: Build the logic to compare current file stats against the DB to identify creates, updates, and deletes.
- [x] 5. **Real-Time Watcher**: Integrate `fsnotify` to listen for OS-level file events and trigger immediate processing for active directories.
- [x] 6. **Concurrency Control**: Implement a worker pool pattern to handle file crawling without overwhelming file descriptors or CPU.

## Phase 2.2: Ingestion & Mapping (The "Brain")

> [!TIP]
> **Strategy**: Build the extraction layer as a decoupled, modular internal package. This ensures Hippo remains clean and allows for future open-sourcing of the "Universal Document Extractor."

- [x] 7. **Text Extraction Interface**: Create an interface for extracting text from different file types (Markdown, PDF, Code). (`.go`, `.ts`, `.py`) and markdown.
- [x] 8. **Rich Document Support**: Integrate libraries to parse unstructured text from PDFs and Office documents.
- [x] 9. **Chunking Engine**: Implement a sliding window chunker with configurable overlap to preserve context across boundaries.
- [x] 10. **Metadata Extractor**: Extract semantic metadata (author, created date, meaningful title) to enrich the vector payload.
- [x] 11. **Local Embedding Service**: Create an abstraction for embedding providers (Start with ONNX Runtime/`all-MiniLM-L6-v2` or Ollama).
- [x] 12. **Vector Database Setup**: Initialize `sqlite-vec` tables for storing high-dimensional embeddings.
- [x] 13. **Content Hashing & Robustness**: SHA-256 for dedup, WAL mode for concurrency, and retry logic for stability.

## Phase 2.3: Vector Search Engine [COMPLETED]

- [x] 14. **Vector Search Implementation**: Implement KNN search using `sqlite-vec`.

## Phase 2.4: The CLI Experience [COMPLETED]

- [x] 15. **Command Structure**: Refactor `main.go` to use `cobra` subcommands.
- [x] 16. **`hippo start`**: The daemon process with PID singleton enforcement.
- [x] 17. **`hippo query`**: Semantic search interface with snippets.
- [x] 18. **`hippo status`**: Health check, DB stats, and memory monitoring.
- [x] 19. **`hippo stop`**: Clean daemon termination.

## Phase 3: Advanced Intelligence & Interface

- [ ] 17. **Knowledge Graph Linking**: Create a simple link map (e.g., referencing other files via import statements or wikilinks).
- [ ] 18. **Hybrid Search Logic**: Combine vector results with FTS5 keyword search.
- [ ] 19. **Result Reranking**: Boost results based on recency or exact matches.
- [ ] 20. **HTTP API Server**: Build a lightweight REST API (Delayed).
- [ ] 21. **Cross-Platform Build**: Configure build pipeline.
