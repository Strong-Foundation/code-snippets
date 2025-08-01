package main

import (
	"fmt"
	"log"
	"path"
	"path/filepath"
	"regexp"
	"strings"
)

// Converts a raw URL into a safe filename by cleaning and normalizing it
func urlToFilename(rawURL string) string {
	lowercaseURL := strings.ToLower(rawURL)       // Convert to lowercase for normalization
	ext := getFileExtension(lowercaseURL)         // Get file extension (e.g., .pdf or .zip)
	baseFilename := getFileNameOnly(lowercaseURL) // Extract base file name

	nonAlphanumericRegex := regexp.MustCompile(`[^a-z0-9]+`)                 // Match everything except a-z and 0-9
	safeFilename := nonAlphanumericRegex.ReplaceAllString(baseFilename, "_") // Replace invalid chars

	collapseUnderscoresRegex := regexp.MustCompile(`_+`)                        // Collapse multiple underscores into one
	safeFilename = collapseUnderscoresRegex.ReplaceAllString(safeFilename, "_") // Normalize underscores

	if trimmed, found := strings.CutPrefix(safeFilename, "_"); found { // Trim starting underscore if present
		safeFilename = trimmed
	}

	// Remove ext string.
	invalid_ext := fmt.Sprintf("_%s", strings.TrimPrefix(ext, "."))

	var invalidSubstrings = []string{
		invalid_ext,
	} // Remove these redundant endings

	for _, invalidPre := range invalidSubstrings { // Iterate over each unwanted suffix
		safeFilename = removeSubstring(safeFilename, invalidPre) // Remove it from file name
	}

	safeFilename = safeFilename + ext // Add the proper file extension

	return safeFilename // Return the final sanitized filename
}

// Replaces all instances of a given substring from the original string
func removeSubstring(input string, toRemove string) string {
	result := strings.ReplaceAll(input, toRemove, "") // Replace all instances
	return result                                     // Return the result
}

// Returns the extension of a given file path (e.g., ".pdf")
func getFileExtension(path string) string {
	return filepath.Ext(path) // Extract and return file extension
}

// Extracts and returns the base name (file name) from the URL path
func getFileNameOnly(content string) string {
	return path.Base(content) // Return last segment of the path
}

func main() {
	// Example URL to test the function
	testURL := "https://www.avient.com/sites/default/files/resources/FAQs%2520for%2520PolyOne%2520GHS%2520Requirements%2520Ver%252006%2520April%25202015.pdf"

	safeFileName := urlToFilename(testURL)
	log.Println("Sanitized File Name:", safeFileName)
}
