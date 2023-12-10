package day10

import (
	_ "embed"
	"fmt"
	"github.com/ictrobot/aoc/internal/util/parse"
	"github.com/ictrobot/aoc/internal/util/vec"
)

//go:embed example
var Example string

//go:embed example2
var Example2 string

const (
	north = 1 << iota
	east
	south
	west
)

var directions = [...]uint8{north, east, south, west}

var opposites = [...]uint8{
	north: south,
	east:  west,
	south: north,
	west:  east,
}

var offsets = [...]vec.I2[int]{
	north: {Y: -1},
	east:  {X: 1},
	south: {Y: 1},
	west:  {X: -1},
}

type Day10 struct {
	grid  *vec.Grid[uint8]
	start vec.I2[int]
}

func (d *Day10) Parse(input string) {
	lines := parse.Lines(input)
	d.grid = vec.NewGrid[uint8](0, 0, len(lines[0]), len(lines))

	d.start.X = -1
	for y, l := range lines {
		for x, c := range l {
			switch c {
			case '|':
				d.grid.SetInts(x, y, north|south)
			case '-':
				d.grid.SetInts(x, y, east|west)
			case 'L':
				d.grid.SetInts(x, y, north|east)
			case 'J':
				d.grid.SetInts(x, y, north|west)
			case '7':
				d.grid.SetInts(x, y, south|west)
			case 'F':
				d.grid.SetInts(x, y, south|east)
			case '.':
				// ground
			case 'S':
				d.start.X = x
				d.start.Y = y
			default:
				panic(fmt.Errorf("unknown character: %c", c))
			}
		}
	}
	if d.start.X == -1 {
		panic("no start")
	}

	// replace start with correct pipe type
	var startType uint8
	var startConnections int
	for _, dir := range directions {
		pipe := d.grid.Get(d.start.Add(offsets[dir]))
		if pipe&opposites[dir] != 0 {
			startType |= dir
			startConnections++
		}
	}
	if startConnections != 2 {
		panic(fmt.Errorf("found %d connections to start, expected 2", startConnections))
	}
	d.grid.Set(d.start, startType)
}

func (d *Day10) ParseExample() {
	d.Parse(Example)
}

func (d *Day10) ParseExample2() {
	d.Parse(Example2)
}

func (d *Day10) Part1() any {
	return d.loopPipesCount() / 2
}

func (d *Day10) Part2() any {
	loop := d.loopPipes()
	xMin, yMin, xMax, yMax := loop.Bounds()

	// for each row, calculate which tiles are inside by the parity of the
	// numbers of pipes to the left with north connections
	count := 0
	for y := yMin; y <= yMax; y++ {
		inside := false
		for x := xMin; x <= xMax; x++ {
			if loop.ContainsInts(x, y) {
				if d.grid.GetInts(x, y)&north != 0 {
					inside = !inside
				}
			} else if inside {
				count++
			}
		}
	}

	return count
}

// loopPipesCount returns the number of pipes in the loop from Day10.start
func (d *Day10) loopPipesCount() int {
	pos := d.start
	dir := oneDir(d.grid.Get(d.start))
	count := 0
	for {
		count++

		// get the next pipe in direction dir
		pos = pos.Add(offsets[dir])
		// find the next pipe's direction that isn't the way we've come from
		dir = d.grid.Get(pos) & (^opposites[dir])

		if pos == d.start {
			return count
		}
	}
}

// loopPipes returns a grid with the positions for all the pipes in the loop
// from Day10.start marked
func (d *Day10) loopPipes() *vec.Grid[bool] {
	pos := d.start
	dir := oneDir(d.grid.Get(pos))
	pipes := vec.NewGrid[bool](d.grid.Bounds())
	for {
		pipes.Set(pos, true)

		// get the next pipe in direction dir
		pos = pos.Add(offsets[dir])
		// find the next pipe's direction that isn't the way we've come from
		dir = d.grid.Get(pos) & (^opposites[dir])

		if pos == d.start {
			return pipes
		}
	}
}

// oneDir returns one of the directions a pipe is connected
func oneDir(pipe uint8) uint8 {
	for _, dir := range directions {
		if dir&pipe != 0 {
			return dir
		}
	}
	panic(fmt.Errorf("pipe has no connections: %b", pipe))
}
