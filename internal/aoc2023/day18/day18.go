package day18

import (
	_ "embed"
	"github.com/ictrobot/aoc-go/internal/util/numbers"
	"github.com/ictrobot/aoc-go/internal/util/parse"
	"github.com/ictrobot/aoc-go/internal/util/vec"
	"strconv"
	"strings"
)

//go:embed example
var Example string

type Day18 struct {
	steps1 []vec.I2[int64]
	steps2 []vec.I2[int64]
}

func (d *Day18) Parse(input string) {
	input = strings.NewReplacer("(", "", ")", "", "#", "").Replace(input)
	lines := parse.Lines(input)

	d.steps1 = make([]vec.I2[int64], len(lines))
	d.steps2 = make([]vec.I2[int64], len(lines))
	for i, l := range lines {
		s := strings.SplitN(l, " ", 3)

		// part 1 uses first 2 fields
		switch s[0] {
		case "U":
			d.steps1[i].Y = 1
		case "D":
			d.steps1[i].Y = -1
		case "L":
			d.steps1[i].X = -1
		case "R":
			d.steps1[i].X = 1
		default:
			panic("unknown dir")
		}
		d.steps1[i] = d.steps1[i].Mul(parse.Int64(s[1]))

		// part 2 uses last field
		switch s[2][5] {
		case '3':
			d.steps2[i].Y = 1
		case '1':
			d.steps2[i].Y = -1
		case '2':
			d.steps2[i].X = -1
		case '0':
			d.steps2[i].X = 1
		default:
			panic("unknown dir")
		}

		c, err := strconv.ParseUint(s[2][:5], 16, 20)
		if err != nil {
			panic(err)
		}
		d.steps2[i] = d.steps2[i].Mul(int64(c))
	}
}

func (d *Day18) ParseExample() {
	d.Parse(Example)
}

func (d *Day18) Part1() any {
	return shapeArea(d.steps1)
}

func (d *Day18) Part2() any {
	return shapeArea(d.steps2)
}

func shapeArea(steps []vec.I2[int64]) int64 {
	// See 2023 Day 10 - Pick's theorem & shoelace formula
	var pos vec.I2[int64]

	var twiceArea, perimeter int64
	for _, s := range steps {
		next := pos.Add(s)

		twiceArea += pos.X*next.Y - next.X*pos.Y
		perimeter += s.Abs().MaxComponent()

		pos = next
	}

	return numbers.IntAbs(twiceArea)/2 + (perimeter / 2) + 1
}
