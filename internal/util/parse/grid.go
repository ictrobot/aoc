package parse

import (
	"fmt"
	"github.com/ictrobot/aoc-go/internal/util/structures"
)

type Splitter func(string) []string

func Grid[T any](row Splitter, items func(string) []T, s string) [][]T {
	rows := row(s)
	results := make([][]T, len(rows))
	for i := 0; i < len(rows); i++ {
		results[i] = items(rows[i])
	}
	return results
}

func ReflectGrid[T any](row, col Splitter, s string) [][]T {
	rows := row(s)
	results := make([][]T, len(rows))
	for i := 0; i < len(rows); i++ {
		var err error
		results[i], err = Reflect[[]T](col(rows[i]))
		if err != nil {
			panic(fmt.Sprintf("failed to parse grid: reading slice element %d: %v", i, err))
		}
	}
	return results
}

// ByteGrid converts a 2D grid represented as a string into a
// structures.FlatGrid. Each row should be represented by a line, and each
// item within that line should be represented by a single byte. This is useful
// for e.g. maze tasks where coordinates always stay inside the input grid
func ByteGrid(s string) *structures.FlatGrid[byte] {
	if s == "" {
		return nil
	}

	b := make([]byte, 0, len(s))
	width := -1
	height := 0

	for i := 0; i < len(s); {
		j := i
		for j < len(s) && s[j] != '\n' {
			j++
		}

		var row string
		if j > 0 && s[j-1] == '\r' {
			row = s[i : j-1]
		} else {
			row = s[i:j]
		}

		if width == -1 {
			width = len(row)
		} else if width != len(row) {
			panic("rows have different lengths")
		}

		b = append(b, row...)
		height++

		i = j + 1
	}

	return &structures.FlatGrid[byte]{S: b, Width: width, Height: height}
}
