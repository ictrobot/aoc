package day03

import (
	_ "embed"
	"fmt"
	"github.com/ictrobot/aoc/internal/util/parse"
	"slices"
	"strconv"
)

//go:embed example
var Example string

type Day03 struct {
	nums []string
}

func (d *Day03) Parse(input string) {
	d.nums = parse.Lines(input)
}

func (d *Day03) ParseExample() {
	d.Parse(Example)
}

func (d *Day03) Part1() any {
	onesCount := make([]int, len(d.nums[0]))
	for _, n := range d.nums {
		for i, c := range n {
			if c == '1' {
				onesCount[i]++
			}
		}
	}

	var epsilon, gamma int
	for i, c := range onesCount {
		if c > len(d.nums)-c {
			// epsilon rate is the most common bit in each position, so if more
			// than half the numbers had a 1 in position i, add a 1 to epsilon
			epsilon += 1 << (len(d.nums[0]) - 1 - i)
		} else {
			// gamma rate is the least common bit in each position, so if less
			// than half the numbers had a 1 in position i, add a 1 to gamma
			gamma += 1 << (len(d.nums[0]) - 1 - i)
		}
	}
	return epsilon * gamma
}

func (d *Day03) Part2() any {
	oxygen, err := strconv.ParseInt(filter(true, d.nums), 2, strconv.IntSize)
	if err != nil {
		panic(fmt.Errorf("failed to parse binary: %w", err))
	}
	co2, err := strconv.ParseInt(filter(false, d.nums), 2, strconv.IntSize)
	if err != nil {
		panic(fmt.Errorf("failed to parse binary: %w", err))
	}
	return oxygen * co2
}

func filter(mostCommon bool, nums []string) string {
	before := slices.Clone(nums)
	after := make([]string, 0, len(nums))

	for i := 0; i < len(nums[0]); i++ {
		ones := 0
		for _, n := range before {
			if n[i] == '1' {
				ones++
			}
		}

		var keepOnes bool
		if mostCommon {
			keepOnes = ones >= len(before)-ones
		} else {
			keepOnes = ones < len(before)-ones
		}

		for _, n := range before {
			if keepOnes == (n[i] == '1') {
				after = append(after, n)
			}
		}

		if len(after) == 1 {
			return after[0]
		}

		before, after = after, before[:0]
	}

	panic("did not filter to single result")
}
