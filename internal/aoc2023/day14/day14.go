package day14

import (
	_ "embed"
	"github.com/ictrobot/aoc/internal/util/parse"
	"github.com/ictrobot/aoc/internal/util/structures"
)

//go:embed example
var Example string

type Day14 struct {
	grid *structures.FlatGrid[byte]
}

const part2Cycles = 1_000_000_000

const (
	rounded = 'O'
	empty   = '.'
)

func (d *Day14) Parse(input string) {
	d.grid = parse.ByteGrid(input)
}

func (d *Day14) ParseExample() {
	d.Parse(Example)
}

func (d *Day14) Part1() any {
	g := d.grid.Clone()
	rollNorth(g)
	return load(g)
}

func (d *Day14) Part2() any {
	g := d.grid.Clone()
	previous := make(map[string]int)
	for i := 1; i <= part2Cycles; i++ {
		rollNorth(g)
		rollWest(g)
		rollSouth(g)
		rollEast(g)

		s := string(g.S)
		if j, ok := previous[s]; ok {
			since := i - j
			remaining := part2Cycles - i
			i += remaining / since * since
			clear(previous)
		} else {
			previous[s] = i
		}
	}

	return load(g)
}

func rollNorth(g structures.FlatGrid[byte]) {
	for x := 0; x < g.Width; x++ {
		free := 0
		for y := 0; y < g.Height; y++ {
			t := g.Get(x, y)
			if t == empty {
				free++
			} else if t == rounded && free > 0 {
				g.Set(x, y-free, rounded)
				g.Set(x, y, empty)
			} else {
				free = 0
			}
		}
	}
}

func rollWest(g structures.FlatGrid[byte]) {
	for y := 0; y < g.Height; y++ {
		free := 0
		for x := 0; x < g.Width; x++ {
			t := g.Get(x, y)
			if t == empty {
				free++
			} else if t == rounded && free > 0 {
				g.Set(x-free, y, rounded)
				g.Set(x, y, empty)
			} else {
				free = 0
			}
		}
	}
}

func rollSouth(g structures.FlatGrid[byte]) {
	for x := 0; x < g.Width; x++ {
		free := 0
		for y := g.Height - 1; y >= 0; y-- {
			t := g.Get(x, y)
			if t == empty {
				free++
			} else if t == rounded && free > 0 {
				g.Set(x, y+free, rounded)
				g.Set(x, y, empty)
			} else {
				free = 0
			}
		}
	}
}

func rollEast(g structures.FlatGrid[byte]) {
	for y := 0; y < g.Height; y++ {
		free := 0
		for x := g.Width - 1; x >= 0; x-- {
			t := g.Get(x, y)
			if t == empty {
				free++
			} else if t == rounded && free > 0 {
				g.Set(x+free, y, rounded)
				g.Set(x, y, empty)
			} else {
				free = 0
			}
		}
	}
}

func load(g structures.FlatGrid[byte]) (total int) {
	for x := 0; x < g.Width; x++ {
		for y := 0; y < g.Height; y++ {
			if g.Get(x, y) == rounded {
				total += g.Height - y
			}
		}
	}
	return
}
