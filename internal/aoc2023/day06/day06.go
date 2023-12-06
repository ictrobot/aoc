package day06

import (
	_ "embed"
	"github.com/ictrobot/aoc/internal/util/parse"
	"strings"
)

//go:embed example
var Example string

type Day06 struct {
	races []race
	part2 race
}

type race struct {
	time, dist int
}

func (d *Day06) Parse(input string) {
	lines := parse.Lines(input)

	d.races = nil
	times := parse.ExtractInts(lines[0])
	distances := parse.ExtractInts(lines[1])
	for i := range times {
		d.races = append(d.races, race{times[i], distances[i]})
	}

	p2 := parse.ExtractInts(strings.ReplaceAll(input, " ", ""))
	d.part2 = race{p2[0], p2[1]}
}

func (d *Day06) ParseExample() {
	d.Parse(Example)
}

func (d *Day06) Part1() any {
	ans := 1
	for _, r := range d.races {
		wins := 0
		for h := 1; h < r.time; h++ {
			if h*(r.time-h) > r.dist {
				wins++
			}
		}
		ans *= wins
	}
	return ans
}

func (d *Day06) Part2() any {
	wins := 0
	for h := 1; h < d.part2.time; h++ {
		if h*(d.part2.time-h) > d.part2.dist {
			wins++
		}
	}
	return wins
}
