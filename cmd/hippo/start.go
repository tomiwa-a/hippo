package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"path/filepath"
	"strconv"
	"syscall"

	"github.com/spf13/cobra"
	"github.com/tomiwa-a/hippo/internal/config"
	"github.com/tomiwa-a/hippo/internal/crawler"
	"github.com/tomiwa-a/hippo/internal/db"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start the Hippo ingestion engine",
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()
		cfg, err := config.Load()
		if err != nil {
			log.Fatalf("Failed to load config: %v", err)
		}

		pidPath := "hippo.pid"
		if abs, err := filepath.Abs(pidPath); err == nil {
			pidPath = abs
		}

		if data, err := os.ReadFile(pidPath); err == nil {
			if pid, err := strconv.Atoi(string(data)); err == nil {
				if process, err := os.FindProcess(pid); err == nil {
					if err := process.Signal(syscall.Signal(0)); err == nil {
						log.Fatalf("Hippo is already running (PID: %d)", pid)
					}
				}
			}
		}

		if err := os.WriteFile(pidPath, []byte(fmt.Sprintf("%d", os.Getpid())), 0644); err != nil {
			log.Printf("Warning: Failed to write PID file: %v", err)
		}
		defer os.Remove(pidPath)

		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt, syscall.SIGTERM)
		go func() {
			<-c
			os.Remove(pidPath)
			os.Exit(0)
		}()

		fmt.Println("🦛 Hippo Engine Started")
		fmt.Printf("PID: %d\n", os.Getpid())
		fmt.Printf("Database Path: %s\n", cfg.DBPath)

		database, err := db.New(cfg.DBPath)
		if err != nil {
			log.Fatalf("Failed to initialize database: %v", err)
		}
		defer database.Close()

		fmt.Println("Database connected and migrated.")

		engine := crawler.New(database, cfg)
		engine.Start(ctx)

		fmt.Println("Starting Sync...")
		if err := engine.Sync(ctx); err != nil {
			log.Fatalf("Sync failed: %v", err)
		}
		fmt.Println("Initial Sync completed.")

		fmt.Println("Watching for changes...")
		if err := engine.Watch(ctx); err != nil {
			log.Fatalf("Watcher failed: %v", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}
