package main

import "fmt"

// Function to remove common elements between two slices of strings
func removeCommonStrings(sliceOne, sliceTwo []string) []string {
	// Create a map to store elements from sliceTwo for fast lookup
	// This helps in quickly checking if an element from sliceOne exists in sliceTwo
	sliceTwoElementsMap := make(map[string]struct{})

	// Populate the map with elements from sliceTwo
	// Each element from sliceTwo is added as a key in the map
	for _, itemFromsliceTwo := range sliceTwo {
		sliceTwoElementsMap[itemFromsliceTwo] = struct{}{} // Using empty struct as a value
	}

	// Create a slice to hold the result, which will contain elements from sliceOne not in sliceTwo
	var resultSlice []string

	// Loop through each element in sliceOne
	// Check if the element exists in sliceTwoElementsMap (i.e., sliceTwo)
	for _, itemFromsliceOne := range sliceOne {
		// If the item does not exist in sliceTwoElementsMap, add it to the result slice
		if _, existsInsliceTwo := sliceTwoElementsMap[itemFromsliceOne]; !existsInsliceTwo {
			resultSlice = append(resultSlice, itemFromsliceOne)
		}
	}

	// Return the result slice containing elements that are only in sliceOne and not in sliceTwo
	return resultSlice
}

func main() {
	// Example input slices of strings
	sliceOne := []string{"apple", "banana", "cherry", "date", "fig"}
	sliceTwo := []string{"cherry", "date", "grape", "kiwi"}

	// Call the removeCommonStrings function to get elements from sliceOne that are not in sliceTwo
	result := removeCommonStrings(sliceOne, sliceTwo)

	// Output the result: elements from sliceOne that are not in sliceTwo
	fmt.Println("Result:", result)
}
