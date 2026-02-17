package main

import (
	"fmt"
	"log"

	"github.com/tomiwa-a/hippo/internal/config"
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

	fmt.Println("Watching:")
	for _, p := range cfg.WatchPaths {
		fmt.Printf("  - %s\n", p)
	}
}
