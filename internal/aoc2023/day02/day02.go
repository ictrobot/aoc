package day02

import (
	_ "embed"
	"github.com/ictrobot/aoc/internal/util/parse"
	"regexp"
)

//go:embed example
var Example string

var extractRegex = regexp.MustCompile(`\d+|blue|red|green|;`)

type Day02 struct {
	games []game
}

type game struct {
	ID   int
	Sets []struct {
		Cubes []struct {
			Count  int
			Colour string
		}
		_ parse.Placeholder `match:";" flags:"optional"`
	}
}

func (d *Day02) Parse(input string) {
	d.games = nil
	for _, l := range parse.Lines(input) {
		d.games = append(d.games, parse.MustReflect[game](extractRegex.FindAllString(l, -1)))
	}
}

func (d *Day02) ParseExample() {
	d.Parse(Example)
}

func (d *Day02) Part1() any {
	sum := 0
	for _, g := range d.games {
		if g.Possible(12, 13, 14) {
			sum += g.ID
		}
	}
	return sum
}

func (d *Day02) Part2() any {
	sum := 0
	for _, g := range d.games {
		red, green, blue := g.MinRequired()
		sum += red * green * blue
	}
	return sum
}

func (g game) Possible(redLimit, greenLimit, blueLimit int) bool {
	for _, s := range g.Sets {
		for _, c := range s.Cubes {
			var l int
			switch c.Colour {
			case "red":
				l = redLimit
			case "green":
				l = greenLimit
			case "blue":
				l = blueLimit
			default:
				panic("unknown colour" + c.Colour)
			}

			if c.Count > l {
				return false
			}
		}
	}
	return true
}

func (g game) MinRequired() (red, green, blue int) {
	for _, s := range g.Sets {
		for _, c := range s.Cubes {
			switch c.Colour {
			case "red":
				red = max(red, c.Count)
			case "green":
				green = max(green, c.Count)
			case "blue":
				blue = max(blue, c.Count)
			default:
				panic("unknown colour" + c.Colour)
			}
		}
	}
	return
}
