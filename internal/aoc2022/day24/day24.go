package day24

import (
	_ "embed"
	"github.com/ictrobot/aoc-go/internal/util/collections"
	"github.com/ictrobot/aoc-go/internal/util/numbers"
	"github.com/ictrobot/aoc-go/internal/util/parse"
	"github.com/ictrobot/aoc-go/internal/util/vec"
	"slices"
	"strings"
)

//go:embed example
var Example string

type Day24 struct {
	start, end    vec.I2[int]
	width, height int
	blizzards     []blizzard
}

type blizzard struct {
	pos, dir vec.I2[int]
}

func (d *Day24) Parse(input string) {
	lines := parse.Lines(input)

	d.start = vec.I2[int]{strings.IndexByte(lines[0], '.') - 1, -1}
	d.end = vec.I2[int]{strings.IndexByte(lines[len(lines)-1], '.') - 1, len(lines) - 2}

	d.width = len(lines[0]) - 2
	d.height = len(lines) - 2

	d.blizzards = make([]blizzard, 0, (len(lines)-2)*(len(lines[0])-2))
	for y := 1; y < len(lines)-1; y++ {
		for x := 1; x < len(lines[y])-1; x++ {
			var dir vec.I2[int]
			switch lines[y][x] {
			case '>':
				dir.X = 1
			case '<':
				dir.X = -1
			case 'v':
				dir.Y = 1
			case '^':
				dir.Y = -1
			default:
				continue
			}

			d.blizzards = append(d.blizzards, blizzard{vec.I2[int]{x - 1, y - 1}, dir})
		}
	}
}

func (d *Day24) ParseExample() {
	d.Parse(Example)
}

func (d *Day24) Part1() any {
	blizzards := slices.Clone(d.blizzards)
	return d.fastestRoute(d.start, d.end, vec.I2[int]{0, 1}, blizzards)
}

func (d *Day24) Part2() any {
	blizzards := slices.Clone(d.blizzards)
	mins1 := d.fastestRoute(d.start, d.end, vec.I2[int]{0, 1}, blizzards)
	mins2 := d.fastestRoute(d.end, d.start, vec.I2[int]{0, -1}, blizzards)
	mins3 := d.fastestRoute(d.start, d.end, vec.I2[int]{0, 1}, blizzards)
	return mins1 + mins2 + mins3
}

// 4 cardinal directions + staying still
var directions = [5]vec.I2[int]{{0, 1}, {0, -1}, {1, 0}, {-1, 0}, {0, 0}}

func (d *Day24) fastestRoute(start, end, startDir vec.I2[int], blizzards []blizzard) int {
	blocked := make([]bool, d.width*d.height)
	positions := make([]bool, d.width*d.height)
	newPositions := make([]bool, d.width*d.height)

	gridStart := start.Add(startDir)

	minute := 1
	for {
		for i := range blizzards {
			blizzards[i].pos.X = numbers.IntMod(blizzards[i].pos.X+blizzards[i].dir.X, d.width)
			blizzards[i].pos.Y = numbers.IntMod(blizzards[i].pos.Y+blizzards[i].dir.Y, d.height)
			blocked[(blizzards[i].pos.X*d.height)+blizzards[i].pos.Y] = true
		}

		// moving from start
		if !blocked[(gridStart.X*d.height)+gridStart.Y] {
			newPositions[(gridStart.X*d.height)+gridStart.Y] = true
		}

		// moving from previous minute positions
		for i := range positions {
			if !positions[i] {
				continue
			}

			for j := range directions {
				x := (i / d.height) + directions[j].X
				y := (i % d.height) + directions[j].Y

				if x == end.X && y == end.Y {
					return minute
				}

				if x < 0 || x >= d.width || y < 0 || y >= d.height {
					continue
				}

				if !blocked[(x*d.height)+y] {
					newPositions[(x*d.height)+y] = true
				}
			}
		}

		positions, newPositions = newPositions, positions
		collections.Fill(newPositions, false)
		collections.Fill(blocked, false)
		minute++
	}
}
