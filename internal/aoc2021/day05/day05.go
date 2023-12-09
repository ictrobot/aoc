package day05

import (
	_ "embed"
	"github.com/ictrobot/aoc/internal/util/parse"
	"github.com/ictrobot/aoc/internal/util/vec"
)

//go:embed example
var Example string

type Day05 struct {
	lines []line
}

type line struct {
	Start, End vec.I2[int]
}

func (d *Day05) Parse(input string) {
	d.lines = parse.MustReflect[[]line](parse.UintStrings(input))
}

func (d *Day05) ParseExample() {
	d.Parse(Example)
}

func (d *Day05) Part1() any {
	return d.calcOverlaps(true)
}

func (d *Day05) Part2() any {
	return d.calcOverlaps(false)
}

func (d *Day05) calcOverlaps(skipDiagonals bool) int {
	g := &vec.Grid[uint8]{}
	overlaps := 0
	for _, l := range d.lines {
		if skipDiagonals && l.Start.X != l.End.X && l.Start.Y != l.End.Y {
			continue
		}

		delta := l.End.Sub(l.Start)
		sign := delta.Sign()
		steps := delta.Abs().MaxComponent()

		pos := l.Start
		for i := 0; i <= steps; i++ {
			c := g.Get(pos)
			if c == 1 {
				overlaps++
			}
			g.Set(pos, c+1)
			pos = pos.Add(sign)
		}
	}
	return overlaps
}
