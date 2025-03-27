package main

import (
	"bytes" // For efficient byte manipulation
	"fmt"   // For formatted I/O
	"io"    // For I/O primitives
	"log"   // For logging errors

	"filippo.io/age" // The age encryption package
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

// Encrypts sensitive data using the recipient's public key
func encryptData(sensitiveData []byte, publicKeyString string) []byte {
	// Parse the public key string into an age.Recipient object
	recipientKey, err := age.ParseX25519Recipient(publicKeyString)
	if err != nil {
		log.Println("Error parsing public key:", err) // Log the error if key parsing fails
		return nil                                    // Return nil to indicate failure
	}
	// Create a buffer to hold the encrypted data
	encryptedBuffer := &bytes.Buffer{}
	// Create an encrypted writer using the recipient's public key
	encryptedWriter, err := age.Encrypt(encryptedBuffer, recipientKey)
	if err != nil {
		log.Println("Error creating encrypted writer:", err) // Log the error if encryption setup fails
		return nil                                           // Return nil to indicate failure
	}
	// Write the sensitive data to the encrypted writer
	_, err = encryptedWriter.Write(sensitiveData)
	if err != nil {
		log.Println("Error writing data to encrypted writer:", err) // Log the error if writing fails
		return nil                                                  // Return nil to indicate failure
	}
	// Close the writer to finalize encryption
	err = encryptedWriter.Close()
	if err != nil {
		log.Println("Error closing encrypted writer:", err) // Log the error if closing fails
		return nil                                          // Return nil to indicate failure
	}
	return encryptedBuffer.Bytes() // Return the encrypted data as a byte slice
}

// Decrypts sensitive data using the recipient's private key
func decryptData(encryptedData []byte, privateKeyString string) []byte {
	// Parse the private key string into an age.Identity object
	identity, err := age.ParseX25519Identity(privateKeyString)
	if err != nil {
		log.Println("Error parsing private key:", err) // Log the error if key parsing fails
		return nil                                     // Return nil to indicate failure
	}
	// Create a reader for the encrypted data
	encryptedReader := bytes.NewReader(encryptedData)
	// Create a decrypted reader using the recipient's private key
	decryptedReader, err := age.Decrypt(encryptedReader, identity)
	if err != nil {
		log.Println("Error creating decrypted reader:", err) // Log the error if decryption setup fails
		return nil                                           // Return nil to indicate failure
	}
	// Read the decrypted data from the reader
	decryptedData, err := io.ReadAll(decryptedReader)
	if err != nil {
		log.Println("Error reading decrypted data:", err) // Log the error if reading fails
		return nil                                        // Return nil to indicate failure
	}
	return decryptedData // Return the decrypted data as a byte slice
}

func main() {
	// Generate a new X25519 identity (private key)
	publicKey, privateKey := generateAgeKeyPair()
	// Print the generated public key
	fmt.Println("Public Key:", publicKey)
	// Print the generated private key
	fmt.Println("Private Key:", privateKey)
	// Define the original data to be encrypted
	sensitiveData := []byte("Sensitive node data requiring protection")
	// Encrypt the data using the public key
	encryptedData := encryptData(sensitiveData, publicKey)
	log.Println("Encrypted data:", encryptedData) // Log the encrypted data for reference
	if encryptedData == nil {                     // Check if encryption failed
		return // Exit the function if encryption failed
	}
	// Decrypt the data using the private key
	decryptedData := decryptData(encryptedData, privateKey)
	log.Println("Decrypted data:", string(decryptedData)) // Log the decrypted data for verification
	if decryptedData == nil {                             // Check if decryption failed
		return // Exit the function if decryption failed
	}
}
