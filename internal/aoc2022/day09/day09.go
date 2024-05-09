package day09

import (
	_ "embed"
	"github.com/ictrobot/aoc-go/internal/util/parse"
	"github.com/ictrobot/aoc-go/internal/util/vec"
)

//go:embed example
var Example string

//go:embed example2
var Example2 string

type Day09 struct {
	Motions []struct {
		Dir   string
		Steps int
	}
}

func (d *Day09) Parse(input string) {
	*d = parse.MustReflect[Day09](parse.Whitespace(input))
}

func (d *Day09) ParseExample() {
	d.Parse(Example)
}

func (d *Day09) ParseExample2() {
	d.Parse(Example2)
}

func (d *Day09) Part1() any {
	return d.tailPositions(1)
}

func (d *Day09) Part2() any {
	return d.tailPositions(9)
}

func (d *Day09) tailPositions(numTails uint) int {
	h := vec.I2[int]{}
	tails := make([]vec.I2[int], numTails)
	visited := make(vec.I2Set[int])

	for _, m := range d.Motions {
		for i := 0; i < m.Steps; i++ {
			switch m.Dir {
			case "U":
				h.Y++
			case "D":
				h.Y--
			case "R":
				h.X++
			case "L":
				h.X--
			}

			prev := h
			for i := range tails {
				diff := prev.Sub(tails[i])
				if diff.Abs().MaxComponent() > 1 {
					tails[i] = tails[i].Add(diff.Sign())
				}
				prev = tails[i]
			}

			visited.Add(prev)
		}
	}

	return len(visited)
}
