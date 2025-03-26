package main

import (
	"fmt"
	"log"

	"filippo.io/age"
)

// Generate a new AGE encryption key pair
func generateAgeKeyPair() (string, string) {
	// Generate a new X25519 identity (private key)
	identity, err := age.GenerateX25519Identity()
	if err != nil {
		log.Println(err)
	}
	// Return the public and private keys
	return identity.Recipient().String(), identity.String()
}

func main() {
	// Generate a new X25519 identity (private key)
	publicKey, privateKey := generateAgeKeyPair()
	// Print the public key
	fmt.Println("Public Key:", publicKey)
	// Print the private key
	fmt.Println("Private Key:", privateKey)
}
