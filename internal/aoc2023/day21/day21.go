package day21

import (
	_ "embed"
	"github.com/ictrobot/aoc-go/internal/util/numbers"
	"github.com/ictrobot/aoc-go/internal/util/parse"
	"github.com/ictrobot/aoc-go/internal/util/structures"
	"github.com/ictrobot/aoc-go/internal/util/vec"
	"golang.org/x/exp/maps"
)

//go:embed example
var Example string

type Day21 struct {
	size    int
	grid    [][]bool
	start   vec.I2[int]
	example bool
}

func (d *Day21) Parse(input string) {
	lines := parse.Lines(input)

	d.size = len(lines)
	d.grid = make([][]bool, d.size)
	d.start = vec.I2[int]{}
	d.example = false

	for y, l := range lines {
		if len(l) != len(lines) {
			panic("grid must be square")
		}

		d.grid[y] = make([]bool, d.size)
		for x, c := range l {
			if c == 'S' {
				d.start.X = x
				d.start.Y = y
			}

			d.grid[y][x] = c == '#'
		}
	}
}

func (d *Day21) ParseExample() {
	d.Parse(Example)
	d.example = true
}

func (d *Day21) Part1() any {
	steps := 64
	if d.example {
		steps = 6
	}

	return d.reachablePlots(steps)
}

func (d *Day21) Part2() any {
	steps := 26501365
	if d.example {
		steps = 5000
	}

	return d.reachablePlots(steps)
}

func (d *Day21) reachablePlots(steps int) int64 {
	repeatAfter := d.size
	repeats := steps / repeatAfter
	extra := steps % repeatAfter

	if repeats <= 6 {
		firstReached := d.firstReached(steps)
		xMin, yMin, xMax, yMax := firstReached.Bounds()
		return reachable(firstReached, xMin, yMin, xMax, yMax, steps)
	}

	// looking at the example input with steps=5000, repeats=454 and extra=6.
	// after 4 repeats (6 + 4*11 = 50 steps) the number of reachable garden
	// plots within each repeating grid looks like:
	// -4, -4                4, -4                  4x  2
	//   0  0  0  6 17  9  0  0  0                  4x  6
	//   0  0  6 36 39 38  9  0  0                  8x  9
	//   0  6 36 39 42 39 38  9  0                  2x  15
	//   6 36 39 42 39 42 39 38  9                  1x  17
	//  18 39 42 39 42 39 42 39 15                  1x  18
	//   9 38 39 42 39 42 39 29  2                  3x  29
	//   0  9 38 39 42 39 29  2  0                  3x  36
	//   0  0  9 38 39 29  2  0  0                  6x  38
	//   0  0  0  9 15  2  0  0  0                  16x 39
	// -4, 4                  4, 4                  9x  42
	//
	// after 5 repeats (6 + 5*11 = 61 steps):
	// -5, -5                      5, -5            5x  2
	//   0  0  0  0  6 17  9  0  0  0  0            5x  6
	//   0  0  0  6 36 39 38  9  0  0  0            10x 9
	//   0  0  6 36 39 42 39 38  9  0  0            2x  15
	//   0  6 36 39 42 39 42 39 38  9  0            1x  17
	//   6 36 39 42 39 42 39 42 39 38  9            1x  18
	//  18 39 42 39 42 39 42 39 42 39 15            4x  29
	//   9 38 39 42 39 42 39 42 39 29  2            4x  36
	//   0  9 38 39 42 39 42 39 29  2  0            8x  38
	//   0  0  9 38 39 42 39 29  2  0  0            25x 39
	//   0  0  0  9 38 39 29  2  0  0  0            16x 42
	//   0  0  0  0  9 15  2  0  0  0  0
	// -5, 5                        5, 5
	//
	// after 6 repeats (6 + 6*11 = 72 steps):
	// -6, -6                            6, -6      6x  2
	//   0  0  0  0  0  6 17  9  0  0  0  0  0      6x  6
	//   0  0  0  0  6 36 39 38  9  0  0  0  0      12x 9
	//   0  0  0  6 36 39 42 39 38  9  0  0  0      2x  15
	//   0  0  6 36 39 42 39 42 39 38  9  0  0      1x  17
	//   0  6 36 39 42 39 42 39 42 39 38  9  0      1x  18
	//   6 36 39 42 39 42 39 42 39 42 39 38  9      5x  29
	//  18 39 42 39 42 39 42 39 42 39 42 39 15      5x  36
	//   9 38 39 42 39 42 39 42 39 42 39 29  2      10x 38
	//   0  9 38 39 42 39 42 39 42 39 29  2  0      36x 39
	//   0  0  9 38 39 42 39 42 39 29  2  0  0      25x 42
	//   0  0  0  9 38 39 42 39 29  2  0  0  0
	//   0  0  0  0  9 38 39 29  2  0  0  0  0
	//   0  0  0  0  0  9 15  2  0  0  0  0  0
	// -6, 6                              6, 6
	//
	// therefore extrapolate how many grids there are with each count from the
	// pattern at the 4th, 5th and 6th repeat.

	firstReached := d.firstReached(extra + 6*repeatAfter)

	minX, minY, maxX, maxY := firstReached.Bounds()
	minX = (minX - d.size + 1) / d.size * d.size
	minY = (minY - d.size + 1) / d.size * d.size

	repeatCounts := [...]map[int64]int64{make(map[int64]int64), make(map[int64]int64), make(map[int64]int64)}
	for x := minX; x <= maxX; x += d.size {
		for y := minY; y <= maxY; y += d.size {
			for i := 0; i < 3; i++ {
				c := reachable(firstReached, x, y, x+d.size-1, y+d.size-1, extra+(4+i)*repeatAfter)
				if c > 0 {
					repeatCounts[i][c]++
				}
			}
		}
	}

	unique := maps.Keys(repeatCounts[0])
	var total int64
	for _, c := range unique {
		total += c * extrapolate(repeatCounts[0][c], repeatCounts[1][c], repeatCounts[2][c], int64(repeats)-6)
	}
	return total
}

func (d *Day21) firstReached(maxSteps int) *vec.Grid[int] {
	// use Z to store the step the position is reached first
	var queue structures.Deque[vec.I3[int]]
	queue.PushBack(d.start.WithZ(0))

	reached := vec.NewGrid[int](d.start.X-maxSteps, d.start.Y-maxSteps, d.start.X+maxSteps, d.start.Y+maxSteps)

	for !queue.IsEmpty() {
		p := queue.PopFront()

		if !reached.SetIfZero(p.XY(), p.Z) {
			continue
		}

		if p.Z >= maxSteps {
			continue
		}

		for _, dir := range vec.I2Directions {
			n := p.Add(dir.WithZ(1))

			if d.grid[numbers.IntMod(n.Y, d.size)][numbers.IntMod(n.X, d.size)] {
				continue
			}

			if reached.Contains(n.XY()) {
				continue
			}

			queue.PushBack(n)
		}
	}

	return reached
}

func reachable(firstReached *vec.Grid[int], xMin, yMin, xMax, yMax, steps int) (plots int64) {
	for x := xMin; x <= xMax; x++ {
		for y := yMin; y <= yMax; y++ {
			v := firstReached.GetInts(x, y)
			if v > 0 && v <= steps && v%2 == steps%2 {
				plots++
			}
		}
	}
	return
}

func extrapolate(x0, x1, x2, n int64) int64 {
	diff0 := x1 - x0
	diff1 := x2 - x1

	if diff0 == diff1 {
		return x2 + n*diff0
	}

	return x2 + n*((n+3)*diff1-(n+1)*diff0)/2
}
