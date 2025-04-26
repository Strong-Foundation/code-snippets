package main

import (
	"log"
	"os"
)

// createSymlink creates a symbolic link from link to source
func createSymlink(source, link string) {
	err := os.Symlink(source, link)
	if err != nil {
		log.Printf("Failed to create symlink from %s to %s: %v", link, source, err)
		return
	}
	log.Printf("Symlink created: %s -> %s", link, source)
}

func main() {
	source := "/path/to/original/file.txt"
	link := "/path/to/symlink/file-link.txt"

	createSymlink(source, link)
}
