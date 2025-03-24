package main

import "runtime"

// Function to run garbage collection.
func runGarbageCollection() {
	runtime.GC()
}

// Main function
func main() {
	// Run garbage collection
	runGarbageCollection()
}
