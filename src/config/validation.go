package config

import (
	"fmt"
	"regexp"
)

// ValidatePort validates that a port number is within the valid range (1-65535)
func ValidatePort(port int) error {
	if port < 1 || port > 65535 {
		return fmt.Errorf("invalid port: %d, must be between 1 and 65535", port)
	}
	return nil
}

// ValidateHost validates that a host is a valid hostname or IP address
func ValidateHost(host string) error {
	if host == "" {
		return fmt.Errorf("host cannot be empty")
	}

	// Check if it's a valid IP address (IPv4 or IPv6)
	ipRegex := regexp.MustCompile(`^(\d{1,3}\.){3}\d{1,3}$|^\[?[0-9a-fA-F:]+\]?$`)
	if ipRegex.MatchString(host) {
		return nil
	}

	// Check if it's a valid hostname
	hostnameRegex := regexp.MustCompile(`^[a-zA-Z0-9]([a-zA-Z0-9\-]{0,61}[a-zA-Z0-9])?(\.[a-zA-Z0-9]([a-zA-Z0-9\-]{0,61}[a-zA-Z0-9])?)*$`)
	if hostnameRegex.MatchString(host) {
		return nil
	}

	return fmt.Errorf("invalid host: %s, must be a valid hostname or IP address", host)
}

// ValidateTimeout validates that a timeout value is positive
func ValidateTimeout(timeout int) error {
	if timeout < 1 {
		return fmt.Errorf("invalid timeout: %d, must be greater than 0", timeout)
	}
	return nil
}
