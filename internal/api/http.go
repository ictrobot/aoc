//go:build !nohttp

package api

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

const (
	tokenEnvVar   = "AOC_TOKEN"
	tokenFilename = ".aoc_token"
)

func retrieveInput(year, day int, file string) (string, error) {
	token, tokenFile, err := getSessionToken()
	if err != nil {
		return "", fmt.Errorf("retrieving session token: %w", err)
	}

	req, err := http.NewRequest("GET", fmt.Sprintf("http://adventofcode.com/%d/day/%d/input", year, day), nil)
	if err != nil {
		panic(err)
	}
	req.Header.Set("User-Agent", "github.com/ictrobot/aoc")
	req.Header.Set("Cookie", "session="+token)

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("fetching input: %w", err)
	}

	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)
	body, err := io.ReadAll(resp.Body)

	if resp.StatusCode == 200 {
		// failing to save input shouldn't prevent running solution
		dir := filepath.Dir(file)
		if err := os.MkdirAll(dir, 0755); err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "creating inputs directories: %v\n", err)
		}

		if err := os.WriteFile(file, body, 0644); err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "saving input: %v\n", err)
		}

		return string(body), nil
	}

	var action string
	if (resp.StatusCode == 400 || resp.StatusCode == 500) && tokenFile != "" {
		// invalid token returns HTTP 400 or 500
		if err := os.Remove(tokenFile); err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "deleting saved token: %v\n", err)
		} else {
			action = ", deleted token"
		}
	}

	return "", fmt.Errorf("fetching input: HTTP %v returned%v", resp.StatusCode, action)
}

func getSessionToken() (string, string, error) {
	if token, present := os.LookupEnv(tokenEnvVar); present && token != "" {
		return token, "", nil
	}

	home, err := os.UserHomeDir()
	if err != nil {
		return "", "", fmt.Errorf("getting user home: %w", err)
	}

	file := filepath.Join(home, tokenFilename)
	if b, err := os.ReadFile(file); err == nil && len(b) > 0 {
		return string(b), file, nil
	} else if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "reading token from file: %v\n", err)
	}

	_, _ = fmt.Fprintf(os.Stderr, "\nplease supply session token for downloading inputs, or manually populate the input file\nsession token: ")

	var token string
	if _, err := fmt.Scan(&token); err != nil {
		return "", "", fmt.Errorf("reading token from stdin: %w", err)
	} else if token == "" {
		return "", "", fmt.Errorf("no session token provided")
	}

	if err := os.WriteFile(file, []byte(token), 0600); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "writing token to file: %v\n", err)
	}

	return token, file, nil
}
