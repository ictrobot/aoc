package main

import (
	"errors"
	"fmt"
	"github.com/ictrobot/aoc/internal/api"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

func main() {
	baseDir, err := getBaseDir()
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	// If a year is provided, generate day files for each released puzzle
	// If a year & day is provided, generate day files for that day
	if len(os.Args) > 1 {
		year, err := strconv.Atoi(os.Args[1])
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "invalid year: %v\n", err)
			os.Exit(2)
		}

		var days []int
		if len(os.Args) > 2 {
			day, err := strconv.Atoi(os.Args[2])
			if err != nil {
				_, _ = fmt.Fprintf(os.Stderr, "invalid day: %v\n", err)
				os.Exit(2)
			}

			// Ignore release time as long as it's valid
			if _, err := api.ReleaseTime(year, day); err != nil {
				_, _ = fmt.Fprintln(os.Stderr, err)
				os.Exit(2)
			}

			days = []int{day}
		} else {
			now := time.Now()

			// Find days which have been released
			for i := 1; i < 25; i++ {
				if t, err := api.ReleaseTime(year, i); err == nil && now.After(t) {
					days = append(days, i)
				}
			}
		}

		for _, day := range days {
			if err := generateDay(baseDir, year, day); err != nil {
				_, _ = fmt.Fprintf(os.Stderr, "generating day %d-%02d: %v\n", year, day, err)
				os.Exit(1)
			}
		}
	}

	if err := generateSolutionsFile(baseDir); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "generating solutions file: %v\n", err)
		os.Exit(1)
	}
}

func getBaseDir() (string, error) {
	baseDir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	if _, present := os.LookupEnv("GOFILE"); present {
		if !strings.HasSuffix(baseDir, filepath.Join("internal", "solution")) {
			return "", errors.New("go:generate can only be used in internal/solution")
		}
		if len(os.Args) > 1 {
			return "", errors.New("go:generate can not be used with arguments")
		}
		baseDir = filepath.Dir(filepath.Dir(baseDir))
	}

	if _, err := os.Stat(filepath.Join(baseDir, "go.mod")); errors.Is(err, os.ErrNotExist) {
		return "", err
	}

	return baseDir, nil
}
