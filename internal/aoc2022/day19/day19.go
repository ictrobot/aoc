package day19

import (
	_ "embed"
	"github.com/ictrobot/aoc/internal/util/parse"
	"github.com/samber/lo"
	"github.com/samber/lo/parallel"
	"math"
)

//go:embed example
var Example string

const (
	ore = iota
	clay
	obsidian
	geode
	numTypes
)

type Day19 struct {
	blueprints []Blueprint
}

type Blueprint struct {
	num   int
	costs [numTypes][numTypes]int
}

func (d *Day19) Parse(input string) {
	chunks := lo.Chunk(parse.ExtractInts(input), 7)
	d.blueprints = make([]Blueprint, len(chunks))
	for i, ints := range chunks {
		d.blueprints[i] = Blueprint{
			num: ints[0],
			costs: [numTypes][numTypes]int{
				{ints[1]},             // ore
				{ints[2]},             // clay
				{ints[3], ints[4]},    // obsidian
				{ints[5], 0, ints[6]}, // geode
			},
		}
	}
}

func (d *Day19) ParseExample() {
	d.Parse(Example)
}

func (d *Day19) Part1() any {
	return lo.Sum(parallel.Map(d.blueprints, func(b Blueprint, i int) int {
		return b.num * b.maxGeodes(24)
	}))
}

func (d *Day19) Part2() any {
	return lo.Reduce(
		parallel.Map(lo.Subset(d.blueprints, 0, 3), func(b Blueprint, i int) int {
			return b.maxGeodes(32)
		}),
		func(total int, v int, i int) int {
			return total * v
		},
		1,
	)
}

func (b *Blueprint) maxGeodes(runtime int) int {
	// it's not worth having more e.g. ore robots than ore can be used a minute
	maxRobots := []int{
		max(b.costs[ore][ore], b.costs[clay][ore], b.costs[obsidian][ore], b.costs[geode][ore]),
		b.costs[obsidian][clay],
		b.costs[geode][obsidian],
		math.MaxInt,
	}

	maxFound := 0

	var search func(int, [numTypes]int, [numTypes]int)
	search = func(minute int, robots [numTypes]int, resources [numTypes]int) {
		remaining := runtime - minute
		geodes := resources[geode] + (remaining * robots[geode])
		maxFound = max(maxFound, geodes)

		// no point building a robot in the last minute
		if remaining <= 1 {
			return
		}

		// calculate the maximum number of geodes if we could build a geode
		// robot every minute from now - if it is less than or equal to the max
		// already found, stop
		if (geodes+(geodes+remaining))*(remaining+1)/2 <= maxFound {
			return
		}

		// try craft geode robot first
	robotTypes:
		for _, robotType := range []int{geode, ore, clay, obsidian} {
			// check if we already have enough robots for the type
			if robots[robotType] >= maxRobots[robotType] {
				continue
			}

			// calculate how many minutes we would need to wait to build the
			// given robot type without building any other robots in between
			mins := 0
			for i := 0; i < numTypes; i++ {
				needed := b.costs[robotType][i] - resources[i]
				if needed <= 0 {
					continue
				}
				if robots[i] <= 0 {
					// no robots building required resource, cannot just wait
					continue robotTypes
				}
				// (n+(d-1))/d = n/d rounded up
				mins = max(mins, (needed+robots[i]-1)/robots[i])
			}
			// takes a minute to build the robot
			mins++
			// check if we can build the robot before the end
			if minute+mins > runtime {
				continue
			}

			// continue the search after the new robot is built
			newRobots := robots
			newRobots[robotType]++

			newResources := resources
			for i := 0; i < numTypes; i++ {
				newResources[i] += (robots[i] * mins) - b.costs[robotType][i]
			}

			search(minute+mins, newRobots, newResources)

			if mins == 1 && robotType == geode {
				break
			}
		}
	}

	search(0, [numTypes]int{1}, [numTypes]int{})
	return maxFound
}
