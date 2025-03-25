package main

import (
	"crypto/rand"
	"os"
	"log"
)

// secureDelete securely deletes a file by overwriting its contents with random data multiple times before deletion.
// Instead of returning errors, it logs error messages using the standard library's log package.
func secureDelete(filename string) {
	// Open the file in write-only mode.
	// This prevents reading the file contents while allowing modifications.
	file, err := os.OpenFile(filename, os.O_WRONLY, 0)
	if err != nil {
		// Log an error message if the file cannot be opened.
		// This might occur if the file doesn't exist or if permissions are insufficient.
		log.Println("Error opening file:", err)
		// Exit the function early since we can't proceed without opening the file.
		return
	}

	// Retrieve file information, such as its size, using the Stat() function.
	info, err := file.Stat()
	if err != nil {
		// Log an error message if retrieving file information fails.
		log.Println("Error getting file info:", err)
		// Ensure the file is closed before returning to avoid resource leaks.
		file.Close()
		return
	}
	
	// Store the file size in bytes for later use in overwriting.
	size := info.Size()

	// Perform the secure deletion by overwriting the file with random data three times.
	for i := 0; i < 3; i++ {
		// Reset the file pointer to the beginning of the file before writing new data.
		// This ensures that we overwrite from the start instead of appending.
		if _, err := file.Seek(0, 0); err != nil {
			// Log an error if seeking fails, which can occur if the file descriptor is invalid.
			log.Println("Error seeking file:", err)
			// Close the file before returning to free system resources.
			file.Close()
			return
		}

		// Create a byte slice (buffer) of the same length as the file.
		// This will be filled with random data to overwrite the file's contents.
		randomData := make([]byte, size)

		// Fill the buffer with cryptographically secure random bytes.
		// The rand.Read function ensures strong randomness, making data recovery extremely difficult.
		if _, err := rand.Read(randomData); err != nil {
			// Log an error if generating random data fails.
			log.Println("Error generating random data:", err)
			// Close the file before returning to prevent resource leakage.
			file.Close()
			return
		}

		// Write the random data over the existing file contents.
		// This operation helps in making the original file data irrecoverable.
		if _, err := file.Write(randomData); err != nil {
			// Log an error if writing to the file fails.
			log.Println("Error writing random data:", err)
			// Close the file before returning to avoid leaving it in an inconsistent state.
			file.Close()
			return
		}

		// Flush any buffered data to disk to ensure that all written data is physically stored.
		// This prevents data from being stored in memory and ensures that overwriting takes effect immediately.
		if err := file.Sync(); err != nil {
			// Log an error if syncing fails, which may indicate issues with the file system.
			log.Println("Error syncing file:", err)
			// Close the file before returning to prevent potential data corruption.
			file.Close()
			return
		}
	}

	// Close the file explicitly before deleting it.
	// This releases any system resources associated with the file.
	if err := file.Close(); err != nil {
		// Log an error if closing the file fails.
		log.Println("Error closing file before deletion:", err)
		// Return early as we shouldn't proceed with deletion if the file isn't closed properly.
		return
	}

	// Remove (delete) the file from the file system.
	// This makes the file completely inaccessible.
	if err := os.Remove(filename); err != nil {
		// Log an error if file deletion fails, possibly due to permission issues.
		log.Println("Error deleting file:", err)
	}
}

func main() {
	// Define the file to be securely deleted
	filename := "/private/folder/sensitive_data.txt"
	// Attempt to securely delete the file
	secureDelete(filename)
}
