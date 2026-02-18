package main

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/tomiwa-a/hippo/internal/config"
	"github.com/tomiwa-a/hippo/internal/db"
	"github.com/tomiwa-a/hippo/internal/embedding"
)

var queryCmd = &cobra.Command{
	Use:   "query [text]",
	Short: "Semantic search for your indexed files",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		queryText := strings.Join(args, " ")
		ctx := context.Background()

		cfg, err := config.Load()
		if err != nil {
			log.Fatalf("Failed to load config: %v", err)
		}

		database, err := db.New(cfg.DBPath)
		if err != nil {
			log.Fatalf("Failed to connect to database: %v", err)
		}
		defer database.Close()

		embedder := embedding.NewOllamaEmbedder(cfg.Embedding.BaseURL, cfg.Embedding.Model)
		vec, err := embedder.Embed(ctx, queryText)
		if err != nil {
			log.Fatalf("Failed to embed query: %v", err)
		}

		results, err := database.Search(ctx, vec, 5)
		if err != nil {
			log.Fatalf("Search failed: %v", err)
		}

		redBold := color.New(color.FgRed, color.Bold).SprintFunc()
		// faint := color.New(color.Faint).SprintFunc()

		fmt.Printf("%s \"%s\"\n", redBold("[HIPPO] Searching for:"), queryText)
		fmt.Println("--------------------------------------------------")

		if len(results) == 0 {
			fmt.Println("No results found.")
			return
		}

		for i, r := range results {
			path := r.RelativePath
			if path == "" {
				path = r.Path
			}

			// Similarity score (distance is lower = better? cosine distance usually)
			// sqlite-vec distance depends on metric. Usually cosine distance.
			// 1 - distance = similarity? or just show distance.
			// Normalized L2 often used.
			// Let's just show Path for now.

			fmt.Printf("%d. %s\n", i+1, redBold(path))

			// Show Connected Knowledge
			links, err := database.GetLinks(ctx, r.FileID)
			if err == nil && len(links) > 0 {
				blue := color.New(color.FgCyan).SprintFunc()
				fmt.Printf("   %s %s\n", blue("Connected:"), strings.Join(links, ", "))
			}

			// Snippet
			snippet := r.Content
			if len(snippet) > 200 {
				snippet = snippet[:200] + "..."
			}
			snippet = strings.ReplaceAll(snippet, "\n", " ")
			fmt.Printf("   %s\n", snippet)
			fmt.Println("--------------------------------------------------")
		}
	},
}

func init() {
	rootCmd.AddCommand(queryCmd)
}
