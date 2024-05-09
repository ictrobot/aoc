package day18

import (
	_ "embed"
	"github.com/ictrobot/aoc-go/internal/util/parse"
	"github.com/ictrobot/aoc-go/internal/util/vec"
	"math"
)

//go:embed example
var Example string

type Day18 struct {
	cubes    []vec.I3[int]
	grid     vec.I3Set[int]
	min, max int
}

func (d *Day18) Parse(input string) {
	d.cubes = parse.MustReflect[[]vec.I3[int]](parse.IntStrings(input))

	d.grid = make(vec.I3Set[int])
	d.min = math.MaxInt
	d.max = math.MinInt
	for _, c := range d.cubes {
		d.grid.Add(c)
		d.min = min(d.min, c.X, c.Y, c.Z)
		d.max = max(d.max, c.X, c.Y, c.Z)
	}
}

func (d *Day18) ParseExample() {
	d.Parse(Example)
}

func (d *Day18) Part1() any {
	surfaceArea := 0
	for _, c := range d.cubes {
		for _, dir := range vec.I3Directions {
			if !d.grid.Contains(c.Add(dir)) {
				surfaceArea++
			}
		}
	}
	return surfaceArea
}

func (d *Day18) Part2() any {
	exteriorArea := 0
	visited := make(vec.I3Set[int])
	queue := []vec.I3[int]{{d.min - 1, d.min - 1, d.min - 1}}

	for len(queue) > 0 {
		v := queue[0]
		queue = queue[1:]

		if !visited.Add(v) {
			continue
		}

		for _, dir := range vec.I3Directions {
			n := v.Add(dir)
			if d.grid.Contains(n) {
				exteriorArea++

			} else if n.X >= d.min-1 && n.X <= d.max+1 &&
				n.Y >= d.min-1 && n.Y <= d.max+1 &&
				n.Z >= d.min-1 && n.Z <= d.max+1 {

				queue = append(queue, n)
			}
		}
	}

	return exteriorArea
}
