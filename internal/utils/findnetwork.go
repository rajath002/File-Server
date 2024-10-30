package utils

import (
	"fmt"
	"net"
)

// GetWiFiIPv4Address retrieves the IPv4 address of the active WiFi interface.
func GetWiFiIPv4Address() (string, error) {
	interfaces, err := net.Interfaces()
	if err != nil {
		return "", err
	}

	for _, iface := range interfaces {
		// Filter to skip interfaces that are not up or are loopback
		if iface.Flags&net.FlagUp == 0 || iface.Flags&net.FlagLoopback != 0 {
			continue
		}

		addrs, err := iface.Addrs()
		if err != nil {
			return "", err
		}

		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}

			// Check if IP is IPv4 and not loopback
			if ip != nil && ip.To4() != nil {
				return ip.String(), nil // Return the first found IPv4 address
			}
		}
	}

	return "", fmt.Errorf("no active WiFi IPv4 address found")
}
