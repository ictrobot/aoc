package day05

import (
	_ "embed"
	"fmt"
	"github.com/ictrobot/aoc/internal/util/parse"
	"github.com/samber/lo"
	"github.com/samber/lo/parallel"
	"math"
	"strings"
)

//go:embed example
var Example string

type Day05 struct {
	seeds []int
	maps  []rangeMap
}

type rangeMap struct {
	from, to string
	elements []mapping
}

type mapping struct {
	destStart   int
	sourceStart int
	length      int
}

func (d *Day05) Parse(input string) {
	chunks := parse.Chunks(input)

	d.seeds = parse.ExtractInts(chunks[0])

	d.maps = nil
	for _, c := range chunks[1:] {
		s := strings.Split(strings.Split(c, " ")[0], "-to-")
		r := rangeMap{from: s[0], to: s[1]}

		ints := parse.ExtractInts(c)
		for i := 0; i < len(ints)-2; i += 3 {
			r.elements = append(r.elements, mapping{ints[i+0], ints[i+1], ints[i+2]})
		}

		d.maps = append(d.maps, r)
	}
}

func (d *Day05) ParseExample() {
	d.Parse(Example)
}

func (d *Day05) Part1() any {
	m := math.MaxInt
	for _, s := range d.seeds {
		m = min(m, d.MapFrom("seed", "location", s))
	}
	return m
}

func (d *Day05) Part2() any {
	return lo.Min(parallel.Map(lo.Chunk(d.seeds, 2), func(r []int, _ int) int {
		m := math.MaxInt
		for s := r[0]; s <= r[0]+r[1]; s++ {
			m = min(m, d.MapFrom("seed", "location", s))
		}
		return m
	}))
}

func (d *Day05) MapFrom(from, to string, v int) int {
	for _, r := range d.maps {
		if r.from == from {
			out := v
			for _, e := range r.elements {
				if v >= e.sourceStart && v < e.sourceStart+e.length {
					out = e.destStart + (v - e.sourceStart)
				}
			}

			if r.to == to {
				return out
			}
			return d.MapFrom(r.to, to, out)
		}
	}

	panic(fmt.Errorf("%v %v %v", from, to, v))
}
