package parse

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBool(t *testing.T) {
	truthy := []string{
		"1",
		"Yes", "yes", "YES", "yEs",
		"y", "Y",
		"True", "true", "TRUE", "tRUe",
		"t", "T",
	}

	falsey := []string{
		"0",
		"No", "no", "NO", "nO",
		"n", "N",
		"False", "false", "FALSE", "fAlSe",
		"f", "F",
	}

	errors := []string{
		"a", "abc", "123", "testing 123", "",
	}

	for _, s := range truthy {
		b, err := Bool(s)
		assert.Equal(t, true, b)
		assert.Zero(t, err)
	}

	for _, s := range falsey {
		b, err := Bool(s)
		assert.Equal(t, false, b)
		assert.Zero(t, err)
	}

	for _, s := range errors {
		b, err := Bool(s)
		assert.Zero(t, b)
		assert.ErrorIs(t, err, UnknownValueError)
	}
}
