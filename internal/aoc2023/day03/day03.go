package day03

import (
	_ "embed"
	"github.com/ictrobot/aoc-go/internal/util/parse"
)

//go:embed example
var Example string

type Day03 struct {
	lines []string
}

func (d *Day03) Parse(input string) {
	d.lines = parse.Lines(input)
}

func (d *Day03) ParseExample() {
	d.Parse(Example)
}

func (d *Day03) Part1() any {
	sum := 0
	for i, line := range d.lines {
		for j := 0; j < len(line); j++ {
			if line[j] < '0' || line[j] > '9' {
				continue
			}

			j2 := j
			for j2+1 < len(line) && line[j2+1] >= '0' && line[j2+1] <= '9' {
				j2++
			}

			if neighbouringSymbol(d.lines, i, j, j2) {
				sum += parse.Int(line[j : j2+1])
			}

			j = j2 + 1
		}
	}

	return sum
}

func (d *Day03) Part2() any {
	sum := 0
	for i, line := range d.lines {
		for j := 0; j < len(line); j++ {
			if line[j] != '*' {
				continue
			}

			nums := getNums(d.lines, i, j)
			if len(nums) != 2 {
				continue
			}

			sum += nums[0] * nums[1]
		}
	}

	return sum
}

func neighbouringSymbol(l []string, Y, X1, X2 int) bool {
	// check for at least one symbol nearby (neither digit nor dot)
	for y := max(Y-1, 0); y < min(Y+2, len(l)); y++ {
		for x := max(X1-1, 0); x < min(X2+2, len(l)); x++ {
			if !(l[y][x] >= '0' && l[y][x] <= '9') && l[y][x] != '.' {
				return true
			}
		}
	}
	return false
}

func getNums(l []string, Y, X int) []int {
	// extract neighbouring numbers from position
	var nums []int
	for y := max(Y-1, 0); y < min(Y+2, len(l)); y++ {
		for x := max(X-1, 0); x < min(X+2, len(l)); x++ {
			if l[y][x] < '0' || l[y][x] > '9' {
				continue
			}

			x1, x2 := numIndex(l[y], x)
			nums = append(nums, parse.Int(l[y][x1:x2+1]))
			x = x2 + 1
		}
	}
	return nums
}

func numIndex(l string, X int) (int, int) {
	// find start and end index of number
	x1 := X
	for x1 > 0 && l[x1-1] >= '0' && l[x1-1] <= '9' {
		x1--
	}

	x2 := X
	for x2+1 < len(l) && l[x2+1] >= '0' && l[x2+1] <= '9' {
		x2++
	}

	return x1, x2
}
