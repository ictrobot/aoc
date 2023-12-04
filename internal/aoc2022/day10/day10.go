package day10

import (
	_ "embed"
	"github.com/ictrobot/aoc/internal/util/numbers"
	"github.com/ictrobot/aoc/internal/util/parse"
	"strings"
)

//go:embed example
var Example string

type Day10 struct {
	Instructions []struct {
		Opcode  string
		Operand *int
	}
}

func (d *Day10) Parse(input string) {
	*d = parse.MustReflect[Day10](parse.Whitespace(input))
}

func (d *Day10) ParseExample() {
	d.Parse(Example)
}

func (d *Day10) Part1() any {
	xs := d.runProgram()

	signalStrength := 0
	for i := 19; i < len(xs); i += 40 {
		signalStrength += (i + 1) * xs[i]
	}
	return signalStrength
}

func (d *Day10) Part2() any {
	xs := d.runProgram()

	var output strings.Builder
	for i, x := range xs {
		if numbers.IntAbsDiff(i%40, x) <= 1 {
			output.WriteByte('#')
		} else {
			output.WriteByte('.')
		}

		if i%40 == 39 && i < len(xs)-1 {
			output.WriteByte('\n')
		}
	}

	return output.String()
}

func (d *Day10) runProgram() []int {
	x := 1

	var xs []int
	for _, i := range d.Instructions {
		xs = append(xs, x)
		if i.Opcode == "addx" {
			xs = append(xs, x)
			x += *i.Operand
		}
	}

	return xs
}
