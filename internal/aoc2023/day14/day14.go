package day14

import (
	_ "embed"
	"github.com/ictrobot/aoc/internal/util/parse"
	"github.com/ictrobot/aoc/internal/util/vec"
)

//go:embed example
var Example string

type Day14 struct {
	grid *vec.Grid[tile]
}

const part2Cycles = 1_000_000_000

type tile uint8

const (
	rounded = 'O'
	empty   = '.'
)

func (d *Day14) Parse(input string) {
	lines := parse.Lines(input)

	d.grid = vec.NewGrid[tile](0, 0, len(lines[0])-1, len(lines)-1)
	for y, l := range lines {
		for x, c := range l {
			d.grid.SetInts(x, y, tile(c))
		}
	}
}

func (d *Day14) ParseExample() {
	d.Parse(Example)
}

func (d *Day14) Part1() any {
	g := d.grid.Clone()
	rollNorth(g)
	// fmt.Println(g.PrettyPrint(1, false, true))
	return load(g)
}

func (d *Day14) Part2() any {
	g := d.grid.Clone()
	previous := make(map[string]int)
	// TODO use a faster function than PrettyPrint!
	previous[g.PrettyPrint(1, false, true)] = 0

	for i := 1; i <= part2Cycles; i++ {
		rollNorth(g)
		rollWest(g)
		rollSouth(g)
		rollEast(g)

		p := g.PrettyPrint(1, false, true)
		if j, ok := previous[p]; ok {
			since := i - j
			remaining := part2Cycles - i
			i += remaining / since * since
			clear(previous)
		} else {
			previous[p] = i
		}
	}

	return load(g)
}

func rollNorth(g *vec.Grid[tile]) {
	yMin, xMin, xMax, yMax := g.Bounds()
	for y := yMin + 1; y <= yMax; y++ {
		for x := xMin; x <= xMax; x++ {
			if g.GetInts(x, y) != rounded || g.GetInts(x, y-1) != empty {
				continue
			}

			dstY := y - 1
			for dstY > yMin && g.GetInts(x, dstY-1) == empty {
				dstY--
			}

			g.SetInts(x, dstY, rounded)
			g.SetInts(x, y, empty)
		}
	}
}

func rollWest(g *vec.Grid[tile]) {
	yMin, xMin, xMax, yMax := g.Bounds()
	for x := xMin + 1; x <= xMax; x++ {
		for y := yMin; y <= yMax; y++ {
			if g.GetInts(x, y) != rounded || g.GetInts(x-1, y) != empty {
				continue
			}

			dstX := x - 1
			for dstX > xMin && g.GetInts(dstX-1, y) == empty {
				dstX--
			}

			g.SetInts(dstX, y, rounded)
			g.SetInts(x, y, empty)
		}
	}
}

func rollSouth(g *vec.Grid[tile]) {
	yMin, xMin, xMax, yMax := g.Bounds()
	for y := yMax - 1; y >= yMin; y-- {
		for x := xMin; x <= xMax; x++ {
			if g.GetInts(x, y) != rounded || g.GetInts(x, y+1) != empty {
				continue
			}

			dstY := y + 1
			for dstY < yMax && g.GetInts(x, dstY+1) == empty {
				dstY++
			}

			g.SetInts(x, dstY, rounded)
			g.SetInts(x, y, empty)
		}
	}
}

func rollEast(g *vec.Grid[tile]) {
	yMin, xMin, xMax, yMax := g.Bounds()
	for x := xMax - 1; x >= xMin; x-- {
		for y := yMin; y <= yMax; y++ {
			if g.GetInts(x, y) != rounded || g.GetInts(x+1, y) != empty {
				continue
			}

			dstX := x + 1
			for dstX < xMax && g.GetInts(dstX+1, y) == empty {
				dstX++
			}

			g.SetInts(dstX, y, rounded)
			g.SetInts(x, y, empty)
		}
	}
}

func load(g *vec.Grid[tile]) (total int) {
	yMin, xMin, xMax, yMax := g.Bounds()
	for y := yMin; y <= yMax; y++ {
		for x := xMin; x <= xMax; x++ {
			if g.GetInts(x, y) == rounded {
				total += (yMax - yMin + 1) - y
			}
		}
	}
	return
}
