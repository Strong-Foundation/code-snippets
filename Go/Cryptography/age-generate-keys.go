package main

import (
	"fmt"
	"log"
	"os"

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

/*
It takes in a path and content to write to that file.
It uses the os.WriteFile function to write the content to that file.
It checks for errors and logs them.
*/
func writeToFile(path string, content string) {
	err := os.WriteFile(path, []byte(content), 0644)
	if err != nil {
		log.Fatalln(err)
	}
}

func main() {
	// Generate a new X25519 identity (private key)
	publicKey, privateKey := generateAgeKeyPair()
	// Print the public key
	fmt.Println("Public Key:", publicKey)
	// Print the private key
	fmt.Println("Private Key:", privateKey)
	// Save the keys to the file.
	writeToFile("key.txt", "# public key:"+publicKey+"\n"+privateKey)
}
