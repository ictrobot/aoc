package day16

import (
	_ "embed"
	"github.com/ictrobot/aoc/internal/util/collections"
	"github.com/ictrobot/aoc/internal/util/parse"
	"github.com/samber/lo"
	"math"
	"regexp"
	"slices"
)

//go:embed example
var Example string

var punctuationRegex = regexp.MustCompile(`[=;,]`)

type Day16 struct {
	valveNames       []string
	startIdx         int
	valveRates       []int16
	rateValves       []int
	tunnels          [][]int
	pressureReleased []int16
}

func (d *Day16) Parse(input string) {
	// remove punctuation the parse into grid of lines & fields
	lines := parse.Grid(parse.Lines, parse.Whitespace, punctuationRegex.ReplaceAllString(input, " "))

	// use indexes to refer to valves, so first iterate to get all names
	d.valveNames = make([]string, 0, len(lines))
	for _, l := range lines {
		d.valveNames = append(d.valveNames, l[1])
	}

	d.startIdx = slices.Index(d.valveNames, "AA")
	if d.startIdx < 0 {
		panic("valve AA not found")
	}

	d.valveRates = make([]int16, 0, len(lines))
	d.rateValves = nil
	d.tunnels = make([][]int, 0, len(lines))
	for i, l := range lines {
		rate := parse.Int16(l[5])
		d.valveRates = append(d.valveRates, rate)
		if rate > 0 {
			d.rateValves = append(d.rateValves, i)
		}

		tunnels := make([]int, 0, len(l)-10)
		for _, v := range l[10:] {
			tunnels = append(tunnels, slices.Index(d.valveNames, v))
		}
		d.tunnels = append(d.tunnels, tunnels)
	}

	if len(d.rateValves) > 32 {
		panic("too many valves with non-zero rates")
	}

	// pressureReleased[0b1101] = how much pressure is released when valves 1, 3 & 4 are open
	d.pressureReleased = make([]int16, 1<<len(d.rateValves))
	for open := range d.pressureReleased {
		for i := 0; i < len(d.rateValves); i++ {
			if open&(1<<i) != 0 {
				d.pressureReleased[open] += d.valveRates[d.rateValves[i]]
			}
		}
	}
}

func (d *Day16) ParseExample() {
	d.Parse(Example)
}

func (d *Day16) Part1() any {
	return lo.Max(d.simulate(30))
}

func (d *Day16) Part2() any {
	state := d.simulate(26)

	// released[open] = max pressure released with those valves open, regardless of pos
	released := make([]int16, len(d.pressureReleased))
	for pos := 0; pos < len(d.valveNames); pos++ {
		for open := 0; open < len(d.pressureReleased); open++ {
			released[open] = max(released[open], state[(pos*len(d.pressureReleased))+open])
		}
	}

	most := int16(0)
	for you, youRelease := range released {
		if youRelease <= 0 {
			continue
		}

		for elephant := you + 1; elephant < len(released); elephant++ {
			if you&elephant != 0 {
				// can't both open the same valves
				continue
			}

			elephantRelease := released[elephant]
			if elephantRelease <= 0 {
				continue
			}

			most = max(most, youRelease+elephantRelease)
		}
	}
	return most
}

func (d *Day16) simulate(minutes int) []int16 {
	last := make([]int16, len(d.valveNames)*len(d.pressureReleased))
	collections.Fill(last, math.MinInt16)
	curr := slices.Clone(last)

	// last[pos][open] = last[(pos * m) + open]
	m := len(d.pressureReleased)

	// start at startIdx with zero valves open
	curr[(d.startIdx*m)+0] = 0

	for minute := 1; minute <= minutes; minute++ {
		// swap state slices and reset now current slice
		curr, last = last, curr
		collections.Fill(curr, math.MinInt16)

		for pos := 0; pos < len(d.valveNames); pos++ {
			posIdx := pos * m // [pos]

			if d.valveRates[pos] > 0 {
				// stay at location and open valve
				valveBit := 1 << slices.Index(d.rateValves, pos)

				for prevOpen, b := range last[posIdx : posIdx+m] {
					if b < 0 {
						// previous state not reachable
						continue
					}

					if prevOpen&valveBit > 0 {
						// valve already open
						continue
					}

					open := prevOpen | valveBit
					curr[posIdx+open] = max(curr[posIdx+open], b+d.pressureReleased[prevOpen])
				}
			}

			for _, prevPos := range d.tunnels[pos] {
				// travelled from prevPos
				prevPosIdx := prevPos * m

				for open, b := range last[prevPosIdx : prevPosIdx+m] {
					if b < 0 {
						// previous state not reachable
						continue
					}

					curr[posIdx+open] = max(curr[posIdx+open], b+d.pressureReleased[open])
				}
			}
		}
	}

	return curr
}
