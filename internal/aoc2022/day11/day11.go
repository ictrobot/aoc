package day11

import (
	_ "embed"
	"github.com/ictrobot/aoc/internal/util/deep"
	"github.com/ictrobot/aoc/internal/util/parse"
	"regexp"
	"slices"
)

//go:embed example
var Example string

var regex = regexp.MustCompile(`\d+|old|\+|\*`)

type Day11 struct {
	Monkeys []monkey
}

type monkey struct {
	_           parse.Placeholder `regex:"^[0-9]+$"` // Monkey _:
	Items       []int
	_           parse.Placeholder `match:"old"` // old on LHS
	Operation   string
	Operand     *int              // nil for old
	_           parse.Placeholder `match:"old" flags:"optional"` // old on RHS
	Test        int
	TrueMonkey  int
	FalseMonkey int
}

func (d *Day11) Parse(input string) {
	*d = parse.MustReflect[Day11](regex.FindAllString(input, -1))
}

func (d *Day11) ParseExample() {
	d.Parse(Example)
}

func (d *Day11) Part1() any {
	return d.monkeyBusiness(20, 3)
}

func (d *Day11) Part2() any {
	return d.monkeyBusiness(10_000, 1)
}

func (d *Day11) monkeyBusiness(rounds int, divisor int) int {
	monkeys := *deep.Clone(&d.Monkeys)
	inspected := make([]int, len(monkeys))

	mod := 1
	for _, m := range monkeys {
		mod *= m.Test
	}

	for round := 0; round < rounds; round++ {
		for i := range monkeys {
			items := monkeys[i].Items
			monkeys[i].Items = nil

			for _, item := range items {
				inspected[i]++

				operand := item
				if monkeys[i].Operand != nil {
					operand = *monkeys[i].Operand
				}

				if monkeys[i].Operation == "*" {
					item *= operand
				} else {
					item += operand
				}

				item /= divisor
				item %= mod

				next := monkeys[i].FalseMonkey
				if item%monkeys[i].Test == 0 {
					next = monkeys[i].TrueMonkey
				}
				monkeys[next].Items = append(monkeys[next].Items, item)
			}
		}
	}

	slices.Sort(inspected)
	return inspected[len(inspected)-1] * inspected[len(inspected)-2]
}
