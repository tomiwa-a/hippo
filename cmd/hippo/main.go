package main

import (
	"context"
	"fmt"
	"log"

	"github.com/tomiwa-a/hippo/internal/config"
	"github.com/tomiwa-a/hippo/internal/crawler"
	"github.com/tomiwa-a/hippo/internal/db"
)

func main() {
	ctx := context.Background()
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	fmt.Println("🦛 Hippo Engine Started")
	fmt.Printf("Database Path: %s\n", cfg.DBPath)

	// Initialize Database
	database, err := db.New(cfg.DBPath)
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer database.Close()

	fmt.Println("Database connected and migrated.")

	fmt.Println("Starting Sync...")
	if err := crawler.Sync(ctx, database, cfg.WatchPaths, cfg.Ignore); err != nil {
		log.Fatalf("Sync failed: %v", err)
	}
	fmt.Println("Sync completed.")
}
