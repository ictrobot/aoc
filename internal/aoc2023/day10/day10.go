package day10

import (
	_ "embed"
	"fmt"
	"github.com/ictrobot/aoc-go/internal/util/numbers"
	"github.com/ictrobot/aoc-go/internal/util/parse"
	"github.com/ictrobot/aoc-go/internal/util/vec"
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
	d.grid = vec.NewGrid[uint8](0, 0, len(lines[0])-1, len(lines)-1)

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
	pipe := d.start
	dir := oneDir(d.grid.Get(d.start))
	count := 0
	for {
		count++

		pipe = pipe.Add(offsets[dir])
		if pipe == d.start {
			return count / 2
		}

		dir = d.grid.Get(pipe) & (^opposites[dir])
	}
}

// Part2 calculates the enclosed tiles using the [shoelace formula] and
// [Pick's theorem]. Pick's theorem can be rearranged to give
//
//	interiorPoints = area - perimeter/2 + 1
//
// The perimeter is equal to the number of pipes in the loop, as each pipe is
// one unit apart. The area can be calculated using the shoelace formula
//
//	area = abs(sum(x_i*y_{i+1} - x_{i+1}*y_i)))/2
//
// Credit to all the users on the [reddit solutions thread] for the idea
//
// [shoelace formula]: https://en.wikipedia.org/wiki/Shoelace_formula#Triangle_formula
// [Pick's theorem]: https://en.wikipedia.org/wiki/Pick's_theorem
// [reddit solutions thread]: https://www.reddit.com/r/adventofcode/comments/18evyu9/2023_day_10_solutions
func (d *Day10) Part2() any {
	pipe := d.start
	dir := oneDir(d.grid.Get(pipe))
	twiceArea := 0
	perimeter := 0
	for {
		next := pipe.Add(offsets[dir])

		twiceArea += pipe.X*next.Y - next.X*pipe.Y
		perimeter++

		if next == d.start {
			return numbers.IntAbs(twiceArea)/2 - perimeter/2 + 1
		}

		pipe = next
		dir = d.grid.Get(next) & (^opposites[dir])
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
