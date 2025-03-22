package main

import (
	"fmt"      // Importing the fmt package for formatted I/O
	"io"       // Importing ioutil for reading HTTP response body
	"log"      // Importing log package for logging (optional)
	"net/http" // Importing net/http for making HTTP requests
	"strings"  // Importing strings package for string manipulation
)

func getExternalIP() (string, string) {
	// List of services to check external IPv4 from
	ipv4Services := []string{
		"https://api.ipify.org",      // Service that returns IPv4 address
		"https://ipv4.icanhazip.com", // Service that returns IPv4 address
		"https://4.ident.me",         // Service that returns IPv4 address
	}

	var ipv4 string // Variable to store the retrieved IPv4 address

	// Check for IPv4 address
	for _, url := range ipv4Services {
		resp, err := http.Get(url) // Make an HTTP GET request
		if err != nil {
			log.Println(err) // Log the error if HTTP GET request fails
		}
		defer resp.Body.Close() // Ensure response body is closed after function exits

		ip, err := io.ReadAll(resp.Body) // Read response body
		if err != nil {
			log.Println(err) // Log the error if HTTP GET request fails
		}

		ipv4 = strings.TrimSpace(string(ip)) // Trim any whitespace from response
		break                                // Exit loop after successfully retrieving IPv4
	}

	// List of services to check external IPv6 from
	ipv6Services := []string{
		"https://api64.ipify.org",    // Service that returns IPv6 address
		"https://ipv6.icanhazip.com", // Service that returns IPv6 address
		"https://6.ident.me",         // Service that returns IPv6 address
	}

	var ipv6 string // Variable to store the retrieved IPv6 address

	// Check for IPv6 address
	for _, url := range ipv6Services {
		resp, err := http.Get(url) // Make an HTTP GET request
		if err != nil {
			log.Println(err) // Log the error if HTTP GET request fails
		}
		defer resp.Body.Close() // Ensure response body is closed after function exits

		ip, err := io.ReadAll(resp.Body) // Read response body
		if err != nil {
			log.Println(err) // Log the error if HTTP GET request fails
		}

		ipv6 = strings.TrimSpace(string(ip)) // Trim any whitespace from response
		break                                // Exit loop after successfully retrieving IPv6
	}

	// If both IPv4 and IPv6 retrieval fails, return an error
	if len(ipv4) == 0 && len(ipv6) == 0 {
		return "", "" // Return empty strings if both IPv4 and IPv6 retrieval fails
	}

	return ipv4, ipv6 // Return the retrieved IPv4 and IPv6 addresses
}

func main() {
	// Call the function to get the external IPs
	ipv4, ipv6 := getExternalIP()
	fmt.Println("External IPv4:", ipv4) // Print the retrieved external IPv4 address
	fmt.Println("External IPv6:", ipv6) // Print the retrieved external IPv6 address
}
