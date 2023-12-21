package day21

import (
	_ "embed"
	"github.com/ictrobot/aoc/internal/util/numbers"
	"github.com/ictrobot/aoc/internal/util/parse"
	"github.com/ictrobot/aoc/internal/util/vec"
	"math"
)

//go:embed example
var Example string

type Day21 struct {
	grid    [][]bool
	start   vec.I2[int]
	example bool
}

func (d *Day21) Parse(input string) {
	lines := parse.Lines(input)

	d.grid = make([][]bool, len(lines))
	d.start = vec.I2[int]{}
	d.example = false

	for y, l := range lines {
		d.grid[y] = make([]bool, len(l))
		for x, c := range l {
			if c == 'S' {
				d.start.X = x
				d.start.Y = y
			}

			d.grid[y][x] = c == '#'
		}
	}
}

func (d *Day21) ParseExample() {
	d.Parse(Example)
	d.example = true
}

func (d *Day21) Part1() any {
	steps := 64
	if d.example {
		steps = 6
	}

	plots := d.plots(steps)
	return plots[steps]
}

func (d *Day21) Part2() any {
	if d.example {
		return "part 2 unsupported for example input"
	}
	// this only works because the start position is on a blank row & col

	steps := 26501365
	repeatAfter := len(d.grid) + len(d.grid[0])
	repeats := steps / repeatAfter
	leftOver := steps % repeatAfter

	plots := d.plots(leftOver + repeatAfter + repeatAfter)

	a, b, c := fitQuadratic([3]float64{0, 1, 2}, [3]float64{
		float64(plots[leftOver]),
		float64(plots[leftOver+repeatAfter]),
		float64(plots[leftOver+repeatAfter+repeatAfter]),
	})

	return evalQuadratic(a, b, c, repeats)
}

func (d *Day21) plots(steps int) []int64 {
	// use Z to store the step the position is reached first
	queue := []vec.I3[int]{d.start.WithZ(0)}
	visited := vec.Grid[bool]{}
	plots := make([]int64, steps+1)

	for len(queue) > 0 {
		p := queue[0]
		queue = queue[1:]

		if !visited.SetIfZero(p.XY(), true) {
			continue
		}

		for step := p.Z; step <= steps; step += 2 {
			plots[step]++
		}

		if p.Z >= steps {
			continue
		}

		for _, dir := range vec.I2Directions {
			n := p.Add(dir.WithZ(1))

			if d.grid[numbers.IntMod(n.Y, len(d.grid))][numbers.IntMod(n.X, len(d.grid[0]))] {
				continue
			}

			if visited.Get(n.XY()) {
				continue
			}

			queue = append(queue, n)
		}
	}

	return plots
}

func fitQuadratic(x, y [3]float64) (a, b, c float64) {
	a = (x[0]*(y[2]-y[1]) + x[1]*(y[0]-y[2]) + x[2]*(y[1]-y[0])) /
		((x[0] - x[1]) * (x[0] - x[2]) * (x[1] - x[2]))
	b = ((y[1] - y[0]) / (x[1] - x[0])) - a*(x[0]+x[1])
	c = y[0] - (a * x[0] * x[0]) - (b * x[0])
	return
}

func evalQuadratic(a, b, c float64, steps int) int64 {
	x := float64(steps)
	f := a*x*x + b*x + c
	return int64(math.Round(f))
}
