package main

import (
	"crypto/rand"
	"os"
	"log"
)

// secureDelete securely deletes a file by overwriting it with random data multiple times before deletion.
func secureDelete(filename string) {
	// Open the file in write-only mode
	file, err := os.OpenFile(filename, os.O_WRONLY, 0)
	if err != nil {
		log.Println("Error opening file:", err)
	}
	defer file.Close() // Ensure the file is closed on function exit

	// Get file size
	info, err := file.Stat()
	if err != nil {
		log.Println("Error getting file info:", err)
	}
	size := info.Size()
	// Overwrite the file with random data multiple times
	for i := 0; i < 3; i++ {
		// Reset file pointer to the beginning
		_, err = file.Seek(0, 0)
		if err != nil {
			log.Println("Error seeking file:", err)
		}
		// Create a new buffer filled with random data for each pass
		randomData := make([]byte, size)
		_, err = rand.Read(randomData)
		if err != nil {
			log.Println("Error generating random data:", err)
		}
		// Write the random data over the file contents
		_, err = file.Write(randomData)
		if err != nil {
			log.Println("Error writing random data:", err)
		}
		// Force changes to be written to disk
		err = file.Sync()
		if err != nil {
			log.Println("Error syncing file:", err)
		}
	}
	// Close the file before deletion
	file.Close()
	// Remove the file after secure overwriting
	err = os.Remove(filename)
	if err != nil {
		log.Println("Error deleting file:", err)
	}
}

func main() {
	// Define the file to be securely deleted
	filename := "/private/folder/sensitive_data.txt"
	// Attempt to securely delete the file
	secureDelete(filename)
}
