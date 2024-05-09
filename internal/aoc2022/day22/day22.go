package day22

import (
	_ "embed"
	"fmt"
	"github.com/ictrobot/aoc-go/internal/util/numbers"
	"github.com/ictrobot/aoc-go/internal/util/parse"
	"github.com/ictrobot/aoc-go/internal/util/vec"
	"strings"
)

//go:embed example
var Example string

type Day22 struct {
	grid    [][]tile
	path    []interface{}
	example bool
}

type tile uint8

const (
	outOfBounds tile = iota
	open
	wall
)

const (
	right = iota
	down
	left
	up
)

var directions = [...]vec.I2[int]{
	{1, 0},
	{0, 1},
	{-1, 0},
	{0, -1},
}

func (d *Day22) Parse(input string) {
	chunks := parse.Chunks(input)

	gridLines := parse.Lines(chunks[0])
	maxLength := 0
	for _, l := range gridLines {
		maxLength = max(maxLength, len(l))
	}

	d.grid = make([][]tile, len(gridLines))
	for i, l := range gridLines {
		d.grid[i] = make([]tile, maxLength)
		for j, c := range l {
			if c == '.' {
				d.grid[i][j] = open
			} else if c == '#' {
				d.grid[i][j] = wall
			}
		}
	}

	path := strings.TrimSpace(chunks[1])
	d.path = nil
	i := 0
	for j := 1; j < len(path); j++ {
		iDigit := path[i] >= '0' && path[i] <= '9'
		jDigit := path[j] >= '0' && path[j] <= '9'

		if iDigit == jDigit {
			continue
		}

		if iDigit {
			d.path = append(d.path, parse.Int(path[i:j]))
		} else {
			d.path = append(d.path, path[i:j])
		}
		i = j
	}
	if path[i] >= '0' && path[i] <= '9' {
		d.path = append(d.path, parse.Int(path[i:]))
	} else {
		d.path = append(d.path, path[i:])
	}

	d.example = false
}

func (d *Day22) ParseExample() {
	d.Parse(Example)
	d.example = true
}

func (d *Day22) Part1() any {
	return d.password(func(pos vec.I3[int]) vec.I3[int] {
		for {
			pos.X = numbers.IntMod(pos.X+directions[pos.Z].X, len(d.grid[0]))
			pos.Y = numbers.IntMod(pos.Y+directions[pos.Z].Y, len(d.grid))

			if d.grid[pos.Y][pos.X] != outOfBounds {
				return pos
			}
		}
	})
}

func (d *Day22) Part2() any {
	if d.example {
		return "part 2 unsupported for example input due to different net"
	}

	// face layout:
	// .01
	// .2.
	// 43.
	// 5..

	faceSize := 50
	return d.password(func(pos vec.I3[int]) vec.I3[int] {
		n := pos
		n.X += directions[pos.Z].X
		n.Y += directions[pos.Z].Y
		if n.Y >= 0 && n.Y < len(d.grid) && n.X >= 0 && n.X < len(d.grid[0]) && d.grid[n.Y][n.X] != outOfBounds {
			return n
		}

		var face int
		faceX := pos.X / faceSize
		faceY := pos.Y / faceSize
		pos.X %= faceSize
		pos.Y %= faceSize
		if faceX == 1 && faceY == 0 {
			face = 0

			switch pos.Z {
			case left:
				return vec.I3[int]{0, 3*faceSize - 1 - pos.Y, right}
			case up:
				return vec.I3[int]{0, 3*faceSize + pos.X, right}
			}
		} else if faceX == 2 && faceY == 0 {
			face = 1

			switch pos.Z {
			case right:
				return vec.I3[int]{2*faceSize - 1, 3*faceSize - 1 - pos.Y, left}
			case down:
				return vec.I3[int]{2*faceSize - 1, faceSize + pos.X, left}
			case up:
				return vec.I3[int]{pos.X, 4*faceSize - 1, up}
			}
		} else if faceX == 1 && faceY == 1 {
			face = 2

			switch pos.Z {
			case right:
				return vec.I3[int]{2*faceSize + pos.Y, faceSize - 1, up}
			case left:
				return vec.I3[int]{pos.Y, 2 * faceSize, down}
			}
		} else if faceX == 1 && faceY == 2 {
			face = 3

			switch pos.Z {
			case right:
				return vec.I3[int]{3*faceSize - 1, faceSize - 1 - pos.Y, left}
			case down:
				return vec.I3[int]{faceSize - 1, 3*faceSize + pos.X, left}
			}
		} else if faceX == 0 && faceY == 2 {
			face = 4

			switch pos.Z {
			case left:
				return vec.I3[int]{faceSize, faceSize - 1 - pos.Y, right}
			case up:
				return vec.I3[int]{faceSize, faceSize + pos.X, right}
			}
		} else if faceX == 0 && faceY == 3 {
			face = 5

			switch pos.Z {
			case right:
				return vec.I3[int]{faceSize + pos.Y, 3*faceSize - 1, up}
			case down:
				return vec.I3[int]{2*faceSize + pos.X, 0, down}
			case left:
				return vec.I3[int]{faceSize + pos.Y, 0, down}
			}
		} else {
			panic(fmt.Errorf("pos on unknown face: %v", pos))
		}

		panic(fmt.Errorf("missing face mapping for face %d dir %d", face, pos.Z))
	})
}

func (d *Day22) password(forward func(vec.I3[int]) vec.I3[int]) int {
	startX := 0
	for d.grid[0][startX] != open {
		startX++
	}

	// use Z for direction
	pos := vec.I3[int]{startX, 0, right}
	for _, p := range d.path {
		if s, ok := p.(string); ok {
			if s == "L" {
				pos.Z = numbers.IntMod(pos.Z-1, 4)
			} else if s == "R" {
				pos.Z = numbers.IntMod(pos.Z+1, 4)
			}
		} else if c, ok := p.(int); ok {
			for i := 0; i < c; i++ {
				n := forward(pos)
				if n.Y < 0 || n.Y >= len(d.grid) || n.X < 0 || n.X >= len(d.grid[0]) {
					panic(fmt.Errorf("forward from %v gave %v, which is outside the grid", pos, n))
				}

				t := d.grid[n.Y][n.X]
				if t == wall {
					break
				} else if t == outOfBounds {
					panic(fmt.Errorf("forward from %v gave %v, which is out of bounds", pos, n))
				}

				pos = n
			}
		}
	}

	return ((pos.Y + 1) * 1000) + ((pos.X + 1) * 4) + pos.Z
}
