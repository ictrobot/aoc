package day05

import (
	_ "embed"
	"fmt"
	"github.com/ictrobot/aoc-go/internal/util/parse"
	"math"
	"strings"
)

//go:embed example
var Example string

type Day05 struct {
	seeds []int64
	maps  map[string]rangeMap
}

type rangeMap struct {
	from, to string
	elements []mapping
}

type mapping struct {
	dstStart int64
	srcStart int64
	srcEnd   int64
}

func (d *Day05) Parse(input string) {
	chunks := parse.Chunks(input)

	d.seeds = parse.ExtractInt64s(chunks[0])

	d.maps = make(map[string]rangeMap)
	for _, c := range chunks[1:] {
		s := strings.Split(strings.Split(c, " ")[0], "-to-")
		r := rangeMap{from: s[0], to: s[1]}

		ints := parse.ExtractInt64s(c)
		for i := 0; i < len(ints)-2; i += 3 {
			r.elements = append(r.elements, mapping{
				dstStart: ints[i],
				srcStart: ints[i+1],
				srcEnd:   ints[i+1] + ints[i+2] - 1,
			})
		}

		d.maps[r.from] = r
	}
}

func (d *Day05) ParseExample() {
	d.Parse(Example)
}

func (d *Day05) Part1() any {
	m := int64(math.MaxInt64)
	for _, s := range d.seeds {
		m = min(m, d.Map("seed", "location", s))
	}
	return m
}

func (d *Day05) Part2() any {
	m := int64(math.MaxInt64)
	for i := 0; i < len(d.seeds)-1; i += 2 {
		results := d.MapInterval("seed", "location", [2]int64{d.seeds[0], d.seeds[0] + d.seeds[1] - 1})
		for _, r := range results {
			m = min(m, r[0], r[1])
		}
	}
	return m
}

func (d *Day05) Map(from, to string, v int64) int64 {
	for {
		m, mOk := d.maps[from]
		if !mOk {
			panic(fmt.Errorf("no map from %s", from))
		}

		for _, e := range m.elements {
			if v <= e.srcEnd && e.srcStart <= v {
				v = v - e.srcStart + e.dstStart
				break
			}
		}

		if m.to == to {
			return v
		}

		from = m.to
	}
}

func (d *Day05) MapInterval(from, to string, in [2]int64) [][2]int64 {
	var queue, result [][2]int64
	queue = append(queue, in)

	for {
		m, mOk := d.maps[from]
		if !mOk {
			panic(fmt.Errorf("no map from %s", from))
		}

	intervals:
		for len(queue) > 0 {
			i := queue[0]
			queue = queue[1:]

			for _, e := range m.elements {
				if i[0] <= e.srcEnd && e.srcStart <= i[1] {
					result = append(result, [2]int64{
						max(i[0], e.srcStart) - e.srcStart + e.dstStart,
						min(i[1], e.srcEnd) - e.srcStart + e.dstStart,
					})

					if i[0] < e.srcStart {
						queue = append(queue, [2]int64{i[0], e.srcStart - 1})
					}

					if e.srcEnd < i[1] {
						queue = append(queue, [2]int64{e.srcEnd + 1, i[1]})
					}

					continue intervals
				}
			}

			result = append(result, i)
		}

		if m.to == to {
			return result
		}

		from = m.to
		queue = result
		result = nil
	}
}
