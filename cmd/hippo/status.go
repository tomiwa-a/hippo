package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/tomiwa-a/hippo/internal/config"
	"github.com/tomiwa-a/hippo/internal/db"
)

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Check the health and stats of the Hippo engine",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := config.Load()
		if err != nil {
			log.Fatalf("Failed to load config: %v", err)
		}

		database, err := db.New(cfg.DBPath)
		if err != nil {
			color.Red("Error: Database not accessible (%v)", err)
			return
		}
		defer database.Close()

		var count int
		if err := database.QueryRow("SELECT count(*) FROM files").Scan(&count); err != nil {
			color.Red("Error querying file count: %v", err)
			return
		}

		var totalSize int64
		if err := database.QueryRow("SELECT COALESCE(SUM(size), 0) FROM files").Scan(&totalSize); err != nil {
			color.Red("Error querying total size: %v", err)
			return
		}

		// Get DB File Size (Index Size)
		var dbSize int64
		if info, err := os.Stat(cfg.DBPath); err == nil {
			dbSize = info.Size()
		}

		// Check Memory Usage (if PID file exists)
		var memUsage string = "Not Running"
		pidData, err := os.ReadFile("hippo.pid")
		if err == nil {
			pidStr := string(pidData)
			// Run ps -o rss= -p <PID>
			// Output is in KB
			cmd := exec.Command("ps", "-o", "rss=", "-p", pidStr)
			out, err := cmd.Output()
			if err == nil {
				// Parse KB
				var kb int64
				fmt.Sscanf(strings.TrimSpace(string(out)), "%d", &kb)
				memUsage = fmt.Sprintf("%.2f MB", float64(kb)/1024.0)
			}
		}

		redBold := color.New(color.FgRed, color.Bold).SprintFunc()

		fmt.Printf("%s\n", redBold("[HIPPO] Status Report"))
		fmt.Println("-----------------------")
		fmt.Printf("Files Indexed: %d\n", count)
		fmt.Printf("Content Size:  %.2f MB\n", float64(totalSize)/(1024*1024))
		fmt.Printf("Index Size:    %.2f MB\n", float64(dbSize)/(1024*1024))
		fmt.Printf("Memory Usage:  %s\n", memUsage)
		fmt.Println("-----------------------")
	},
}

func init() {
	rootCmd.AddCommand(statusCmd)
}
