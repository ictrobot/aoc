package day20

import (
	_ "embed"
	"github.com/ictrobot/aoc-go/internal/util/numbers"
	"github.com/ictrobot/aoc-go/internal/util/parse"
	"slices"
)

//go:embed example
var Example string

const decryptionKey = 811589153

type Day20 struct {
	nums []int64
}

type mixingValue struct {
	num         int64
	originalIdx int
}

func (d *Day20) Parse(input string) {
	d.nums = parse.ExtractInt64s(input)
}

func (d *Day20) ParseExample() {
	d.Parse(Example)
}

func (d *Day20) Part1() any {
	return coordsSum(d.nums, 1)
}

func (d *Day20) Part2() any {
	decrypted := make([]int64, len(d.nums))
	for i, num := range d.nums {
		decrypted[i] = num * decryptionKey
	}

	return coordsSum(decrypted, 10)
}

func coordsSum(nums []int64, times int) int64 {
	values := make([]mixingValue, len(nums))
	for i, num := range nums {
		values[i] = mixingValue{num, i}
	}

	for repeat := 0; repeat < times; repeat++ {
		for originalIdx := 0; originalIdx < len(values); originalIdx++ {
			currentIdx := 0
			for i, v := range values {
				if v.originalIdx == originalIdx {
					currentIdx = i
					break
				}
			}

			v := values[currentIdx]
			values = slices.Delete(values, currentIdx, currentIdx+1)
			newIdx := numbers.IntMod(int64(currentIdx)+v.num, int64(len(values)))
			values = slices.Insert(values, int(newIdx), v)
		}
	}

	zeroIdx := 0
	for i, v := range values {
		if v.num == 0 {
			zeroIdx = i
			break
		}
	}

	return values[(zeroIdx+1000)%len(values)].num +
		values[(zeroIdx+2000)%len(values)].num +
		values[(zeroIdx+3000)%len(values)].num
}
