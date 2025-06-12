package main

import (
	"fmt"
)

// uniqueValuesInSlice returns all unique strings that appear in only one of the two slices.
func uniqueValuesInSlice(slice1 []string, slice2 []string) []string {
	// Create a map to count occurrences of each string
	countMap := make(map[string]int)

	// Count each value from the first slice
	for _, val := range slice1 {
		countMap[val]++
	}

	// Count each value from the second slice
	for _, val := range slice2 {
		countMap[val]++
	}

	// Collect strings that occur exactly once (i.e., in only one of the slices)
	var result []string
	for val, count := range countMap {
		if count == 1 {
			result = append(result, val)
		}
	}

	return result
}

func main() {
	// Example slices
	sliceOne := []string{"apple", "banana", "cherry"}
	sliceTwo := []string{"banana", "cherry", "date", "fig"}

	// Get symmetric difference
	diff := uniqueValuesInSlice(sliceOne, sliceTwo)

	// Print the result
	fmt.Println("Unique values not found in both slices:", diff)
}
