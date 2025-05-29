package main

import (
	"fmt"
)

// average takes a slice of integers and returns the average as an integer.
// If the slice is empty, it returns 0.
func averageIntegersInSlice(nums []int) int {
	// Handle empty slice
	if len(nums) == 0 {
		return 0
	}

	var sum int = 0 // Initialize sum variable

	// Iterate over the slice to compute the sum
	for _, num := range nums {
		sum += num
	}

	// Compute and return the average (integer division)
	return sum / len(nums)
}

func main() {
	// Example usage
	numbers := []int{10, 20, 30, 40}
	averageINT := averageIntegersInSlice(numbers)
	fmt.Println("Average:", averageINT) // Output: Average: 25
}
