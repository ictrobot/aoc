package day03

import (
	_ "embed"
	"github.com/ictrobot/aoc/internal/util/parse"
	"github.com/samber/lo"
	"strings"
)

//go:embed example
var Example string

type Day03 struct {
	rucksacks []string
}

func (d *Day03) Parse(input string) {
	d.rucksacks = parse.Lines(input)
}

func (d *Day03) ParseExample() {
	d.Parse(Example)
}

func (d *Day03) Part1() any {
	priorities := 0
	for _, rucksack := range d.rucksacks {
		first := lo.Uniq(strings.Split(rucksack[:len(rucksack)/2], ""))
		second := lo.Uniq(strings.Split(rucksack[len(rucksack)/2:], ""))

		priorities += priority(lo.Intersect(first, second)[0][0])
	}
	return priorities
}

func (d *Day03) Part2() any {
	priorities := 0
	for _, rucksacks := range lo.Chunk(d.rucksacks, 3) {
		repeated := lo.Intersect(
			lo.Intersect(
				lo.Uniq(strings.Split(rucksacks[0], "")),
				lo.Uniq(strings.Split(rucksacks[1], "")),
			),
			lo.Uniq(strings.Split(rucksacks[2], "")),
		)[0][0]

		priorities += priority(repeated)
	}
	return priorities
}

func priority(c uint8) int {
	if c >= 'a' && c <= 'z' {
		return int(c - 'a' + 1)
	} else if c >= 'A' && c <= 'Z' {
		return int(c - 'A' + 27)
	}
	panic(c)
}
