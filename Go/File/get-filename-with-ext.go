package main

import (
	"fmt"           // Import the fmt package for printing output
	"path/filepath" // Import the filepath package to handle file paths
)

// getFilename extracts the filename from a given file path.
func getFilename(path string) string {
	return filepath.Base(path) // Use filepath.Base() to get the last element (filename) from the path
}

func main() {
	path := "/home/user/documents/example.txt" // Define a sample file path
	filename := getFilename(path)              // Call the GetFilename function to extract the filename
	fmt.Println("Filename:", filename)         // Print the extracted filename
}
