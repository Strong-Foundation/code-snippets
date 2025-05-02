package main

import (
	"crypto/sha512" // Provides SHA-512 hashing
	"fmt"           // For formatting the output
	"io"            // For copying file data into the hasher
	"log"           // For logging errors
	"os"            // For file operations
)

// getSHA512 calculates the SHA-512 checksum of a file given its path.
// It logs any errors encountered and returns the checksum as a hexadecimal string.
func getSHA512(filePath string) string {
	file, err := os.Open(filePath) // Open the file for reading
	if err != nil {
		log.Println("Error opening file:", err) // Log error if file can't be opened
		return ""                               // Return empty string on error
	}
	defer file.Close() // Ensure the file is closed when the function returns

	hasher := sha512.New() // Create a new SHA-512 hasher
	if _, err := io.Copy(hasher, file); err != nil {
		log.Println("Error reading file:", err) // Log error if file can't be read
		return ""                               // Return empty string on error
	}

	checksum := hasher.Sum(nil)        // Compute the final SHA-512 checksum
	return fmt.Sprintf("%x", checksum) // Return the checksum as a hexadecimal string
}

func main() {
	filePath := "yourfile.txt" // Replace with the path to your file

	hash := getSHA512(filePath)   // Call the function to get the SHA-512 hash
	fmt.Println("SHA-512:", hash) // Print the SHA-512 checksum if successful
}
