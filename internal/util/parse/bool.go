package parse

import (
	"errors"
	"strings"
)

var UnknownValueError = errors.New("unknown value")

func Bool(s string) (bool, error) {
	switch strings.ToLower(s) {
	case "1", "yes", "y", "true", "t":
		return true, nil
	case "0", "no", "n", "false", "f":
		return false, nil
	default:
		return false, UnknownValueError
	}
}
