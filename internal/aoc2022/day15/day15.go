package day15

import (
	_ "embed"
	"github.com/ictrobot/aoc-go/internal/util/numbers"
	"github.com/ictrobot/aoc-go/internal/util/parse"
	"github.com/ictrobot/aoc-go/internal/util/vec"
	"slices"
)

//go:embed example
var Example string

type Day15 struct {
	isExample    bool // part 1 uses a different y value
	beacons      vec.I2Set[int]
	sensorRanges []vec.I3[int] // Z for radius
}

type interval struct {
	min, max int
}

func (d *Day15) Parse(input string) {
	d.isExample = false
	d.beacons = make(vec.I2Set[int]) // beacons must be deduplicated
	d.sensorRanges = nil

	ints := parse.ExtractInts(input)
	for i := 0; i < len(ints)-3; i += 4 {
		s := vec.I2[int]{ints[i], ints[i+1]}
		b := vec.I2[int]{ints[i+2], ints[i+3]}
		r := s.ManhattanDist(b)

		d.beacons.Add(b)
		d.sensorRanges = append(d.sensorRanges, s.WithZ(int(r)))
	}

	// presort sensor ranges by x for a slight speed boost when sorting in sensorIntervals
	slices.SortFunc(d.sensorRanges, func(a, b vec.I3[int]) int {
		return a.X - b.X
	})
}

func (d *Day15) ParseExample() {
	d.Parse(Example)
	d.isExample = true
}

func (d *Day15) Part1() any {
	y := 2000000
	if d.isExample {
		y = 10
	}

	// total area sensors cover, minus any known beacons
	intervals := make([]interval, 0)
	sensorIntervals(d.sensorRanges, y, &intervals)

	totalCovered := 0
	for _, i := range intervals {
		totalCovered += i.max - i.min + 1
	}

	for b := range d.beacons {
		if b.Y == y {
			totalCovered--
		}
	}

	return totalCovered
}

func (d *Day15) Part2() any {
	maxCoord := 4000000
	if d.isExample {
		maxCoord = 20
	}

	intervals := make([]interval, 0, len(d.sensorRanges))
	for y := 0; y < maxCoord; y += 1 {
		sensorIntervals(d.sensorRanges, y, &intervals)

		for _, i := range intervals {
			// since sensorIntervals combines overlapping & adjacent intervals
			// minX-1 or maxX+1 aren't covered by any sensors
			if i.min-1 >= 0 && i.min-1 <= maxCoord {
				return (int64(i.min-1) * int64(4000000)) + int64(y)
			}
			if i.max+1 >= 0 && i.max+1 <= maxCoord {
				return (int64(i.max+1) * int64(4000000)) + int64(y)
			}
		}
	}

	panic("not found")
}

// pass in pointer to existing slice to prevent allocating slice 4M times
func sensorIntervals(sensorRanges []vec.I3[int], y int, s *[]interval) {
	intervals := *s
	clear(intervals)

	// find list of x coordinate intervals sensors cover at y
	for _, r := range sensorRanges {
		yOffset := numbers.IntAbsDiff(r.Y, y)
		if yOffset > r.Z {
			// range does cover y
			continue
		}

		next := interval{
			r.X - (r.Z - yOffset),
			r.X + (r.Z - yOffset),
		}

		// if the new interval overlaps with the last one, combine it now instead
		// of waiting until after we sort the list, reducing the number to sort
		if len(intervals) > 0 && combineIntervals(&intervals[len(intervals)-1], &next) {
			continue
		}

		intervals = append(intervals, next)
	}

	if len(intervals) <= 1 {
		// nothing to sort/combine
		*s = intervals
		return
	}

	// sort & combine intervals
	slices.SortFunc(intervals, func(a, b interval) int {
		return a.min - b.min
	})

	i := 0
	for j := 1; j < len(intervals); j++ {
		if !combineIntervals(&intervals[i], &intervals[j]) {
			i++
			intervals[i] = intervals[j]
		}
	}

	*s = intervals[:i+1]
}

func combineIntervals(a, b *interval) bool {
	// +1 allows adjacent ranges, see TestCombineIntervals
	if a.min <= b.max+1 && b.min <= a.max+1 {
		// overlap or adjacent, can combine
		a.min = min(a.min, b.min)
		a.max = max(a.max, b.max)
		return true
	}
	return false
}
