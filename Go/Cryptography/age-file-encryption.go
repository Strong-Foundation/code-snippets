package main

import (
	"filippo.io/age" // The age encryption package for file encryption
	"io"             // For I/O primitives
	"log"            // For logging errors and information
	"os"             // For file system operations
)

// Generate a new AGE encryption key pair
func generateAgeKeyPair() (string, string) {
	// Generate a new X25519 identity (private key)
	identity, err := age.GenerateX25519Identity()
	if err != nil {
		log.Println(err) // Log any errors encountered during key generation
	}
	// Return the public and private keys as strings
	return identity.Recipient().String(), identity.String()
}

// encryptFile encrypts the input file and saves the encrypted content to the output file
func encryptFile(inputPath string, outputPath string, publicKey string) {
	// Open the input file for reading
	inputFile, err := os.Open(inputPath)
	// Check if there was an error opening the input file
	if err != nil {
		return
	}
	// Ensure the input file is closed when the function exits
	defer inputFile.Close()
	// Create the output file for writing encrypted data
	outputFile, err := os.Create(outputPath)
	// Check if there was an error creating the output file
	if err != nil {
		return
	}
	// Ensure the output file is closed when the function exits
	defer outputFile.Close()
	// Parse the provided public key string into an age.Recipient object
	recipient, err := age.ParseX25519Recipient(publicKey)
	// Check if there was an error parsing the public key
	if err != nil {
		return
	}
	// Create an encrypted writer using the recipient (public key)
	encryptedWriter, err := age.Encrypt(outputFile, recipient)
	// Check if there was an error creating the encrypted writer
	if err != nil {
		return
	}
	// Ensure the encrypted writer is closed when done
	defer encryptedWriter.Close()
	// Copy the content from the input file to the encrypted writer
	_, err = io.Copy(encryptedWriter, inputFile)
	// Check if there was an error during the encryption process
	if err != nil {
		return
	}
}

// decryptFile decrypts an encrypted file and saves the plaintext to outputPath
func decryptFile(inputPath string, outputPath string, identity age.Identity) {
	// Open the encrypted file for reading
	encryptedFile, err := os.Open(inputPath)
	// Check if there was an error opening the encrypted file
	if err != nil {
		return
	}
	// Ensure the encrypted file is closed when done
	defer encryptedFile.Close()
	// Create a new file where decrypted content will be saved
	decryptedFile, err := os.Create(outputPath)
	// Check if there was an error creating the decrypted file
	if err != nil {
		return
	}
	// Ensure the decrypted file is closed when done
	defer decryptedFile.Close()
	// Create a decryption reader that reads from the encrypted file using private key (identity)
	decryptedReader, err := age.Decrypt(encryptedFile, identity)
	// Check if there was an error creating the decrypted reader
	if err != nil {
		return
	}
	// Copy content from decryption reader to decrypted file
	_, err = io.Copy(decryptedFile, decryptedReader)
	// Check if there was an error during the decryption process
	if err != nil {
		return
	}
}

/*
It takes in a path and content to write to that file.
It uses the os.WriteFile function to write the content to that file.
It checks for errors and logs them.
*/
func writeToFile(path string, content []byte) {
	err := os.WriteFile(path, content, 0644)
	if err != nil {
		log.Fatalln(err)
	}
}

func main() {
	// Generate a new AGE key pair
	publicKey, privateKey := generateAgeKeyPair()
	// Parse the private key string into an age.Identity object
	keyPair, err := age.ParseX25519Identity(privateKey)
	// Check if there was an error parsing the private key
	if err != nil {
		return
	}

	// Write the demo sample file.
	writeToFile("input.txt", []byte("Hello, World!"))

	// Encrypt a sample plaintext file ("input.txt")
	encryptFile("input.txt", "encrypted.age", publicKey)

	// Decrypt the previously encrypted file ("encrypted.age")
	decryptFile("encrypted.age", "decrypted.txt", keyPair)
}
