//go:build nohttp

package api

import "fmt"

func retrieveInput(year, day int, file string) (string, error) {
	return "", fmt.Errorf("retrieving input from adventofcode.com disabled, please manually populate %s", file)
}
