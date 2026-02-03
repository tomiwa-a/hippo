# Hippo

[![Status: Development](https://img.shields.io/badge/Status-Development-yellow)](https://github.com/tomiwaAmole/hippo)
[![Language: Go](https://img.shields.io/badge/Language-Go-00ADD8?logo=go)](https://golang.org)
[![License: MIT](https://img.shields.io/badge/License-MIT-green)](LICENSE)
[![Open Source](https://img.shields.io/badge/Open%20Source-Yes-brightgreen)](https://opensource.org)

A high-performance, local-first background daemon that gives AI agents semantic understanding of your local environment.

> **Hippo** — named after the hippocampus, the brain's memory center.

## What is Hippo?

Hippo is a knowledge engine that runs on your machine. It watches your files, builds a searchable knowledge graph, and exposes it to AI agents via the [Model Context Protocol (MCP)](https://modelcontextprotocol.io).

Think of it as **semantic RAM for your local machine**—your AI assistant can now understand your codebase, documents, and tools without external APIs or cloud storage.

## Features

- **Local-first**: All processing happens on your machine. No cloud, no API calls.
- **Multimodal ingestion**: Indexes code, documents, images, and binaries.
- **Semantic search**: Ask "where is the auth logic?" and find related files.
- **Graph awareness**: Understands imports, references, and relationships between files.
- **MCP integration**: Works with Claude Desktop, Relay, and any MCP-compliant agent.

## Getting Started

### Prerequisites

- Go 1.23+
- SQLite3

### Installation

```bash
git clone https://github.com/tomiwaAmole/hippo.git
cd hippo
go build -o hippo ./cmd
```

### Basic Usage

```bash
# Watch a directory and index files
./hippo watch ~/path/to/project

# Search your local context
./hippo query "authentication logic"

# Start the MCP server
./hippo serve
```

## Architecture

Hippo consists of three layers:

1. **The Watcher**: Detects file changes and routes them to the right processor.
2. **The Cortex**: Processes content, generates embeddings, and stores the knowledge graph.
3. **The Interface**: Exposes search and context via MCP.

See the full technical requirements for details.

## Development Roadmap

- **Phase 1 (Current)**: File discovery and metadata storage.
- **Phase 2**: Embedding generation and semantic search.
- **Phase 3**: Code analysis with Tree-sitter and graph construction.
- **Phase 4**: MCP server and agent integration.

## Future Ideas

- **Temporal context**: Query what you were working on yesterday.
- **Active profiling**: Automatic detection of frequently-edited file pairs.
- **Screen context**: Index error messages and terminal output via OCR.
- **Custom knowledge**: Ingest team documentation and external references.

## License

MIT License — see [LICENSE](LICENSE) for details.

## Author

Created by [@tomiwa_amole](https://twitter.com/tomiwa_amole)

---

**Status**: This project is in active development. Expect breaking changes.
