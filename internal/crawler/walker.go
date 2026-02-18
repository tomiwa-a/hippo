package crawler

import (
	"log"
	"os"
	"path/filepath"

	gitignore "github.com/sabhiram/go-gitignore"
)

func Walk(roots []string, gi *gitignore.GitIgnore) <-chan string {
	out := make(chan string)

	go func() {
		defer close(out)

		for _, root := range roots {
			err := filepath.WalkDir(root, func(path string, d os.DirEntry, err error) error {
				if err != nil {
					log.Printf("Error accessing path %q: %v\n", path, err)
					return nil
				}

				if gi.MatchesPath(path) {
					if d.IsDir() {
						log.Printf("DEBUG: Skipping directory %s", path)
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
