package crawler

import (
	"log"
	"os"
	"path/filepath"
	"strings"
)

func Walk(roots []string, ignores []string) <-chan string {
	out := make(chan string)

	go func() {
		defer close(out)

		for _, root := range roots {
			err := filepath.WalkDir(root, func(path string, d os.DirEntry, err error) error {
				if err != nil {
					log.Printf("Error accessing path %q: %v\n", path, err)
					return nil
				}

				if isIgnored(path, ignores) {
					if d.IsDir() {
						return filepath.SkipDir
					}
					return nil
				}

				if !d.IsDir() {
					out <- path
				}

				return nil
			})

			if err != nil {
				log.Printf("Error walking root %q: %v\n", root, err)
			}
		}
	}()

	return out
}

func isIgnored(path string, ignores []string) bool {
	for _, ignore := range ignores {
		if strings.Contains(path, ignore) {
			return true
		}
	}
	return false
}
