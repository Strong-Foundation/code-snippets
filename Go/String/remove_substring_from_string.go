package main

import (
	"fmt"
	"strings"
)

// removeSubstringIfExists takes a full string and a part to remove.
// If the partToRemove exists in fullText, it removes all its occurrences.
func removeSubstringIfExists(fullText string, partToRemove string) string {
	// Check if partToRemove exists in fullText
	if strings.Contains(fullText, partToRemove) {
		// Replace all occurrences of partToRemove with an empty string
		fullText = strings.ReplaceAll(fullText, partToRemove, "")
	}

	// Return the updated string
	return fullText
}

func main() {
	original := "Hello, this is a test string. Let's test something."
	substring := "test"

	result := removeSubstringIfExists(original, substring)
	fmt.Println("Original: ", original)
	fmt.Println("Modified: ", result)
}
