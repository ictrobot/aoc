package day14

import (
	_ "embed"
	"github.com/ictrobot/aoc/internal/util/parse"
	"github.com/ictrobot/aoc/internal/util/vec"
	"math"
)

//go:embed example
var Example string

var source = vec.I2[int]{500, 0}

type Day14 struct {
	Grid *vec.Grid[mat]
	maxY int
}

type mat uint8

const (
	rock mat = '#'
	sand mat = 'o'
)

func (d *Day14) Parse(input string) {
	d.Grid = new(vec.Grid[mat])
	d.maxY = math.MinInt

	for _, l := range parse.Lines(input) {
		ints := parse.ExtractInts(l)
		for i := 0; i < len(ints)-3; i += 2 {
			from := vec.I2[int]{ints[i], ints[i+1]}
			to := vec.I2[int]{ints[i+2], ints[i+3]}

			d.maxY = max(d.maxY, from.Y, to.Y)

			s := to.Sub(from).Sign()
			d.Grid.Set(from, rock)
			for from != to {
				from = from.Add(s)
				d.Grid.Set(from, rock)
			}
		}
	}

	// fmt.Println(d.Grid.PrettyPrint(1, false, true))
}

func (d *Day14) ParseExample() {
	d.Parse(Example)
}

func (d *Day14) Part1() any {
	grid := d.Grid.Clone()
	i := 0
	for dropSand(grid, d.maxY, math.MaxInt) {
		i++
	}
	// fmt.Println(grid.PrettyPrint(1, false, true))
	return i
}

func (d *Day14) Part2() any {
	grid := d.Grid.Clone()
	i := 0
	for !grid.Contains(source) {
		i++
		dropSand(grid, math.MaxInt, d.maxY+2)
	}
	// fmt.Println(grid.PrettyPrint(1, false, false))
	return i
}

func dropSand(filled *vec.Grid[mat], abyssY, floorY int) bool {
	v := source
iteration:
	for {
		if v.Y > abyssY {
			return false
		}

		for _, offset := range []vec.I2[int]{{0, 1}, {-1, 1}, {1, 1}} {
			next := v.Add(offset)
			if !filled.Contains(next) && next.Y < floorY {
				v = next
				continue iteration
			}
		}

		filled.Set(v, sand)
		return true
	}
}
