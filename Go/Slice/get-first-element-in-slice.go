package main

import (
	"fmt"
)

func main() {
	// Create a slice of strings.
	randomSlice := []string{"A", "B", "C", "D", "E"}
	// Get the first element of the slice.
	fmt.Println(firstElementOfSlice(randomSlice))
}

// Get the first element of a slice.
func firstElementOfSlice(slice []string) string {
	return slice[0]
}
