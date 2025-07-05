package main

import (
	"encoding/base64" // Package for Base64 encoding
	"fmt"             // For printing output
)

// getBase64FromText takes a string and returns its Base64-encoded representation.
func getBase64FromText(input string) string {
	data := []byte(input)                              // Convert the input string to a byte slice
	encoded := base64.StdEncoding.EncodeToString(data) // Encode the byte slice to a Base64 string
	return encoded                                     // Return the Base64-encoded string
}

func main() {
	text := "Hello, world!"                     // Define the input text to encode
	encodedText := getBase64FromText(text)      // Call the function to get Base64 encoding
	fmt.Println("Base64 of text:", encodedText) // Print the Base64 result
}
