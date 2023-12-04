package day01

import (
	_ "embed"
	"github.com/ictrobot/aoc/internal/util/parse"
	"github.com/samber/lo"
	"slices"
)

//go:embed example
var Example string

type Day01 struct {
	elves []int
}

func (d *Day01) Parse(input string) {
	d.elves = lo.Map(parse.Chunks(input), func(s string, index int) int {
		return lo.Sum(parse.ExtractInts(s))
	})
}

func (d *Day01) ParseExample() {
	d.Parse(Example)
}

func (d *Day01) Part1() any {
	return lo.Max(d.elves)
}

func (d *Day01) Part2() any {
	sorted := slices.Clone(d.elves)
	slices.Sort(sorted)

	return lo.Sum(lo.Subset(sorted, -3, 3))
}
