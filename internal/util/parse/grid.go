package parse

import (
	"fmt"
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
