package parse

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestLines(t *testing.T) {
	assert.Equal(t, ([]string)(nil), Lines(""))
	assert.Equal(t, []string{""}, Lines("\n"))
	assert.Equal(t, []string{"a", "b", "c"}, Lines("a\nb\nc"))
	assert.Equal(t, []string{"a", "b", "c"}, Lines("a\r\nb\nc"))
	assert.Equal(t, []string{"a", "b", "c"}, Lines("a\r\nb\nc\n"))
}

func TestChunks(t *testing.T) {
	assert.Equal(t, []string{"a\nb\nc", "1\n2\n3"}, Chunks("a\nb\nc\n\n1\n2\n3\n"))
	assert.Equal(t, []string{"a\nb\nc", "1\n2\n3"}, Chunks("a\r\nb\r\nc\r\n\r\n1\r\n2\r\n3\r\n"))
	assert.Equal(t, []string{"a\nb", "1", "c\nd"}, Chunks("\na\nb\n\n\n1\n\nc\nd\n\n"))
}

func TestWhitespace(t *testing.T) {
	assert.Equal(t, []string{"a", "b", "1", "2", "3", "4", "5", "6", "cde"}, Whitespace("\na b 1\n2 3\r\n4\r\n\r\n5  \n  6\n\n\ncde\n"))
}

func TestCharacters(t *testing.T) {
	assert.Equal(t, ([]string)(nil), Characters(""))
	assert.Equal(t, []string{"a", "b", "c", "d", "e", "f"}, Characters("abcdef"))
	assert.Equal(t, []string{" ", "\n", " ", "\r", "\n", " "}, Characters(" \n \r\n "))

	// Unicode fallback
	assert.Equal(t, []string{"1", "2", "3", "🚧", "🌴"}, Characters("123🚧🌴"))
}

func BenchmarkCharacters(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Characters("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	}
}

func BenchmarkStringsSplit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		strings.Split("ABCDEFGHIJKLMNOPQRSTUVWXYZ", "")
	}
}
