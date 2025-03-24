package main

import "runtime/debug"

// Function to run FreeOSMemory
func runFreeOSMemory() {
	// Run garbage collection
	debug.FreeOSMemory() // Suggest the OS to free unused memory

}

// Main function
func main() {
	// Run free OS memory
	runFreeOSMemory()
}
