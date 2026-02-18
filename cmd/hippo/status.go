package main

import (
	"fmt"
	"log"

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

		// Connect to DB
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

		// Styling
		redBold := color.New(color.FgRed, color.Bold).SprintFunc()

		fmt.Printf("%s\n", redBold("[HIPPO] Status Report"))
		fmt.Println("-----------------------")
		fmt.Printf("Database:      %s\n", cfg.DBPath)
		fmt.Printf("Files Indexed: %d\n", count)
		fmt.Printf("Total Size:    %.2f MB\n", float64(totalSize)/(1024*1024))
		fmt.Println("-----------------------")
	},
}

func init() {
	rootCmd.AddCommand(statusCmd)
}
