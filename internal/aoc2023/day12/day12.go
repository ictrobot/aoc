package day12

import (
	_ "embed"
	"github.com/ictrobot/aoc-go/internal/util/collections"
	"github.com/ictrobot/aoc-go/internal/util/parse"
	"slices"
	"strings"
)

//go:embed example
var Example string

type Day12 struct {
	rows []row
}

type row struct {
	springs []spring
	groups  []int
}

type spring uint8

const (
	operational spring = '.'
	broken      spring = '#'
	unknown     spring = '?'
)

func (d *Day12) Parse(input string) {
	lines := parse.Lines(input)
	d.rows = make([]row, len(lines))

	for i, l := range lines {
		split := strings.SplitN(l, " ", 2)

		d.rows[i].springs = make([]spring, len(split[0]))
		for j, c := range split[0] {
			d.rows[i].springs[j] = spring(c)
		}

		d.rows[i].groups = parse.ExtractInts(split[1])
	}
}

func (d *Day12) ParseExample() {
	d.Parse(Example)
}

func (d *Day12) Part1() any {
	var sum int64
	for _, r := range d.rows {
		sum += r.arrangements()
	}
	return sum
}

func (d *Day12) Part2() any {
	var sum int64
	for _, r := range d.rows {
		sum += row{ // remove last unknown from end to get 5 copies joined with unknown
			springs: collections.Repeat(append(r.springs, unknown), 5)[:5*(len(r.springs)+1)-1],
			groups:  collections.Repeat(r.groups, 5),
		}.arrangements()
	}
	return sum
}

func (r row) arrangements() int64 {
	m := make(map[uint32]int64)

	// currentGroup is the number of broken springs left in the current
	// contiguous group. 0 means a group just ended, so the next spring must be
	// operational. -1 is a placeholder for at least one operational spring has
	// passed since the last broken group (so a new group can start)
	var f func([]spring, []int, int, int) int64
	f = func(springs []spring, groups []int, currentGroup, remainingGroup int) int64 {
		if len(springs) < remainingGroup {
			// not enough remaining springs to match group pattern
			return 0
		}

		// process any known springs
		for len(springs) > 0 && (springs[0] == operational || springs[0] == broken) {
			s := springs[0]
			springs = springs[1:]

			if s == operational && currentGroup == -1 {
				// operational spring not immediately after broken group
			} else if s == operational && currentGroup == 0 {
				// operational spring immediately after broken group
				currentGroup = -1
				remainingGroup--
			} else if s == broken && currentGroup > 0 {
				// in contiguous broken group
				currentGroup--
				remainingGroup--
			} else if s == broken && currentGroup == -1 && len(groups) > 0 {
				// start next broken group
				groups, currentGroup = groups[1:], groups[0]-1
				remainingGroup--
			} else {
				// invalid combination
				return 0
			}
		}

		// if there are remaining springs, the first must be unknown
		if len(springs) > 0 {
			if currentGroup > 0 || (currentGroup == -1 && len(springs) == remainingGroup) {
				// unknown spring must be broken as inside contiguous broken group
				// or the number of remaining springs equals the minimum to match
				// the pattern
				springs[0] = broken
				result := f(springs, groups, currentGroup, remainingGroup)
				springs[0] = unknown
				return result
			}

			if currentGroup == 0 {
				// unknown spring must be operational as immediately after broken group
				springs[0] = operational
				result := f(springs, groups, currentGroup, remainingGroup)
				springs[0] = unknown
				return result
			}

			// unknown spring could be either - sum both possibilities and cache
			// use a packed integer for the map key for faster map performance
			key := (uint32(uint16(len(springs))) << 16) + uint32(uint16(remainingGroup))
			if result, ok := m[key]; ok {
				return result
			}

			springs[0] = broken
			result := f(springs, groups, currentGroup, remainingGroup)
			springs[0] = operational
			result += f(springs, groups, currentGroup, remainingGroup)
			springs[0] = unknown

			m[key] = result
			return result
		}

		// no more springs to process
		if remainingGroup > 0 {
			// invalid combination, expected more broken springs
			return 0
		}
		return 1
	}

	// calculate the minimum number of springs required to match the broken
	// groups pattern
	totalGroupSprings := 0
	for _, g := range r.groups {
		// all groups need to be followed by an operational spring so +1
		totalGroupSprings += g + 1
	}
	if totalGroupSprings > 0 {
		// last group doesn't need to be followed by operational spring
		totalGroupSprings--
	}

	return f(slices.Clone(r.springs), r.groups, -1, totalGroupSprings)
}
