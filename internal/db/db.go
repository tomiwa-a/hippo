package db

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"

	_ "github.com/asg017/sqlite-vec-go-bindings/ncruces"
	_ "github.com/ncruces/go-sqlite3/driver"
	"github.com/tomiwa-a/hippo/internal/ingestion"
)

type DB struct {
	*sql.DB
}

func New(path string) (*DB, error) {
	sqlDB, err := sql.Open("sqlite3", path)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	if err := sqlDB.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	// Enable WAL mode for better concurrency
	if _, err := sqlDB.Exec("PRAGMA journal_mode=WAL;"); err != nil {
		return nil, fmt.Errorf("failed to enable WAL mode: %w", err)
	}
	// logical busy timeout (10s)
	if _, err := sqlDB.Exec("PRAGMA busy_timeout=10000;"); err != nil {
		return nil, fmt.Errorf("failed to set busy timeout: %w", err)
	}
	// synchronisation setting for WAL
	if _, err := sqlDB.Exec("PRAGMA synchronous=NORMAL;"); err != nil {
		return nil, fmt.Errorf("failed to set synchronous mode: %w", err)
	}

	db := &DB{sqlDB}
	if err := db.migrate(context.Background()); err != nil {
		sqlDB.Close()
		return nil, fmt.Errorf("failed to migrate database: %w", err)
	}

	return db, nil
}

func (db *DB) migrate(ctx context.Context) error {
	query := `
	CREATE TABLE IF NOT EXISTS files (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		path TEXT NOT NULL UNIQUE,
		hash TEXT NOT NULL,
		last_modified INTEGER NOT NULL,
		size INTEGER NOT NULL,
		indexed_at INTEGER NOT NULL
	);

	CREATE INDEX IF NOT EXISTS idx_files_path ON files(path);
	CREATE INDEX IF NOT EXISTS idx_files_hash ON files(hash);
	
	CREATE TABLE IF NOT EXISTS chunks (
		file_id INTEGER NOT NULL,
		chunk_index INTEGER NOT NULL,
		chunk_id TEXT NOT NULL, 
		content TEXT NOT NULL, 
		metadata TEXT, 
		PRIMARY KEY(file_id, chunk_index),
		FOREIGN KEY(file_id) REFERENCES files(id) ON DELETE CASCADE
	);

	CREATE VIRTUAL TABLE IF NOT EXISTS vec_chunks USING vec0(
		chunk_id TEXT PRIMARY KEY,
		embedding FLOAT[768]
	);

	CREATE INDEX IF NOT EXISTS idx_chunks_file_id ON chunks(file_id);
	CREATE INDEX IF NOT EXISTS idx_chunks_chunk_id ON chunks(chunk_id);
	`

	_, err := db.ExecContext(ctx, query)
	if err != nil {
		return fmt.Errorf("migration failed: %w", err)
	}

	log.Println("Database migration completed successfully.")
	return nil
}

func (db *DB) HasEmbedding(ctx context.Context, chunkID string) (bool, error) {
	var count int
	err := db.QueryRowContext(ctx, "SELECT count(*) FROM vec_chunks WHERE chunk_id = ?", chunkID).Scan(&count)
	return count > 0, err
}

func (db *DB) SaveChunk(ctx context.Context, c ingestion.Chunk, embedding []float32) error {
	metaJSON, err := json.Marshal(c.Meta)
	if err != nil {
		return err
	}

	// Simple retry logic for busy database
	maxRetries := 5
	for i := 0; i < maxRetries; i++ {
		tx, err := db.BeginTx(ctx, nil)
		if err != nil {
			if i == maxRetries-1 {
				return err
			}
			continue
		}

		_, err = tx.ExecContext(ctx, `
			INSERT OR REPLACE INTO chunks (file_id, chunk_index, chunk_id, content, metadata)
			VALUES (?, ?, ?, ?, ?)`,
			c.FileID, c.Index, c.ID, c.Content, string(metaJSON))
		if err != nil {
			tx.Rollback()
			if i == maxRetries-1 {
				return err
			}
			continue
		}

		if embedding != nil {
			exists, err := db.HasEmbedding(ctx, c.ID)
			if err != nil {
				tx.Rollback()
				return err
			}

			if !exists {
				blob := serializeFloat32(embedding)
				_, err = tx.ExecContext(ctx, `
					INSERT INTO vec_chunks (chunk_id, embedding)
					VALUES (?, ?)`,
					c.ID, blob)
				if err != nil {
					tx.Rollback()
					if i == maxRetries-1 {
						return err
					}
					continue
				}
			}
		}

		if err := tx.Commit(); err != nil {
			if i == maxRetries-1 {
				return err
			}
			continue
		}
		return nil
	}
	return fmt.Errorf("max retries exceeded")
}

type File struct {
	ID           int64
	Path         string
	Hash         string
	LastModified int64
	Size         int64
	IndexedAt    int64
}

func (db *DB) GetFile(ctx context.Context, path string) (*File, error) {
	var f File
	query := `SELECT id, path, hash, last_modified, size, indexed_at FROM files WHERE path = ?`
	err := db.QueryRowContext(ctx, query, path).Scan(&f.ID, &f.Path, &f.Hash, &f.LastModified, &f.Size, &f.IndexedAt)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &f, nil
}

func (db *DB) UpsertFile(ctx context.Context, f *File) error {
	query := `
	INSERT INTO files (path, hash, last_modified, size, indexed_at)
	VALUES (?, ?, ?, ?, ?)
	ON CONFLICT(path) DO UPDATE SET
		hash = excluded.hash,
		last_modified = excluded.last_modified,
		size = excluded.size,
		indexed_at = excluded.indexed_at
	`
	// Simple retry logic for busy database
	maxRetries := 5
	for i := 0; i < maxRetries; i++ {
		_, err := db.ExecContext(ctx, query, f.Path, f.Hash, f.LastModified, f.Size, f.IndexedAt)
		if err == nil {
			return nil
		}
		// If error is not relevant to locking, return immediately (simplified for now to just retry on all errors for robustness)
		// rigorous check would be scanning error string for "database is locked"
		if i == maxRetries-1 {
			return err
		}
		continue
	}
	return fmt.Errorf("max retries exceeded for upsert file")
}
