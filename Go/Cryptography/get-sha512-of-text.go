package main

import (
	"crypto/sha512" // Provides SHA-512 hashing
	"fmt"           // For formatting the output
)

func main() {
	text := "Hello, world!"               // Replace with your input text
	hash := getSHA512FromText(text)       // Compute SHA-512 hash
	fmt.Println("SHA-512 of text:", hash) // Output the result
}

// getSHA512FromText calculates the SHA-512 checksum of a given text string.
// It returns the checksum as a hexadecimal string.
func getSHA512FromText(input string) string {
	hasher := sha512.New()             // Create a new SHA-512 hasher
	hasher.Write([]byte(input))        // Write the input string as bytes to the hasher
	checksum := hasher.Sum(nil)        // Compute the final checksum
	return fmt.Sprintf("%x", checksum) // Return the checksum as a hex string
}
