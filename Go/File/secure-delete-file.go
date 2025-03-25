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
	file, err := os.OpenFile(filename, os.O_WRONLY, 0)
	if err != nil {
		// Log an error message if the file cannot be opened and exit the function.
		log.Println("Error opening file:", err)
		return
	}
	// Ensure the file is closed when the function exits.
	defer func() {
		// Attempt to close the file and log any error during closing.
		if cerr := file.Close(); cerr != nil {
			log.Println("Error closing file:", cerr)
		}
	}()

	// Retrieve file information to determine its size.
	info, err := file.Stat()
	if err != nil {
		// Log an error message if file info cannot be retrieved and exit the function.
		log.Println("Error getting file info:", err)
		return
	}
	// Get the file size in bytes.
	size := info.Size()

	// Overwrite the file contents 3 times with random data.
	for i := 0; i < 3; i++ {
		// Reset the file pointer to the beginning of the file.
		if _, err := file.Seek(0, 0); err != nil {
			// Log an error message if seeking fails and exit the function.
			log.Println("Error seeking file:", err)
			return
		}

		// Create a byte slice buffer with the same length as the file.
		randomData := make([]byte, size)

		// Fill the buffer with cryptographically secure random bytes.
		if _, err := rand.Read(randomData); err != nil {
			// Log an error message if random data generation fails and exit the function.
			log.Println("Error generating random data:", err)
			return
		}

		// Write the random data over the file's current contents.
		if _, err := file.Write(randomData); err != nil {
			// Log an error message if writing fails and exit the function.
			log.Println("Error writing random data:", err)
			return
		}

		// Flush the file system's in-memory copy to disk.
		if err := file.Sync(); err != nil {
			// Log an error message if syncing fails and exit the function.
			log.Println("Error syncing file:", err)
			return
		}
	}

	// Explicitly close the file before deletion (attempted once more in case of deferred closure errors).
	if err := file.Close(); err != nil {
		log.Println("Error closing file before deletion:", err)
	}

	// Remove the file from the file system.
	if err := os.Remove(filename); err != nil {
		// Log an error message if file removal fails.
		log.Println("Error deleting file:", err)
	}
}

func main() {
	// Define the file to be securely deleted
	filename := "/private/folder/sensitive_data.txt"
	// Attempt to securely delete the file
	secureDelete(filename)
}
