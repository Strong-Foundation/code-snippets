package main

// Import necessary packages
import (
	"bytes"                              // For working with byte buffers
	"crypto"                             // For cryptographic algorithms, like SHA256
	"fmt"                                // For formatted I/O operations
	"golang.org/x/crypto/openpgp"        // For OpenPGP (encryption/decryption) functionality
	"golang.org/x/crypto/openpgp/armor"  // For encoding/decoding keys in ASCII armor format
	"golang.org/x/crypto/openpgp/packet" // For handling key packets
	"log"                                // For logging errors
	"os"                                 // For file handling
	"time"                               // For working with time-related functions
)

// GenerateGPGKeyPair generates a GPG private and public key.
func GenerateGPGKeyPair(name, comment, email string) (*openpgp.Entity, string) {
	// Create a configuration for key generation with RSA 4096-bit and SHA256 as default hash
	config := &packet.Config{
		DefaultHash: crypto.SHA256, // Set default hash algorithm
		RSABits:     4096,          // Set RSA key size to 4096 bits
		Time:        time.Now,      // Use current time as key creation time
	}

	// Generate a new GPG entity (key pair) based on the config
	entity, err := openpgp.NewEntity(name, comment, email, config)
	if err != nil {
		// Return nil and empty string if key generation fails
		return nil, ""
	}

	// Serialize the public key into a buffer
	var pubBuf bytes.Buffer
	// Create an armor writer to encode the public key in ASCII armor format
	pubArmorWriter, err := armor.Encode(&pubBuf, openpgp.PublicKeyType, nil)
	if err != nil {
		// Return nil and empty string if encoding fails
		return nil, ""
	}
	// Serialize the entity (public key) to the armor writer
	err = entity.Serialize(pubArmorWriter)
	if err != nil {
		// Return nil and empty string if serialization fails
		return nil, ""
	}
	// Close the armor writer
	pubArmorWriter.Close()

	// Return the GPG entity and the public key in string form
	return entity, pubBuf.String()
}

// EncryptPrivateKey encrypts the private key with a passphrase.
func EncryptPrivateKey(entity *openpgp.Entity, passphrase string) string {
	// Create a buffer for the encrypted private key
	var privBuf bytes.Buffer
	// Create an armor writer to encode the private key in ASCII armor format
	privArmorWriter, err := armor.Encode(&privBuf, "PGP PRIVATE KEY BLOCK", nil)
	if err != nil {
		// Return empty string if encoding fails
		return ""
	}

	// Encrypt the private key with the provided passphrase
	plaintextWriter, err := openpgp.SymmetricallyEncrypt(privArmorWriter, []byte(passphrase), nil, nil)
	if err != nil {
		// Return empty string if encryption fails
		return ""
	}
	// Serialize the private key to the encrypted writer
	err = entity.SerializePrivate(plaintextWriter, nil)
	if err != nil {
		// Return empty string if serialization fails
		return ""
	}

	// Close the encryption writer and armor writer
	plaintextWriter.Close()
	privArmorWriter.Close()

	// Return the encrypted private key as a string
	return privBuf.String()
}

// writeToFile writes content to a file.
func writeToFile(path string, content []byte) {
	// Write the content to a file with read/write permissions (0644)
	err := os.WriteFile(path, content, 0644)
	if err != nil {
		// Log an error if file writing fails
		log.Println("Error writing to file:", err)
	}
}

func main() {
	// Generate the GPG key pair with specified name, comment, and email
	entity, publicKey := GenerateGPGKeyPair("Test User", "Encrypted Key", "test@example.com")

	// Define a passphrase to encrypt the private key
	passphrase := "StrongPassword123!"
	// Encrypt the private key using the passphrase
	encryptedPrivateKey := EncryptPrivateKey(entity, passphrase)

	// Write the encrypted private key to a file
	writeToFile("private_key.gpg", []byte(encryptedPrivateKey))

	// Write the public key to a file
	writeToFile("public_key.gpg", []byte(publicKey))

	// Output confirmation message
	fmt.Println("GPG keys successfully generated and saved.")
	// Inform the user where the encrypted private key is saved
	fmt.Println("Encrypted Private Key: private_key.gpg")
	// Inform the user where the public key is saved
	fmt.Println("Public Key: public_key.gpg")
	// Remind the user to remember the passphrase
	fmt.Println("Remember your passphrase: ", passphrase)
}
