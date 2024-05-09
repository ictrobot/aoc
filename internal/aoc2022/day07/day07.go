package day07

import (
	_ "embed"
	"github.com/ictrobot/aoc-go/internal/util/parse"
	"math"
	"path"
	"strings"
)

//go:embed example
var Example string

type Day07 struct {
	folderSizes map[string]int
}

func (d *Day07) Parse(input string) {
	d.folderSizes = make(map[string]int)

	var currentPath []string
	for _, l := range parse.Lines(input) {
		if strings.HasPrefix(l, "$ cd") {
			dir := l[5:]
			if dir == "/" {
				currentPath = nil
			} else if dir == ".." {
				currentPath = currentPath[:len(currentPath)-1]
			} else {
				currentPath = append(currentPath, dir)
			}
		} else if !strings.HasPrefix(l, "$ ") && !strings.HasPrefix(l, "dir ") {
			size := parse.Int(parse.Whitespace(l)[0])
			for i := 0; i <= len(currentPath); i++ {
				d.folderSizes["/"+path.Join(currentPath[:i]...)] += size
			}
		}
	}
}

func (d *Day07) ParseExample() {
	d.Parse(Example)
}

func (d *Day07) Part1() any {
	size := 0
	for _, s := range d.folderSizes {
		if s < 100000 {
			size += s
		}
	}
	return size
}

func (d *Day07) Part2() any {
	toDelete := d.folderSizes["/"] - 40000000

	smallest := math.MaxInt
	for _, s := range d.folderSizes {
		if s > toDelete && s < smallest {
			smallest = s
		}
	}
	return smallest
}
