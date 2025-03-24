package main

import (
	"fmt"     // Import the "fmt" package for formatted I/O operations
	"os"      // Import the "os" package for interacting with the operating system (like retrieving executable path)
	"os/exec" // Import the "os/exec" package to run external commands (used to start the application as an admin)
	"runtime" // Import the "runtime" package to check the operating system at runtime
	"syscall" // Import the "syscall" package for interacting with low-level system calls (used to request admin rights)
)

// runAsAdmin tries to execute the current program with administrator privileges
func runAsAdmin() error {
	// Check if the operating system is Windows, as this method works only on Windows
	if runtime.GOOS != "windows" {
		return fmt.Errorf("this function only works on Windows") // Return an error if not running on Windows
	}

	// Get the current executable's path (the app itself)
	exePath, err := os.Executable()
	if err != nil {
		return err // Return the error if unable to retrieve the executable path
	}

	// Create a command to execute the same application with admin privileges
	cmd := exec.Command(exePath) // Create an exec command to run the same executable
	cmd.SysProcAttr = &syscall.SysProcAttr{
		CreationFlags: syscall.CREATE_NEW_CONSOLE | syscall.CREATE_NO_WINDOW, // Set flags to run the app with new console and no window
	}

	// Start the command (launch the app with elevated privileges)
	cmd.Start()
	return nil // Return nil to indicate that the operation succeeded
}

func main() {
	// Attempt to run the app with admin privileges
	err := runAsAdmin()
	if err != nil {
		fmt.Println("Error running as admin:", err) // Print an error message if the app fails to run with admin privileges
		return
	}

	// If the app is running as admin, print a success message
	fmt.Println("Application running with admin privileges.")
	// Your application's code would go here, after ensuring admin rights are granted
}
