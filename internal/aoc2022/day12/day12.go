package day12

import (
	_ "embed"
	"github.com/ictrobot/aoc/internal/util/parse"
	"github.com/ictrobot/aoc/internal/util/structures"
	"github.com/ictrobot/aoc/internal/util/vec"
)

//go:embed example
var Example string

type Day12 struct {
	grid       [][]int
	start, end vec.I2[int]
	zeroes     []vec.I2[int]
}

func (d *Day12) Parse(input string) {
	d.grid = nil
	d.zeroes = nil
	for y, l := range parse.Lines(input) {
		var row []int
		for x, c := range l {
			var h int
			switch c {
			case 'S':
				d.start.X = x
				d.start.Y = y
				h = 0
			case 'E':
				d.end.X = x
				d.end.Y = y
				h = 25
			default:
				h = int(c - 'a')
			}

			if h == 0 {
				d.zeroes = append(d.zeroes, vec.I2[int]{x, y})
			}

			row = append(row, h)
		}
		d.grid = append(d.grid, row)
	}
}

func (d *Day12) ParseExample() {
	d.Parse(Example)

}

func (d *Day12) Part1() any {
	return d.shortestPath(d.start)
}

func (d *Day12) Part2() any {
	return d.shortestPath(d.zeroes...)
}

func (d *Day12) shortestPath(starts ...vec.I2[int]) int {
	var queue structures.BucketQueue[vec.I3[int]]
	for _, s := range starts {
		// store distance in Z
		queue.Push(s.WithZ(0), 0)
	}

	visited := vec.Grid[bool]{}
	for !queue.IsEmpty() {
		v := queue.Pop()

		if !visited.SetIfZero(v.XY(), true) {
			continue
		}

		if v.XY() == d.end {
			return v.Z
		}

		height := d.grid[v.Y][v.X]
		for _, dir := range vec.I2Directions {
			n := v.Add(dir.WithZ(1)) // add one to distance

			if n.Y < 0 || n.Y >= len(d.grid) || n.X < 0 || n.X >= len(d.grid[n.Y]) {
				continue
			}

			h := d.grid[n.Y][n.X]
			if h <= height+1 {
				queue.Push(n, n.Z)
			}
		}
	}

	panic("no route found")
}
