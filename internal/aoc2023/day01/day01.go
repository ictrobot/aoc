package day01

import (
	_ "embed"
	"github.com/ictrobot/aoc-go/internal/util/parse"
	"strconv"
	"strings"
)

//go:embed example
var Example string

//go:embed example2
var Example2 string

var spelled = []string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

type Day01 struct {
	nums1 [][]int
	nums2 [][]int
}

func (d *Day01) Parse(input string) {
	d.nums1 = parse.Grid(parse.Lines, parse.ExtractDigits, input)

	// spelled numbers can overlap!
	// e.g. "eightwo" = [8, 2]
	for i := 1; i <= 9; i++ {
		input = strings.ReplaceAll(input, spelled[i], spelled[i]+strconv.Itoa(i)+spelled[i])
	}

	d.nums2 = parse.Grid(parse.Lines, parse.ExtractDigits, input)
}

func (d *Day01) ParseExample() {
	d.Parse(Example)
}

func (d *Day01) ParseExample2() {
	d.Parse(Example2)
}

func (d *Day01) Part1() any {
	sum := 0
	for _, a := range d.nums1 {
		sum += (10 * a[0]) + a[len(a)-1]
	}
	return sum
}

func (d *Day01) Part2() any {
	sum := 0
	for _, a := range d.nums2 {
		sum += (10 * a[0]) + a[len(a)-1]
	}
	return sum
}
