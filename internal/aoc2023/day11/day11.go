package day11

import (
	_ "embed"
	"github.com/ictrobot/aoc/internal/util/parse"
	"github.com/ictrobot/aoc/internal/util/vec"
)

//go:embed example
var Example string

type Day11 struct {
	galaxies               []vec.I2[int64]
	nonEmptyXs, nonEmptyYs []bool
	example                bool
}

func (d *Day11) Parse(input string) {
	lines := parse.Lines(input)

	d.nonEmptyXs = make([]bool, len(lines[0]))
	d.nonEmptyYs = make([]bool, len(lines))

	d.galaxies = nil
	for y, l := range lines {
		for x, c := range l {
			if c != '#' {
				continue
			}

			d.galaxies = append(d.galaxies, vec.I2[int64]{int64(x), int64(y)})
			d.nonEmptyXs[x] = true
			d.nonEmptyYs[y] = true
		}
	}

	d.example = false
}

func (d *Day11) ParseExample() {
	d.Parse(Example)
	d.example = true
}

func (d *Day11) Part1() any {
	return sumGalaxyDistances(d.expandEmptySpace(2))
}

func (d *Day11) Part2() any {
	var galaxies []vec.I2[int64]
	if d.example {
		galaxies = d.expandEmptySpace(100)
	} else {
		galaxies = d.expandEmptySpace(1000000)
	}

	return sumGalaxyDistances(galaxies)
}

func (d *Day11) expandEmptySpace(factor int64) []vec.I2[int64] {
	addX := make([]int64, len(d.nonEmptyXs))
	var total int64
	for i, b := range d.nonEmptyXs {
		if !b {
			total += factor - 1
		}
		addX[i] = total
	}

	addY := make([]int64, len(d.nonEmptyYs))
	total = 0
	for i, b := range d.nonEmptyYs {
		if !b {
			total += factor - 1
		}
		addY[i] = total
	}

	out := make([]vec.I2[int64], len(d.galaxies))
	for i, g := range d.galaxies {
		out[i] = g.Add(vec.I2[int64]{addX[g.X], addY[g.Y]})
	}
	return out
}

func sumGalaxyDistances(galaxies []vec.I2[int64]) int64 {
	sum := int64(0)
	for i := 0; i < len(galaxies); i++ {
		for j := i + 1; j < len(galaxies); j++ {
			sum += galaxies[i].ManhattanDist(galaxies[j])
		}
	}
	return sum
}
