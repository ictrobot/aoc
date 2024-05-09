package day23

import (
	_ "embed"
	"github.com/ictrobot/aoc-go/internal/util/collections"
	"github.com/ictrobot/aoc-go/internal/util/parse"
	"github.com/ictrobot/aoc-go/internal/util/vec"
	"strings"
)

//go:embed example
var Example string

type Day23 struct {
	start, end vec.I2[int]
	grid       [][]uint8
}

type vertex []edge

type edge struct {
	vertex, length int
}

const (
	path   = '.'
	forest = '#'
	slopeN = '^'
	slopeE = '>'
	slopeS = 'v'
	slopeW = '<'
)

func (d *Day23) Parse(input string) {
	lines := parse.Lines(input)

	d.start = vec.I2[int]{strings.IndexByte(lines[0], path), 0}
	d.end = vec.I2[int]{strings.IndexByte(lines[len(lines)-1], path), len(lines) - 1}

	d.grid = make([][]uint8, len(lines))
	for y, l := range lines {
		d.grid[y] = []uint8(l)
	}
}

func (d *Day23) ParseExample() {
	d.Parse(Example)
}

func (d *Day23) Part1() any {
	return longestPath(d.graph(false))
}

func (d *Day23) Part2() any {
	return longestPath(d.graph(true))
}

func (d *Day23) graph(ignoreSlopes bool) (int, int, []vertex) {
	vertexGrid := make([][]int, len(d.grid))
	for y := 0; y < len(d.grid); y++ {
		vertexGrid[y] = make([]int, len(d.grid[0]))
		collections.Fill(vertexGrid[y], -1)
	}

	// create initial vertices for start and end
	vertices := []vertex{{}, {}}
	vertexGrid[d.start.Y][d.start.X] = 0
	vertexGrid[d.end.Y][d.end.X] = 1
	possibleEdges := []vec.I3[int]{ // use Z to store next direction
		d.start.WithZ(vec.PosY),
		d.end.WithZ(vec.NegY),
	}

	// find all the junctions in the grid, creating a vertex for each one and
	// populating the list of possible edges to explore
	for y := 0; y < len(d.grid); y++ {
		for x := 0; x < len(d.grid[0]); x++ {
			if d.grid[y][x] == forest {
				continue
			}

			var neighbours, count uint8
			for i, dir := range vec.I2Directions {
				if y+dir.Y < 0 || y+dir.Y >= len(d.grid) || x+dir.X < 0 || x+dir.X >= len(d.grid[0]) {
					continue
				}

				if d.grid[y+dir.Y][x+dir.X] != forest {
					neighbours |= 1 << i
					count++
				}
			}

			if count <= 2 {
				continue
			}

			vertices = append(vertices, vertex{})
			vertexGrid[y][x] = len(vertices) - 1

			for i := range vec.I2Directions {
				if neighbours&(1<<i) != 0 {
					possibleEdges = append(possibleEdges, vec.I3[int]{x, y, i})
				}
			}
		}
	}

	// explore all possible directed edges
	for _, e := range possibleEdges {
		v1 := vertexGrid[e.Y][e.X]

		length := 0
		for {
			currentTile := d.grid[e.Y][e.X]
			nextDir := vec.I2Directions[e.Z]

			if !ignoreSlopes && currentTile != path {
				if (currentTile == slopeN && nextDir.Y >= 0) ||
					(currentTile == slopeE && nextDir.X <= 0) ||
					(currentTile == slopeS && nextDir.Y <= 0) ||
					(currentTile == slopeW && nextDir.X >= 0) {
					// no edge in this direction as there is a slope tile in a
					// different direction
					break
				}
			}

			e = e.Add(nextDir.WithZ(0))
			length++

			if v2 := vertexGrid[e.Y][e.X]; v2 >= 0 {
				// create edge as found another vertex
				vertices[v1] = append(vertices[v1], edge{
					vertex: v2,
					length: length,
				})
				break
			}

			// find the new next direction
			for i, dir := range vec.I2Directions {
				if i == vec.I2Opposites[e.Z] {
					// don't loop back
					continue
				}

				n := e.XY().Add(dir)
				if n.Y < 0 || n.Y >= len(d.grid) || n.X < 0 || n.X >= len(d.grid[0]) {
					continue
				}

				if d.grid[n.Y][n.X] != forest {
					e.Z = i
					break
				}
			}
		}
	}

	return vertexGrid[d.start.Y][d.start.X], vertexGrid[d.end.Y][d.end.X], vertices
}

func longestPath(start, end int, vertices []vertex) int {
	reverseVertices := make([]vertex, len(vertices))
	for i := 0; i < len(vertices); i++ {
		for _, e := range vertices[i] {
			reverseVertices[e.vertex] = append(reverseVertices[e.vertex], edge{vertex: i, length: e.length})
		}
	}

	endLength := 0
	for len(reverseVertices[end]) == 1 {
		e := reverseVertices[end][0]
		endLength += e.length
		end = e.vertex
	}

	maxLength := 0
	iterations := 0
	var f func(int, int, uint64)
	f = func(current, length int, visited uint64) {
		iterations++
		if current == end {
			maxLength = max(maxLength, length)
			return
		}

		for _, e := range vertices[current] {
			if visited&(1<<e.vertex) == 0 {
				f(e.vertex, length+e.length, visited|1<<e.vertex)
			}
		}
	}

	f(start, 0, 1<<start)
	return maxLength + endLength
}
