package day01

import (
	_ "embed"
	"github.com/ictrobot/aoc/internal/util/parse"
)

//go:embed example
var Example string

type Day01 struct {
	depths []int
}

func (d *Day01) Parse(input string) {
	d.depths = parse.ExtractInts(input)
}

func (d *Day01) ParseExample() {
	d.Parse(Example)
}

func (d *Day01) Part1() any {
	count := 0
	for i := 1; i < len(d.depths); i++ {
		if d.depths[i] > d.depths[i-1] {
			count++
		}
	}
	return count
}

func (d *Day01) Part2() any {
	count := 0
	for i := 3; i < len(d.depths); i++ {
		if d.depths[i]+d.depths[i-1]+d.depths[i-2] > d.depths[i-1]+d.depths[i-2]+d.depths[i-3] {
			count++
		}
	}
	return count
}
