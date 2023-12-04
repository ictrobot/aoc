package parse

import (
	"regexp"
	"strings"
	"unicode/utf8"
)

var chunkRegexp = regexp.MustCompile(`(?m)^.+$(?:\n^.+$)*`)
var nonWhitespaceRegexp = regexp.MustCompile(`\S+`)

// Lines splits a string based on newlines, removing the last line if empty
func Lines(s string) []string {
	startPos := 0
	var lines []string
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' {
			endPos := i
			if i > 0 && s[i-1] == '\r' {
				endPos--
			}
			lines = append(lines, s[startPos:endPos])
			startPos = i + 1
		}
	}
	if startPos < len(s) {
		lines = append(lines, s[startPos:])
	}
	return lines
}

// Chunks splits a string into "chunks" split by blank lines. Additionally, CLRF is converted to LF
func Chunks(s string) []string {
	return chunkRegexp.FindAllString(strings.ReplaceAll(s, "\r\n", "\n"), -1)
}

// Whitespace splits a string on whitespace
func Whitespace(s string) []string {
	return nonWhitespaceRegexp.FindAllString(s, -1)
}

// Characters splits a string on each character, optimized for ASCII strings
func Characters(s string) []string {
	if len(s) == 0 {
		return nil
	}

	allSingleByte := true
	chars := make([]string, len(s))
	for i := 0; i < len(s); i++ {
		if s[i] >= utf8.RuneSelf {
			allSingleByte = false
			break
		}
		chars[i] = s[i : i+1]
	}
	if allSingleByte {
		return chars
	}

	// Fallback to go's split function which can handle multibyte characters
	return strings.Split(s, "")
}
