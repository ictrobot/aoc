package day17

import (
	_ "embed"
	"github.com/ictrobot/aoc/internal/util/vec"
	"strings"
)

//go:embed example
var Example string

const width = 7

// compare the last cacheKeyHeight rows when finding duplicates
const cacheKeyHeight = 100

type grid []uint8

type Day17 struct {
	jets string
}

type cacheValue struct {
	rockCount  int64
	gridHeight int64
	jetIdx     int
}

func (d *Day17) Parse(input string) {
	d.jets = strings.TrimSpace(input)
}

func (d *Day17) ParseExample() {
	d.Parse(Example)
}

func (d *Day17) Part1() any {
	return d.simulate(2022)
}

func (d *Day17) Part2() any {
	return d.simulate(1_000_000_000_000)
}

func (d *Day17) simulate(rocks int64) int64 {
	g := make(grid, 0)
	rock := make([]vec.I2[int], 5)
	cache := make(map[[cacheKeyHeight]uint8]cacheValue)
	jetIdx := 0
	height := int64(0)
	for rockCount := int64(0); rockCount < rocks; rockCount++ {
		// get next rock (using same slice to avoid allocations)
		rock = rock[:0]
		switch rockCount % 5 {
		case 0:
			rock = append(rock,
				vec.I2[int]{2, len(g) + 3},
				vec.I2[int]{3, len(g) + 3},
				vec.I2[int]{4, len(g) + 3},
				vec.I2[int]{5, len(g) + 3},
			)
		case 1:
			rock = append(rock,
				vec.I2[int]{3, len(g) + 3},
				vec.I2[int]{2, len(g) + 4},
				vec.I2[int]{3, len(g) + 4},
				vec.I2[int]{4, len(g) + 4},
				vec.I2[int]{3, len(g) + 5},
			)
		case 2:
			rock = append(rock,
				vec.I2[int]{2, len(g) + 3},
				vec.I2[int]{3, len(g) + 3},
				vec.I2[int]{4, len(g) + 3},
				vec.I2[int]{4, len(g) + 4},
				vec.I2[int]{4, len(g) + 5},
			)
		case 3:
			rock = append(rock,
				vec.I2[int]{2, len(g) + 3},
				vec.I2[int]{2, len(g) + 4},
				vec.I2[int]{2, len(g) + 5},
				vec.I2[int]{2, len(g) + 6},
			)
		case 4:
			rock = append(rock,
				vec.I2[int]{2, len(g) + 3},
				vec.I2[int]{3, len(g) + 3},
				vec.I2[int]{2, len(g) + 4},
				vec.I2[int]{3, len(g) + 4},
			)
		}

		for {
			if d.jets[jetIdx%len(d.jets)] == '<' { // left
				moveIfPossible(g, rock, -1, 0)
			} else { // right
				moveIfPossible(g, rock, 1, 0)
			}
			jetIdx++

			if !moveIfPossible(g, rock, 0, -1) {
				break
			}
		}

		for _, r := range rock {
			if r.Y >= len(g) {
				height++
				g = append(g, uint8(0))
			}
			g[r.Y] |= 1 << r.X
		}

		if len(g) >= cacheKeyHeight {
			key := ([cacheKeyHeight]uint8)(g[len(g)-cacheKeyHeight:])
			if s, ok := cache[key]; ok {
				rocksSince := rockCount - s.rockCount
				heightSince := height - s.gridHeight

				remainingRepeats := (rocks - rockCount) / rocksSince

				rockCount += rocksSince * remainingRepeats
				height += heightSince * remainingRepeats

				clear(cache)
			} else {
				cache[key] = cacheValue{rockCount, height, jetIdx}
			}
		}
	}

	return height
}

func moveIfPossible(g grid, rock []vec.I2[int], dX, dY int) bool {
	for _, r := range rock {
		if r.X+dX < 0 || r.X+dX >= width {
			return false // hits left or right wall
		}

		if r.Y+dY < 0 {
			return false // hits floor
		}

		if r.Y+dY < len(g) && g[r.Y+dY]&(1<<(r.X+dX)) != 0 {
			return false // hits existing rock
		}
	}

	for i := range rock {
		rock[i].X += dX
		rock[i].Y += dY
	}

	return true
}
