package day16

import (
	_ "embed"
	"github.com/ictrobot/aoc/internal/util/parse"
)

//go:embed example
var Example string

type Day16 struct {
	grid []string
}

const (
	up = 1 << iota
	right
	down
	left
)

func (d *Day16) Parse(input string) {
	d.grid = parse.Lines(input)
}

func (d *Day16) ParseExample() {
	d.Parse(Example)
}

func (d *Day16) Part1() any {
	beam := make([]uint8, len(d.grid)*len(d.grid[0]))
	return d.energize(0, 0, right, beam)
}

func (d *Day16) Part2() any {
	beam := make([]uint8, len(d.grid)*len(d.grid[0]))

	var best int
	for x := 0; x < len(d.grid[0]); x++ {
		clear(beam)
		best = max(best, d.energize(x, 0, down, beam))

		clear(beam)
		best = max(best, d.energize(x, len(d.grid)-1, up, beam))
	}
	for y := 0; y < len(d.grid[1]); y++ {
		clear(beam)
		best = max(best, d.energize(0, y, right, beam))

		clear(beam)
		best = max(best, d.energize(len(d.grid[0])-1, y, left, beam))
	}
	return best
}

func (d *Day16) energize(x, y int, dir uint8, beam []uint8) (count int) {
	for x >= 0 && y >= 0 && x < len(d.grid[0]) && y < len(d.grid) {
		idx := (x * len(d.grid[0])) + y
		if v := beam[idx]; v&dir != 0 {
			break
		} else {
			if v == 0 {
				count++
			}
			beam[idx] = v | dir
		}

		switch d.grid[y][x] {
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
