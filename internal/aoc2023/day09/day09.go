package day09

import (
	_ "embed"
	"github.com/ictrobot/aoc-go/internal/util/parse"
)

//go:embed example
var Example string

type Day09 struct {
	sequences [][]int64
}

func (d *Day09) Parse(input string) {
	d.sequences = nil
	for _, l := range parse.Lines(input) {
		d.sequences = append(d.sequences, parse.ExtractInt64s(l))
	}
}

func (d *Day09) ParseExample() {
	d.Parse(Example)
}

func (d *Day09) Part1() any {
	var sum int64
	for _, s := range d.sequences {
		sum += findNext(s)
	}
	return sum
}

func (d *Day09) Part2() any {
	var sum int64
	for _, s := range d.sequences {
		sum += findPrevious(s)
	}
	return sum
}

func findNext(s []int64) int64 {
	diff := make([]int64, len(s)-1)
	allZero := true
	for i := 1; i < len(s); i++ {
		diff[i-1] = s[i] - s[i-1]
		if diff[i-1] != 0 {
			allZero = false
		}
	}

	if allZero {
		return s[len(s)-1]
	}

	return s[len(s)-1] + findNext(diff)
}

func findPrevious(s []int64) int64 {
	diff := make([]int64, len(s)-1)
	allZero := true
	for i := 1; i < len(s); i++ {
		diff[i-1] = s[i] - s[i-1]
		if diff[i-1] != 0 {
			allZero = false
		}
	}

	if allZero {
		return s[0]
	}

	return s[0] - findPrevious(diff)
}
