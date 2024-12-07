package main

import (
	"fmt"
	"net"
)

// GetIPsInCIDR returns all IP addresses in a given CIDR block.
func GetIPsInCIDR(cidr string) ([]string, error) {
	// Parse the CIDR block
	_, ipNet, err := net.ParseCIDR(cidr)
	if err != nil {
		return nil, err
	}

	// Calculate the number of IPs in the CIDR block
	ips := []string{}
	for ip := ipNet.IP.Mask(ipNet.Mask); ipNet.Contains(ip); {
		ips = append(ips, ip.String())

		// Increment the IP address by 1
		for j := len(ip) - 1; j >= 0; j-- {
			ip[j]++
			if ip[j] != 0 {
				break
			}
		}
	}

	return ips, nil
}

func main() {
	cidr := "192.168.1.0/29"
	ips, err := GetIPsInCIDR(cidr)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Print all IPs in the CIDR
	for _, ip := range ips {
		fmt.Println(ip)
	}
}
