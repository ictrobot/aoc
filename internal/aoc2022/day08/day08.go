package day08

import (
	_ "embed"
	"github.com/ictrobot/aoc-go/internal/util/parse"
)

//go:embed example
var Example string

type Day08 struct {
	grid [][]int
}

func (d *Day08) Parse(input string) {
	d.grid = parse.Grid(parse.Lines, parse.ExtractDigits, input)
}

func (d *Day08) ParseExample() {
	d.Parse(Example)
}

func (d *Day08) Part1() any {
	visible := 0
	for y := 0; y < len(d.grid); y++ {
		for x := 0; x < len(d.grid[y]); x++ {
			if d.isVisible(x, y) {
				visible++
			}
		}
	}
	return visible
}

func (d *Day08) Part2() any {
	best := 0
	for y := 0; y < len(d.grid); y++ {
		for x := 0; x < len(d.grid[y]); x++ {
			best = max(best, d.scenicScore(x, y))
		}
	}
	return best
}

func (d *Day08) isVisible(x, y int) bool {
	h := d.grid[y][x]

	visible := true
	for y2 := y + 1; y2 < len(d.grid); y2++ {
		if d.grid[y2][x] >= h {
			visible = false
			break
		}
	}
	if visible {
		return true
	}

	visible = true
	for y2 := y - 1; y2 >= 0; y2-- {
		if d.grid[y2][x] >= h {
			visible = false
			break
		}
	}
	if visible {
		return true
	}

	visible = true
	for x2 := x + 1; x2 < len(d.grid[y]); x2++ {
		if d.grid[y][x2] >= h {
			visible = false
			break
		}
	}
	if visible {
		return true
	}

	visible = true
	for x2 := x - 1; x2 >= 0; x2-- {
		if d.grid[y][x2] >= h {
			visible = false
			break
		}
	}
	return visible
}

func (d *Day08) scenicScore(x, y int) int {
	h := d.grid[y][x]

	up := 0
	for y2 := y + 1; y2 < len(d.grid); y2++ {
		up++
		if d.grid[y2][x] >= h {
			break
		}
	}

	down := 0
	for y2 := y - 1; y2 >= 0; y2-- {
		down++
		if d.grid[y2][x] >= h {
			break
		}
	}

	right := 0
	for x2 := x + 1; x2 < len(d.grid[y]); x2++ {
		right++
		if d.grid[y][x2] >= h {
			break
		}
	}

	left := 0
	for x2 := x - 1; x2 >= 0; x2-- {
		left++
		if d.grid[y][x2] >= h {
			break
		}
	}
	return up * down * left * right
}
