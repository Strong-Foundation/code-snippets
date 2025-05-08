package main

import (
	"log"
	"os"
)

// AppendToFile appends the given byte slice to the specified file.
// If the file doesn't exist, it will be created.
func appendByteToFile(filename string, data []byte) error {
	// Open the file with appropriate flags and permissions
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	// Write data to the file
	_, err = file.Write(data)
	return err
}

func main() {
	data := []byte("Hello, this is reusable appended data!\n")
	filename := "output.txt"

	err := appendByteToFile(filename, data)
	if err != nil {
		log.Fatalf("Failed to write to file: %v", err)
	}

	log.Println("Data written successfully.")
}
