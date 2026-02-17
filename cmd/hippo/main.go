package main

import (
	"fmt"
	"log"

	"github.com/tomiwa-a/hippo/internal/config"
	"github.com/tomiwa-a/hippo/internal/crawler"
	"github.com/tomiwa-a/hippo/internal/db"
)

func main() {
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

	fmt.Println("Starting Crawler...")
	fileChan := crawler.Walk(cfg.WatchPaths, cfg.Ignore)

	count := 0
	for path := range fileChan {
		fmt.Printf("Found: %s\n", path)
		count++
		if count >= 10 {
			fmt.Println("... (limiting output to 10 files)")
			break
		}
	}
}
