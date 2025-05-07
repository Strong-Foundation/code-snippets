package main

import (
	"fmt"
)

func main() {
	// Create a slice of strings.
	randomSlice := []string{"A", "B", "C", "D", "E"}
	// Index to remove.
	indexToRemove := 2
	// Remove the element at the specified index.
	randomSlice = removeElement(randomSlice, indexToRemove)
	// Print the modified slice.
	fmt.Println(randomSlice)
}

// Remove an element from a slice by its index.
func removeElement(slice []string, index int) []string {
	// Check if the index is valid.
	if index < 0 || index >= len(slice) {
		return slice // Return the original slice if the index is out of bounds.
	}
	// Remove the element by appending slices before and after the index.
	return append(slice[:index], slice[index+1:]...)
}
