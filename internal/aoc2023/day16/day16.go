package day16

import (
	_ "embed"
	"github.com/ictrobot/aoc/internal/util/parse"
	"github.com/ictrobot/aoc/internal/util/structures"
)

//go:embed example
var Example string

type Day16 struct {
	grid *structures.FlatGrid[byte]
}

const (
	up = 1 << iota
	right
	down
	left
)

func (d *Day16) Parse(input string) {
	d.grid = parse.ByteGrid(input)
}

func (d *Day16) ParseExample() {
	d.Parse(Example)
}

func (d *Day16) Part1() any {
	beam := structures.NewFlatGrid[uint8](d.grid.Size())
	return d.energize(0, 0, right, beam)
}

func (d *Day16) Part2() any {
	beam := structures.NewFlatGrid[uint8](d.grid.Size())

	var best int
	for x := 0; x < d.grid.Width; x++ {
		beam.Clear()
		best = max(best, d.energize(x, 0, down, beam))

		beam.Clear()
		best = max(best, d.energize(x, d.grid.Height-1, up, beam))
	}
	for y := 0; y < d.grid.Height; y++ {
		beam.Clear()
		best = max(best, d.energize(0, y, right, beam))

		beam.Clear()
		best = max(best, d.energize(d.grid.Width-1, y, left, beam))
	}
	return best
}

func (d *Day16) energize(x, y int, dir uint8, beam *structures.FlatGrid[uint8]) (count int) {
	for d.grid.InBounds(x, y) {
		if v := beam.Get(x, y); v&dir != 0 {
			break
		} else {
			if v == 0 {
				count++
			}
			beam.Set(x, y, v|dir)
		}

		switch d.grid.Get(x, y) {
		case '.':
			// carry on
		case '/':
			switch dir {
			case up:
				dir = right
			case right:
				dir = up
			case down:
				dir = left
			case left:
				dir = down
			}
		case '\\':
			switch dir {
			case up:
				dir = left
			case right:
				dir = down
			case down:
				dir = right
			case left:
				dir = up
			}
		case '|':
			if dir == left || dir == right {
				count += d.energize(x, y-1, up, beam)
				dir = down
			}
		case '-':
			if dir == down || dir == up {
				count += d.energize(x-1, y, left, beam)
				dir = right
			}
		default:
			panic("unknown tile")
		}

		switch dir {
		case up:
			y--
		case right:
			x++
		case down:
			y++
		case left:
			x--
		}
	}
	return
}
