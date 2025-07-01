package main

import (
	"fmt"
)

// mapToSlices converts a map with string keys and string values to two slices: one for keys and one for values.
func mapToSlices(personData map[string]string) ([]string, []string) {
	// Declare a slice to store all the keys from the map.
	var keyList []string

	// Declare a slice to store all the values from the map.
	var valueList []string

	// Iterate over the map to extract each key and its corresponding value.
	for key, value := range personData {
		// Append the current key to the keyList slice.
		keyList = append(keyList, key)

		// Append the current value to the valueList slice.
		valueList = append(valueList, value)
	}

	// Return the two slices: keyList and valueList.
	return keyList, valueList
}

func main() {
	// Initialize a map containing personal information with string keys and string values.
	personData := map[string]string{
		"name":     "Alice", // The key "name" has the value "Alice".
		"city":     "Paris", // The key "city" has the value "Paris".
		"language": "Go",    // The key "language" has the value "Go".
	}

	// Call the mapToSlices function, passing in the personData map.
	keys, values := mapToSlices(personData)

	// Print the keys slice to the console.
	fmt.Println("Keys:")
	for _, key := range keys {
		fmt.Println(key)
	}

	fmt.Printf("\n")

	// Print the values slice to the console.
	fmt.Println("Values:")
	for _, value := range values {
		fmt.Println(value)
	}
}
