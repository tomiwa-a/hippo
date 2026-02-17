package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	_ "modernc.org/sqlite"
)

type DB struct {
	*sql.DB
}

func New(path string) (*DB, error) {
	connStr := fmt.Sprintf("%s?_pragma=journal_mode(WAL)&_pragma=synchronous(NORMAL)&_pragma=foreign_keys(ON)", path)

	sqlDB, err := sql.Open("sqlite", connStr)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	if err := sqlDB.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
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
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		file_id INTEGER NOT NULL,
		chunk_index INTEGER NOT NULL,
		content TEXT NOT NULL, 
		metadata TEXT, 
		FOREIGN KEY(file_id) REFERENCES files(id) ON DELETE CASCADE
	);

	CREATE INDEX IF NOT EXISTS idx_chunks_file_id ON chunks(file_id);
	`

	_, err := db.ExecContext(ctx, query)
	if err != nil {
		return fmt.Errorf("migration failed: %w", err)
	}

	log.Println("Database migration completed successfully.")
	return nil
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
	_, err := db.ExecContext(ctx, query, f.Path, f.Hash, f.LastModified, f.Size, f.IndexedAt)
	return err
}
