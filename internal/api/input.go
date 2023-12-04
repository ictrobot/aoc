package api

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"time"
	"unicode/utf8"
)

const firstYear = 2015
const releaseHour = 5 // 5am UTC

var inputCache sync.Map

func ReleaseTime(year, day int) (time.Time, error) {
	if year < firstYear || year > 9999 {
		return time.Time{}, fmt.Errorf("invalid year: %v", year)
	}
	if day < 1 || day > 25 {
		return time.Time{}, fmt.Errorf("invalid day: %v", day)
	}

	return time.Date(year, 12, day, releaseHour, 0, 0, 0, time.UTC), nil
}

func GetInput(year, day int) (string, error) {
	releaseTime, err := ReleaseTime(year, day)
	if err != nil {
		return "", err
	} else if time.Now().Before(releaseTime) {
		return "", fmt.Errorf("%d-%02d is not released yet", year, day)
	}

	file := filepath.Join("inputs", fmt.Sprintf("%d", year), fmt.Sprintf("day%02d", day))

	if v, ok := inputCache.Load(file); ok {
		return v.(string), nil
	}

	if b, err := os.ReadFile(file); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "reading input: %v\n", err)
		// don't return error so we attempt to retrieve input
	} else if len(b) > 0 {
		if !utf8.Valid(b) {
			panic("invalid utf8 in input")
		}

		s := string(b)
		inputCache.Store(file, s)
		return s, nil
	}

	if s, err := retrieveInput(year, day, file); err != nil {
		return "", err
	} else {
		inputCache.Store(file, s)
		return s, nil
	}
}
