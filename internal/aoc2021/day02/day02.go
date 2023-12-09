package day02

import (
	_ "embed"
	"fmt"
	"github.com/ictrobot/aoc/internal/util/parse"
)

//go:embed example
var Example string

type Day02 struct {
	commands []command
}

type command struct {
	Operation string
	Amount    int
}

func (d *Day02) Parse(input string) {
	d.commands = parse.MustReflect[[]command](parse.Whitespace(input))
}

func (d *Day02) ParseExample() {
	d.Parse(Example)
}

func (d *Day02) Part1() any {
	var horizontal, depth int
	for _, c := range d.commands {
		switch c.Operation {
		case "forward":
			horizontal += c.Amount
		case "down":
			depth += c.Amount
		case "up":
			depth -= c.Amount
		default:
			panic(fmt.Errorf("unknown operation: %s", c.Operation))
		}
	}
	return horizontal * depth
}

func (d *Day02) Part2() any {
	var horizontal, depth, aim int
	for _, c := range d.commands {
		switch c.Operation {
		case "forward":
			horizontal += c.Amount
			depth += aim * c.Amount
		case "down":
			aim += c.Amount
		case "up":
			aim -= c.Amount
		default:
			panic(fmt.Errorf("unknown operation: %s", c.Operation))
		}
	}
	return horizontal * depth
}
