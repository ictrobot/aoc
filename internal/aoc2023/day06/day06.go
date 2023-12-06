package day06

import (
	_ "embed"
	"github.com/ictrobot/aoc/internal/util/parse"
	"math"
	"strings"
)

//go:embed example
var Example string

type Day06 struct {
	races []race
	part2 race
}

type race struct {
	time, dist int64
}

func (d *Day06) Parse(input string) {
	lines := parse.Lines(input)

	d.races = nil
	times := parse.ExtractInt64s(lines[0])
	distances := parse.ExtractInt64s(lines[1])
	for i := range times {
		d.races = append(d.races, race{times[i], distances[i]})
	}

	p2 := parse.ExtractInt64s(strings.ReplaceAll(input, " ", ""))
	d.part2 = race{p2[0], p2[1]}
}

func (d *Day06) ParseExample() {
	d.Parse(Example)
}

func (d *Day06) Part1() any {
	result := 1
	for _, r := range d.races {
		result *= r.winningWays()
	}
	return result
}

func (d *Day06) Part2() any {
	return d.part2.winningWays()
}

func (r race) winningWays() int {
	i := r.winningInterval()
	return i[1] - i[0] + 1
}

func (r race) winningInterval() [2]int {
	// x(T-x) > D
	// 0 > x**2 - Tx + D
	time := float64(r.time)
	dist := float64(r.dist)

	d := (time * time) - (4.0 * dist)
	if d <= 0 {
		panic("can't win")
	}
	d = math.Sqrt(d)

	// need the min integer > lower bound and max int < upper bound
	return [2]int{
		int(math.Floor((time-d)/2) + 1),
		int(math.Ceil((time+d)/2) - 1),
	}
}
