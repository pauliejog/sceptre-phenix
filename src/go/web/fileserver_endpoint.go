package web

import (
	"fmt"
	"net"
	"strconv"
	"strings"
)

// normalizeFileServerEndpoint converts the fileserver option into a listen endpoint.
func normalizeFileServerEndpoint(value string) (string, error) {
	value = strings.TrimSpace(value)
	if value == "" || value == "0" {
		return "", nil
	}

	if _, err := parseFileServerPort(value); err == nil {
		return net.JoinHostPort("127.0.0.1", value), nil
	}

	host, port, err := net.SplitHostPort(value)
	if err != nil {
		return "", fmt.Errorf("expected port or host:port")
	}

	if _, err := parseFileServerPort(port); err != nil {
		return "", err
	}

	return net.JoinHostPort(host, port), nil
}

// parseFileServerPort validates a numeric TCP port.
func parseFileServerPort(value string) (int, error) {
	port, err := strconv.Atoi(value)
	if err != nil {
		return 0, fmt.Errorf("invalid port %q", value)
	}

	if port < 0 || port > 65535 {
		return 0, fmt.Errorf("port %d out of range", port)
	}

	return port, nil
}
