package day23

import (
	_ "embed"
	"github.com/ictrobot/aoc-go/internal/util/parse"
	"github.com/ictrobot/aoc-go/internal/util/vec"
	"math"
	"slices"
)

//go:embed example
var Example string

type Day23 struct {
	elves []vec.I2[int]
	grid  *vec.Grid[bool]
}

const (
	north = iota
	south
	west
	east
	dirCount
)

var dirVecs = [dirCount]vec.I2[int]{{0, 1}, {0, -1}, {-1, 0}, {1, 0}}

func (d *Day23) Parse(input string) {
	d.elves = nil
	d.grid = new(vec.Grid[bool])

	lines := parse.Lines(input)
	for y, l := range lines {
		for x, c := range l {
			if c == '#' {
				e := vec.I2[int]{x, len(lines) - 1 - y}
				d.elves = append(d.elves, e)
				d.grid.Set(e, true)
			}
		}
	}
}

func (d *Day23) ParseExample() {
	d.Parse(Example)
}

func (d *Day23) Part1() any {
	g, _ := d.simulate(10)
	return g.Count(false)
}

func (d *Day23) Part2() any {
	_, r := d.simulate(math.MaxInt)
	return r
}

func (d *Day23) simulate(maxRounds int) (*vec.Grid[bool], int) {
	elves := slices.Clone(d.elves)
	grid := d.grid.Clone()

	r := 0
	for ; r < maxRounds; r++ {
		newElves := make([]vec.I2[int], 0, len(elves))
		newGrid := vec.NewGrid[bool](grid.Bounds())

		// moves[to] = [from1, from2, ...]
		moves := make(map[vec.I2[int]][]vec.I2[int])

		for _, e := range elves {
			var blocked [dirCount]bool
			for _, dir := range vec.I2DirectionsWithDiagonals {
				if !grid.Contains(e.Add(dir)) {
					continue
				}

				if dir.Y > 0 {
					blocked[north] = true
				} else if dir.Y < 0 {
					blocked[south] = true
				}

				if dir.X > 0 {
					blocked[east] = true
				} else if dir.X < 0 {
					blocked[west] = true
				}
			}

			if !(blocked[north] || blocked[south] || blocked[east] || blocked[west]) ||
				(blocked[north] && blocked[south] && blocked[east] && blocked[west]) {
				// either no adjacent elves or completely blocked, so not moving
				newElves = append(newElves, e)
				newGrid.Set(e, true)
				continue
			}

			var moveTo vec.I2[int]
			for i := 0; i < dirCount; i++ {
				if !blocked[(i+r)%dirCount] {
					moveTo = e.Add(dirVecs[(i+r)%dirCount])
					break
				}
			}

			moves[moveTo] = append(moves[moveTo], e)
		}

		moved := false
		for moveTo, proposers := range moves {
			if len(proposers) == 1 {
				newElves = append(newElves, moveTo)
				newGrid.Set(moveTo, true)
				moved = true
				continue
			}

			newElves = append(newElves, proposers...)
			for _, e := range proposers {
				newGrid.Set(e, true)
			}
		}

		if !moved {
			break
		}

		elves = newElves
		grid = newGrid
	}

	return grid.Resize(grid.NonZeroBounds()), r + 1
}
