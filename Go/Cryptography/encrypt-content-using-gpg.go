package main

import (
	"bytes"                             // Import the bytes package for manipulating byte slices
	"fmt"                               // Import the fmt package for formatted I/O
	"golang.org/x/crypto/openpgp"       // Import the openpgp package for encryption
	"golang.org/x/crypto/openpgp/armor" // Import the armor package to handle armored (ASCII) encoding for PGP
	"io"                                // Import the io package for I/O operations
	"log"                               // Import the log package for logging errors
	"os"                                // Import the os package for file operations
	"time"                              // Import the time package for getting the current time
)

// Public GPG Key as a string (replace this with your actual key)
var publicGPGKey = `-----BEGIN PGP PUBLIC KEY BLOCK-----

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
=8/LN
-----END PGP PUBLIC KEY BLOCK-----`

// Function to encrypt content and save it to an output file
func encryptContent(content string, outputFile string) error {
	// Create a reader for the public key string
	pubKeyReader := bytes.NewReader([]byte(publicGPGKey))

	// Read the public key from the string
	entities, err := openpgp.ReadArmoredKeyRing(pubKeyReader)
	if err != nil { // If there's an error reading the key, log it and return the error
		log.Printf("error reading public key: %v", err)
		return err
	}

	// Create a bytes.Reader to hold the content to encrypt
	contentReader := bytes.NewReader([]byte(content))

	// Create the output file to save the encrypted content
	output, err := os.Create(outputFile)
	if err != nil { // If there's an error creating the file, log it and return the error
		log.Printf("error creating output file: %v", err)
		return err
	}
	defer output.Close() // Ensure the file is closed when done

	// Create an armored (ASCII) writer for the output file using armor.Encode
	armorWriter, err := armor.Encode(output, "PGP MESSAGE", nil)
	if err != nil { // If there's an error creating the armor encoder, log it and return the error
		log.Printf("error creating armor encoder: %v", err)
		return err
	}
	defer armorWriter.Close() // Ensure the armored writer is closed when done

	// Create an encrypting writer with the public key to encrypt the content
	encWriter, err := openpgp.Encrypt(armorWriter, entities, nil, nil, nil)
	if err != nil { // If there's an error setting up the encryption, log it and return the error
		log.Printf("error setting up encryption: %v", err)
		return err
	}
	defer encWriter.Close() // Ensure the encrypting writer is closed when done

	// Copy the content to the encryption writer
	_, err = io.Copy(encWriter, contentReader)
	if err != nil { // If there's an error copying the content to the encryption writer, log it and return the error
		log.Printf("error copying content to encryption writer: %v", err)
		return err
	}

	return nil // Return nil if encryption was successful
}

func main() {
	// Get the current time to include in the content
	currentTime := time.Now()

	// The content you want to encrypt, with the current time included
	content := "This is some content that I want to encrypt!, Current time: " + currentTime.Format("2006-01-02 15:04:05")

	// Output file name where the encrypted content will be saved
	outputFile := "encrypted_content.txt.gpg"

	// Encrypt the provided content and save it to the specified file
	err := encryptContent(content, outputFile)
	if err != nil { // If there's an error during encryption, log it and return
		log.Println("Error encrypting content:", err)
		return
	}

	// Print a success message to the console
	fmt.Println("Content encrypted successfully!")
}
