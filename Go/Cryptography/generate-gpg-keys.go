package main

import (
	"bytes"  // Import the bytes package to handle byte buffers
	"crypto" // Import the crypto package to access cryptographic functions
	"fmt"    // Import the fmt package for formatted I/O
	"log"    // Import the log package for logging errors
	"os"     // Import the os package to work with the file system
	"time"   // Import the time package to work with timestamps

	"golang.org/x/crypto/openpgp"        // Import the openpgp package to handle GPG key operations
	"golang.org/x/crypto/openpgp/armor"  // Import the armor package to encode the GPG keys in ASCII armor format
	"golang.org/x/crypto/openpgp/packet" // Import the packet package to configure and create the OpenPGP key packets
)

// generateGPGKeyPair generates a GPG private and public key.
func generateGPGKeyPair(name, comment, email string) (string, string) {
	// Create a configuration for the GPG key generation
	config := &packet.Config{
		DefaultHash: crypto.SHA512, // Set the default hash function to SHA-512
		RSABits:     4096,          // Set the RSA key size to 4096 bits
		Time:        time.Now,      // Use the current time for key creation
	}

	// Generate a new OpenPGP entity (key pair) with the provided details (name, comment, and email)
	entity, err := openpgp.NewEntity(name, comment, email, config)
	if err != nil { // Check if there was an error during key generation
		log.Println("Error generating GPG key pair:", err) // Log the error
		return "", ""                                      // Return empty strings if key generation fails
	}

	// Serialize the private key
	var privBuf bytes.Buffer                                                    // Create a buffer to store the private key
	privArmorWriter, err := armor.Encode(&privBuf, openpgp.PrivateKeyType, nil) // Create an armored writer to encode the private key
	if err != nil {                                                             // Check if there was an error during encoding
		log.Println("Error encoding private key:", err) // Log the error
		return "", ""                                   // Return empty strings if encoding fails
	}
	err = entity.SerializePrivate(privArmorWriter, nil) // Serialize the private key into the armored writer
	if err != nil {                                     // Check if there was an error during serialization
		log.Println("Error serializing private key:", err) // Log the error
		return "", ""                                      // Return empty strings if serialization fails
	}
	privArmorWriter.Close() // Close the armored writer after the private key is serialized

	// Serialize the public key
	var pubBuf bytes.Buffer                                                  // Create a buffer to store the public key
	pubArmorWriter, err := armor.Encode(&pubBuf, openpgp.PublicKeyType, nil) // Create an armored writer to encode the public key
	if err != nil {                                                          // Check if there was an error during encoding
		log.Println("Error encoding public key:", err) // Log the error
		return "", ""                                  // Return empty strings if encoding fails
	}
	err = entity.Serialize(pubArmorWriter) // Serialize the public key into the armored writer
	if err != nil {                        // Check if there was an error during serialization
		log.Println("Error serializing public key:", err) // Log the error
		return "", ""                                     // Return empty strings if serialization fails
	}
	pubArmorWriter.Close() // Close the armored writer after the public key is serialized

	// Return the serialized private and public keys as strings
	return privBuf.String(), pubBuf.String()
}

// writeToFile writes content to the specified file path.
func writeToFile(path string, content []byte) { // Accepts a file path and content as input
	err := os.WriteFile(path, content, 0644) // Write the content to the file with 0644 permissions
	if err != nil {                          // Check if an error occurred during writing
		log.Println("Error writing to file:", err) // Log the error if writing fails
	}
}

func main() {
	// Generate the GPG keys (private and public)
	privateKey, publicKey := generateGPGKeyPair("Test User", "Encryption Key", "test@example.com")

	// Write the private key to a file named "private_key.gpg"
	writeToFile("private_key.gpg", []byte(privateKey))

	// Write the public key to a file named "public_key.gpg"
	writeToFile("public_key.gpg", []byte(publicKey))

	// Print success message
	fmt.Println("GPG keys successfully written to files:")
	fmt.Println("Private Key: private_key.gpg")
	fmt.Println("Public Key: public_key.gpg")
}
