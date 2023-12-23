package day23

import (
	_ "embed"
	"github.com/ictrobot/aoc/internal/util/parse"
	"github.com/ictrobot/aoc/internal/util/vec"
	"strings"
)

//go:embed example
var Example string

type Day23 struct {
	start, end vec.I2[int]
	grid       [][]tile
}

type tile uint8

const (
	path   tile = '.'
	forest tile = '#'
	slopeN tile = '^'
	slopeE tile = '>'
	slopeS tile = 'v'
	slopeW tile = '<'
)

func (d *Day23) Parse(input string) {
	lines := parse.Lines(input)

	d.start = vec.I2[int]{strings.IndexByte(lines[0], '.'), 0}
	d.end = vec.I2[int]{strings.IndexByte(lines[len(lines)-1], '.'), len(lines) - 1}

	d.grid = make([][]tile, len(lines))
	for y, l := range lines {
		d.grid[y] = make([]tile, len(l))
		for x, c := range l {
			d.grid[y][x] = tile(c)
		}
	}
}

func (d *Day23) ParseExample() {
	d.Parse(Example)
}

func (d *Day23) Part1() any {
	return d.longestPath(false)
}

func (d *Day23) Part2() any {
	return d.longestPath(true)
}

func (d *Day23) longestPath(ignoreSlopes bool) int {
	visited := make([][]bool, len(d.grid))
	for i := 0; i < len(d.grid); i++ {
		visited[i] = make([]bool, len(d.grid[0]))
	}

	return d.dfs(d.start, 0, visited, ignoreSlopes)
}

func (d *Day23) dfs(p vec.I2[int], length int, visited [][]bool, ignoreSlopes bool) int {
	if p == d.end {
		return length
	}
	visited[p.Y][p.X] = true

	best := 0
	currentTile := d.grid[p.Y][p.X]
	for _, dir := range vec.I2Directions {
		if !ignoreSlopes {
			if (currentTile == slopeN && dir.Y >= 0) || // N y = -1
				(currentTile == slopeE && dir.X <= 0) || // E x = 1
				(currentTile == slopeS && dir.Y <= 0) || // S y = 1
				(currentTile == slopeW && dir.X >= 0) { // W x = -1
				continue
			}
		}

		n := p.Add(dir)
		if n.Y < 0 || n.Y >= len(d.grid) || n.X < 0 || n.X >= len(d.grid[n.Y]) {
			continue
		}

		if visited[n.Y][n.X] {
			continue
		}

		if d.grid[n.Y][n.X] == forest {
			continue
		}

		best = max(best, d.dfs(n, length+1, visited, ignoreSlopes))
	}

	visited[p.Y][p.X] = false
	return best
}
