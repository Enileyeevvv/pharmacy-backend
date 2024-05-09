package utils

import (
	"fmt"
)

func ConnectionURLBuilder(n string) (string, error) {
	var url string

	switch n {
	case "postgres":
		url = fmt.Sprintf(
			"host=localhost port=5432 user=postgres password=password dbname=pharmacy sslmode=disable",
		)

	case "fiber":
		// URL for Fiber connection.
		url = fmt.Sprintf(
			"0.0.0.0:9000",
		)
	default:
		// Return error message.
		return "", fmt.Errorf("connection name '%v' is not supported", n)
	}

	// Return connection URL.
	return url, nil
}
