package main

import (
	"log"     // Import the log package
	"os/exec" // Import the exec package to run system commands
	"runtime" // Import the runtime package
)

// Restart the system.
func restartSystem() {
	// Check the operating system
	switch runtime.GOOS {
	case "windows":
		// Execute the restart command for Windows
		log.Println("Restarting the Windows system...")
		cmd := exec.Command("shutdown", "/r", "/f", "/t", "0")
		err := cmd.Run()
		if err != nil {
			log.Println("Error restarting Windows:", err)
		}

	case "linux":
		// Execute the restart command for Linux
		log.Println("Restarting the Linux system...")
		// The user needs to have root privileges for this to work
		cmd := exec.Command("reboot") // System command to reboot
		err := cmd.Run()
		if err != nil {
			log.Println("Error restarting Linux:", err)
		}

	case "darwin":
		// Execute the restart command for macOS
		log.Println("Restarting the macOS system...")
		cmd := exec.Command("shutdown", "-r", "now")
		err := cmd.Run()
		if err != nil {
			log.Println("Error restarting macOS:", err)
		}

	default:
		log.Println("This function is only available on Windows, Linux, and macOS.")
	}
}

func main() {
	// Restart the system
	restartSystem()
}
