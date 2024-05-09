package day05

import (
	_ "embed"
	"github.com/ictrobot/aoc-go/internal/util/collections"
	"github.com/ictrobot/aoc-go/internal/util/parse"
	"github.com/samber/lo"
)

//go:embed example
var Example string

type Day05 struct {
	startingStacks [][]string
	instructions   [][]int
}

func (d *Day05) Parse(input string) {
	chunks := parse.Chunks(input)

	startingLines := parse.Lines(chunks[0])
	d.startingStacks = make([][]string, len(parse.Whitespace(startingLines[len(startingLines)-1])))
	for _, line := range startingLines[:len(startingLines)-1] {
		for i := 0; 1+4*i < len(line); i++ {
			if line[1+4*i] != ' ' {
				d.startingStacks[i] = append([]string{string(line[1+4*i])}, d.startingStacks[i]...)
			}
		}
	}

	d.instructions = lo.Chunk(parse.ExtractInts(chunks[1]), 3)
}

func (d *Day05) ParseExample() {
	d.Parse(Example)
}

func (d *Day05) Part1() any {
	stacks := collections.Clone2D(d.startingStacks)
	for _, instruction := range d.instructions {
		f, t := instruction[1]-1, instruction[2]-1

		for i := 0; i < instruction[0]; i++ {
			stacks[t] = append(stacks[t], stacks[f][len(stacks[f])-1])
			stacks[f] = stacks[f][:len(stacks[f])-1]
		}
	}

	results := ""
	for _, stack := range stacks {
		results += stack[len(stack)-1]
	}
	return results
}

func (d *Day05) Part2() any {
	stacks := collections.Clone2D(d.startingStacks)
	for _, instruction := range d.instructions {
		c, f, t := instruction[0], instruction[1]-1, instruction[2]-1

		stacks[t] = append(stacks[t], stacks[f][len(stacks[f])-c:]...)
		stacks[f] = stacks[f][:len(stacks[f])-c]
	}

	results := ""
	for _, stack := range stacks {
		results += stack[len(stack)-1]
	}
	return results
}
