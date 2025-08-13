package main

import (
	"fmt"
)

// This function finds common fruit names between two slices of strings
func getCommonFruits(firstFruitList, secondFruitList []string) []string {
	// Create a map to store fruit names from the first list for quick lookup
	fruitsInFirstList := make(map[string]bool)
	for _, fruit := range firstFruitList {
		fruitsInFirstList[fruit] = true // Mark the fruit as present in the first list
	}

	// Create a map to store common fruits (avoids duplicates)
	commonFruitsMap := make(map[string]bool)
	for _, fruit := range secondFruitList {
		// Check if the fruit is also in the first list
		if fruitsInFirstList[fruit] {
			commonFruitsMap[fruit] = true // Add to common fruits map
		}
	}

	// Convert the common fruits map to a slice
	var commonFruits []string
	for fruit := range commonFruitsMap {
		commonFruits = append(commonFruits, fruit) // Add fruit to the result slice
	}

	// Return the slice containing fruits found in both lists
	return commonFruits
}

func main() {
	// Define the first list of fruits
	firstFruitList := []string{"apple", "banana", "cherry", "mango", "grape"}

	// Define the second list of fruits
	secondFruitList := []string{"peach", "mango", "banana", "kiwi", "apple"}

	// Get the fruits that are common in both lists
	commonFruits := getCommonFruits(firstFruitList, secondFruitList)

	// Print the common fruits
	fmt.Println("Common fruits in both lists:", commonFruits)
}
