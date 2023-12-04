package day06

import (
	_ "embed"
	"github.com/samber/lo"
	"strings"
)

//go:embed example
var Example string

type Day06 struct {
	message []string
}

func (d *Day06) Parse(input string) {
	d.message = strings.Split(strings.TrimSpace(input), "")
}

func (d *Day06) ParseExample() {
	d.Parse(Example)
}

func (d *Day06) Part1() any {
	return d.firstUniqueCharacterIndex(4)
}

func (d *Day06) Part2() any {
	return d.firstUniqueCharacterIndex(14)
}

func (d *Day06) firstUniqueCharacterIndex(n int) int {
	for i := n; i < len(d.message); i++ {
		if len(lo.Uniq(d.message[i-n:i])) == n {
			return i
		}
	}
	panic("not found")
}
