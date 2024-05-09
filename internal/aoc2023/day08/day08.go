package day08

import (
	_ "embed"
	"github.com/ictrobot/aoc-go/internal/util/numbers"
	"github.com/ictrobot/aoc-go/internal/util/parse"
	"strings"
)

//go:embed example
var Example string

//go:embed example2
var Example2 string

type Day08 struct {
	directions []string
	nodes      []*node
}

type node struct {
	name        string
	left, right *node
}

func (d *Day08) Parse(input string) {
	split := parse.Whitespace(strings.NewReplacer("=", "", "(", "", ",", "", ")", "").Replace(input))

	d.directions = parse.Characters(split[0])

	// using pointers to traverse left/right is far quicker
	d.nodes = nil
	indexes := make(map[string]int)
	for i := 1; i < len(split)-2; i += 3 {
		indexes[split[i]] = len(d.nodes)
		d.nodes = append(d.nodes, &node{name: split[i]})
	}
	for i := 1; i < len(split)-2; i += 3 {
		d.nodes[i/3].left = d.nodes[indexes[split[i+1]]]
		d.nodes[i/3].right = d.nodes[indexes[split[i+2]]]
	}
}

func (d *Day08) ParseExample() {
	d.Parse(Example)
}

func (d *Day08) ParseExample2() {
	d.Parse(Example2)
}

func (d *Day08) Part1() any {
	var at *node
	for _, n := range d.nodes {
		if n.name == "AAA" {
			at = n
			break
		}
	}
	if at.name == "" {
		panic("node AAA not found")
	}

	step := 0
	for at.name != "ZZZ" {
		dir := d.directions[step%len(d.directions)]
		if dir == "L" {
			at = at.left
		} else {
			at = at.right
		}
		step++
	}

	return step
}

func (d *Day08) Part2() any {
	var at []*node
	for _, n := range d.nodes {
		if n.name[len(n.name)-1] == 'A' {
			at = append(at, n)
		}
	}

	endSteps := make([]int, len(at))
	reachedEnd := 0

	step := 0
	for reachedEnd != len(at) {
		dir := d.directions[step%len(d.directions)]
		step++
		if dir == "L" {
			for i := 0; i < len(at); i++ {
				at[i] = at[i].left
			}
		} else {
			for i := 0; i < len(at); i++ {
				at[i] = at[i].right
			}
		}

		for i, n := range at {
			if n.name[len(n.name)-1] == 'Z' && endSteps[i] == 0 {
				endSteps[i] = step
				reachedEnd++
			}
		}
	}

	result := int64(1)
	for _, p := range endSteps {
		result = numbers.LCM(result, int64(p))
	}
	return result
}
