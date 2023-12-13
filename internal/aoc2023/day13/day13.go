package day13

import (
	_ "embed"
	"github.com/ictrobot/aoc/internal/util/parse"
	"slices"
	"strings"
)

//go:embed example
var Example string

type Day13 struct {
	patterns [][]string
}

func (d *Day13) Parse(input string) {
	d.patterns = parse.Grid(parse.Chunks, parse.Lines, input)
}

func (d *Day13) ParseExample() {
	d.Parse(Example)
}

func (d *Day13) Part1() any {
	return d.summarize(slices.Equal[[]string, string])
}

func (d *Day13) Part2() any {
	return d.summarize(offByOne)
}

func (d *Day13) summarize(check func([]string, []string) bool) int {
	var sum int

	for _, p := range d.patterns {
		if h := checkHorizontalReflection(p, check); h >= 0 {
			sum += 100 * h
			continue
		}

		if v := checkVerticalReflection(p, check); v >= 0 {
			sum += v
			continue
		}

		panic("no reflection")
	}
	return sum
}

func checkHorizontalReflection(p []string, check func([]string, []string) bool) int {
	// reuse slice for reverse slice to reduce allocation
	reversed := make([]string, len(p)-1)

	for y := 1; y < len(p); y++ {
		l := min(y, len(p)-y)
		before := p[y-l : y]

		copy(reversed[:l], p[y:y+l])
		slices.Reverse(reversed[:l])

		if check(before, reversed[:l]) {
			return y
		}
	}

	return -1
}

func checkVerticalReflection(p []string, check func([]string, []string) bool) int {
	// assumes single byte characters
	transposed := make([]string, len(p[0]))
	for i := range p[0] {
		var b strings.Builder
		b.Grow(len(p))

		for j := range p {
			b.WriteByte(p[j][i])
		}

		transposed[i] = b.String()
	}

	return checkHorizontalReflection(transposed, check)
}

func offByOne(a, b []string) bool {
	var diff int
	for i := range a {
		if a[i] == b[i] {
			continue
		}

		for j := range a[i] {
			if a[i][j] == b[i][j] {
				continue
			}

			diff++
			if diff > 1 {
				// early return
				return false
			}
		}
	}
	return diff == 1
}
