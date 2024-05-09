package day04

import (
	_ "embed"
	"github.com/ictrobot/aoc-go/internal/util/parse"
	"github.com/samber/lo"
)

//go:embed example
var Example string

type Day04 struct {
	elves [][]uint
}

func (d *Day04) Parse(input string) {
	d.elves = lo.Chunk(parse.ExtractUints(input), 4)
}

func (d *Day04) ParseExample() {
	d.Parse(Example)
}

func (d *Day04) Part1() any {
	overlap := 0
	for _, elf := range d.elves {
		if (elf[0] >= elf[2] && elf[1] <= elf[3]) || (elf[2] >= elf[0] && elf[3] <= elf[1]) {
			overlap++
		}
	}
	return overlap
}

func (d *Day04) Part2() any {
	overlap := 0
	for _, elf := range d.elves {
		if elf[0] <= elf[3] && elf[2] <= elf[1] {
			overlap++
		}
	}
	return overlap
}
