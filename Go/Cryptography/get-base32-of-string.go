package main

import (
	"encoding/base32"
	"fmt"
)

// base32Encode takes a string and returns its Base32 encoded version.
func base32Encode(input string) string {
	return base32.StdEncoding.EncodeToString([]byte(input))
}

func main() {
	input := "Hello, Base32 in Go!"
	encoded := base32Encode(input)
	fmt.Println("Encoded:", encoded)
}
