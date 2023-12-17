package day17

import (
	_ "embed"
	"github.com/ictrobot/aoc/internal/util/parse"
	"github.com/ictrobot/aoc/internal/util/structures"
	"github.com/ictrobot/aoc/internal/util/vec"
	"math"
)

//go:embed example
var Example string

//go:embed example2
var Example2 string

type Day17 struct {
	grid [][]int
}

type node struct {
	pos     vec.I2[int16]
	prevDir uint8
	loss    int
}

// order is important for bitwise operations in lowestLossPath
// (dir ^ 1) used to get opposite direction
// (dir & 2) used to check if horizontal or vertical
var directions = [4]vec.I2[int16]{
	{1, 0},
	{-1, 0},
	{0, 1},
	{0, -1},
}

func (d *Day17) Parse(input string) {
	d.grid = parse.Grid(parse.Lines, parse.ExtractDigits, input)
}

func (d *Day17) ParseExample() {
	d.Parse(Example)
}

func (d *Day17) ParseExample2() {
	d.Parse(Example2)
}

func (d *Day17) Part1() any {
	return d.lowestLossPath(0, 3)
}

func (d *Day17) Part2() any {
	return d.lowestLossPath(4, 10)
}

func (d *Day17) lowestLossPath(minLength, maxLength int) int {
	queue := structures.Heap[node]{LessThan: func(a, b node) bool {
		return a.loss < b.loss
	}}
	queue.Push(node{prevDir: math.MaxUint8})

	end := vec.I2[int16]{int16(len(d.grid[0]) - 1), int16(len(d.grid) - 1)}

	visited := structures.BitSet32[uint32]{}
	for !queue.IsEmpty() {
		n := queue.Pop()

		if n.pos == end {
			return n.loss
		}

		k := key(n.pos, n.prevDir)
		if visited.Has(k) {
			continue
		}
		visited.Set(k)

		for dir := uint8(0); dir < 4; dir++ {
			if dir == n.prevDir || dir^1 == n.prevDir {
				// cannot continue in previous direction or turn back
				continue
			}

			next := n.pos
			loss := 0
			for dist := 1; dist <= maxLength; dist++ {
				next = next.Add(directions[dir])
				if next.Y < 0 || next.Y > end.Y || next.X < 0 || next.X > end.X {
					break
				}

				loss += d.grid[next.Y][next.X]
				if dist < minLength {
					continue
				}

				if visited.Has(key(next, dir)) {
					continue
				}

				queue.Push(node{next, dir, n.loss + loss})
			}
		}
	}

	panic("no route found")
}

func key(pos vec.I2[int16], dir uint8) uint32 {
	// each position can be visited once horizontally and once vertically
	return uint32(uint16(pos.X))<<17 | uint32(uint16(pos.Y))<<1 | uint32(dir&2)>>1
}
