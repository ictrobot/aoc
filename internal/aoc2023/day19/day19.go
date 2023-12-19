package day19

import (
	"bytes"
	_ "embed"
	"github.com/ictrobot/aoc/internal/util/collections"
	"github.com/ictrobot/aoc/internal/util/parse"
	"strings"
)

//go:embed example
var Example string

type Day19 struct {
	workflows map[string]workflow
	parts     []part
}

type workflow struct {
	rules    []rule
	fallback string
}

type rule struct {
	field    int
	lessThan bool // if true < else >
	value    int
	target   string
}

type part [types]int

type matching [types][maxValue]uint8

const (
	fieldX = iota
	fieldM
	fieldA
	fieldS
	types

	accepted = "A"
	rejected = "R"

	maxValue = 4000
)

func (d *Day19) Parse(input string) {
	chunks := parse.Chunks(input)

	workflows := parse.Grid(parse.Lines, parse.Whitespace, strings.NewReplacer(
		"{", " ",
		"}", "",
		"<", " < ",
		">", " > ",
		":", " ",
		",", " ",
	).Replace(chunks[0]))
	d.workflows = make(map[string]workflow)
	for _, l := range workflows {
		var rules []rule
		for i := 1; i < len(l)-1; i += 4 {
			var r rule
			switch l[i] {
			case "x":
				r.field = fieldX
			case "m":
				r.field = fieldM
			case "a":
				r.field = fieldA
			case "s":
				r.field = fieldS
			default:
				panic(l[i])
			}

			switch l[i+1] {
			case "<":
				r.lessThan = true
			case ">":
				// pass
			default:
				panic(l[i+1])
			}

			r.value = parse.Int(l[i+2])
			r.target = l[i+3]
			rules = append(rules, r)
		}

		d.workflows[l[0]] = workflow{rules, l[len(l)-1]}
	}

	partInts := parse.ExtractInts(chunks[1])
	d.parts = make([]part, 0, len(partInts)/types)
	for i := 0; i < len(partInts)-3; i += types {
		d.parts = append(d.parts, part(partInts[i:i+types]))
	}
}

func (d *Day19) ParseExample() {
	d.Parse(Example)
}

func (d *Day19) Part1() any {
	var result int
	for _, p := range d.parts {
		name := "in"

	workflow:
		for name != accepted && name != rejected {
			w := d.workflows[name]
			for _, r := range w.rules {
				if (r.lessThan && p[r.field] < r.value) || (!r.lessThan && p[r.field] > r.value) {
					name = r.target
					continue workflow
				}
			}
			name = w.fallback
		}

		if name == accepted {
			result += p[0] + p[1] + p[2] + p[3]
		}
	}
	return result
}

func (d *Day19) Part2() any {
	var a matching
	collections.Fill(a[0][:], 1)
	collections.Fill(a[1][:], 1)
	collections.Fill(a[2][:], 1)
	collections.Fill(a[3][:], 1)
	return d.combinations("in", a)
}

func (d *Day19) combinations(name string, m matching) (result int64) {
	if name == accepted {
		result = 1
		for i := 0; i < types; i++ {
			result *= int64(bytes.Count(m[i][:], []byte{1}))
		}
		return
	} else if name == rejected {
		return 0
	}

	w := d.workflows[name]
	nextRule := m
	for _, r := range w.rules {
		thisRule := nextRule

		if r.lessThan {
			// rule matches 1 <= x < r.Value, so unmark r.Value <= x <= max as matching
			collections.Fill(thisRule[r.field][r.value-1:], 0) // -1 for 0 indexed

			// next rule can't match 1 <= x < r.Value
			collections.Fill(nextRule[r.field][:r.value-1], 0) // -1 for 0 indexed
		} else {
			// rule matches r.Value < x <= max, so unmark 1 <= x <= r.Value as matching
			collections.Fill(thisRule[r.field][:r.value], 0) // -1 for 0 indexed, +1 as upper bound <=

			// next rule can't match r.Value < x <= max
			collections.Fill(nextRule[r.field][r.value:], 0) // -1 for 0 indexed, +1 as lower bound <
		}
		result += d.combinations(r.target, thisRule)
	}

	result += d.combinations(w.fallback, nextRule)
	return
}
