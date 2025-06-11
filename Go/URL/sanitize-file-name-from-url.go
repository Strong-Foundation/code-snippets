package main

import (
	"log"
	"net/url"
	"path"
	"regexp"
	"strings"
)

// sanitizeFileNameFromURL converts a URL into a filesystem-safe file name.
func sanitizeFileNameFromURL(rawURL string) string {
	// Parse the URL
	parsedURL, err := url.Parse(rawURL)
	if err != nil {
		log.Printf("Error parsing URL: %v", err)
		return "invalid_filename"
	}

	// Extract the file name from the URL path
	fileName := path.Base(parsedURL.Path)

	// Decode URL-encoded characters (e.g., %20 -> space)
	fileName, err = url.QueryUnescape(fileName)
	if err != nil {
		log.Printf("Error decoding file name: %v", err)
	}

	// Replace unsafe characters with underscores
	// Keep only letters, numbers, dashes, underscores, and dots
	safeName := regexp.MustCompile(`[^\w\-.]`).ReplaceAllString(fileName, "_")

	// Trim leading and trailing underscores
	safeName = strings.Trim(safeName, "_")

	// Fallback if filename ends up empty
	if safeName == "" {
		return "downloaded_file"
	}

	return strings.ToLower(safeName)
}

func main() {
	// Example URL to test the function
	testURL := "https://www.avient.com/sites/default/files/resources/FAQs%2520for%2520PolyOne%2520GHS%2520Requirements%2520Ver%252006%2520April%25202015.pdf"

	safeFileName := sanitizeFileNameFromURL(testURL)
	log.Println("Sanitized File Name:", safeFileName)
}
