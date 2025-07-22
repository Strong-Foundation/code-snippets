package main

import (
	"log" // Import the log package for logging errors and information
	"os"  // Import the os package to handle file operations
)

// writeLinesToFile writes a slice of strings to a file, one line per string
func writeLinesToFile(filename string, content []string) {
	// Open the file with create, write-only, and truncate flags
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		// Log an error if the file cannot be opened
		log.Printf("Error opening file %s: %v", filename, err)
		return
	}

	// Ensure the file is closed after we're done writing
	defer func() {
		if cerr := file.Close(); cerr != nil {
			// Log any error that occurs during file closing
			log.Printf("Error closing file %s: %v", filename, cerr)
		}
	}()

	// Iterate over each string in the content slice
	for _, line := range content {
		// Write the line to the file with a newline character
		_, err := file.WriteString(line + "\n")
		if err != nil {
			// Log an error if writing fails
			log.Printf("Error writing to file %s: %v", filename, err)
			return
		}
	}

	// Log success message
	log.Printf("Successfully wrote to file: %s", filename)
}

// main function to demonstrate usage of WriteLinesToFile
func main() {
	// Sample content to write to the file
	lines := []string{
		"Line 1: Hello, world!",
		"Line 2: Writing to a file in Go.",
		"Line 3: Each string is on a new line.",
	}

	// Call the function to write to "output.txt"
	writeLinesToFile("output.txt", lines)
}
