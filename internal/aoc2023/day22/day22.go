package day22

import (
	_ "embed"
	"github.com/ictrobot/aoc/internal/util/parse"
	"github.com/ictrobot/aoc/internal/util/vec"
	"slices"
)

//go:embed example
var Example string

type Day22 struct {
	bricks []brick
	maxZ   uint
}

type brick [2]vec.I3[uint]

func (d *Day22) Parse(input string) {
	ints := parse.ExtractUints(input)

	d.bricks = nil
	d.maxZ = 0
	for i := 0; i < len(ints)-5; i += 6 {
		d.bricks = append(d.bricks, brick{
			{ints[i+0], ints[i+1], ints[i+2]},
			{ints[i+3], ints[i+4], ints[i+5]},
		})
		d.maxZ = max(d.maxZ, ints[i+5])
	}

	// real input isn't sorted by z ascending unlike example!
	slices.SortFunc(d.bricks, func(a, b brick) int {
		if a[0].Z < b[0].Z {
			return -1
		}

		if a[1].Z < b[1].Z {
			return -1
		}

		return 0
	})
}

func (d *Day22) ParseExample() {
	d.Parse(Example)
}

func (d *Day22) Part1() any {
	bricks, supports, supportedBy := d.Settle()

	var count int
outer:
	for i := range bricks {
		for _, j := range supports[i] {
			if len(supportedBy[j]) == 1 {
				continue outer
			}
		}
		count++
	}
	return count
}

func (d *Day22) Part2() any {
	bricks, supports, supportedBy := d.Settle()

	var total int
	for i := range bricks {
		total += fallCount(bricks, supports, supportedBy, i)
	}
	return total
}

func (d *Day22) Settle() ([]brick, [][]int, [][]int) {
	bricks := slices.Clone(d.bricks)
	supportedBy := make([][]int, len(d.bricks))
	supports := make([][]int, len(d.bricks))
	atZ := make([][]int, d.maxZ)

	for i := range bricks {
		for bricks[i][0].Z > 1 && bricks[i][1].Z > 1 {
			newBrick := brick{
				bricks[i][0].Sub(vec.I3[uint]{0, 0, 1}),
				bricks[i][1].Sub(vec.I3[uint]{0, 0, 1}),
			}

			for _, j := range atZ[newBrick[0].Z-1] {
				if bricks[j].OverlapsWith(newBrick) {
					supportedBy[i] = append(supportedBy[i], j)
					supports[j] = append(supports[j], i)
				}
			}
			if len(supportedBy[i]) > 0 {
				break
			}

			bricks[i] = newBrick
		}

		atZ[bricks[i][1].Z-1] = append(atZ[bricks[i][1].Z-1], i)
	}

	return bricks, supports, supportedBy
}

func (b brick) OverlapsWith(o brick) bool {
	return b[0].Z <= o[1].Z && o[0].Z <= b[1].Z &&
		b[0].Y <= o[1].Y && o[0].Y <= b[1].Y &&
		b[0].X <= o[1].X && o[0].X <= b[1].X
}

func fallCount(bricks []brick, supports [][]int, supportedBy [][]int, d int) int {
	if len(supports[d]) == 0 {
		return 0
	}

	queue := append([]int{}, supports[d]...)
	disintegrated := make([]bool, len(bricks))
	disintegrated[d] = true
	fall := 0

outer:
	for len(queue) > 0 {
		b := queue[0]
		queue = queue[1:]

		if disintegrated[b] {
			continue
		}

		for _, j := range supportedBy[b] {
			if !disintegrated[j] {
				continue outer
			}
		}

		disintegrated[b] = true
		fall++

		queue = append(queue, supports[b]...)
	}

	return fall
}
