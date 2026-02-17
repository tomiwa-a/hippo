package main

import (
	"fmt"
	"log"

	"github.com/tomiwa-a/hippo/internal/config"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	fmt.Println("🦛 Hippo Engine Started")
	fmt.Printf("Database: %s\n", cfg.DBPath)
	fmt.Println("Watching:")
	for _, p := range cfg.WatchPaths {
		fmt.Printf("  - %s\n", p)
	}
	fmt.Println("Ignoring:")
	for _, i := range cfg.Ignore {
		fmt.Printf("  - %s\n", i)
	}
}
