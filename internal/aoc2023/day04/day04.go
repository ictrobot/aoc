package day04

import (
	_ "embed"
	"github.com/ictrobot/aoc-go/internal/util/collections"
	"github.com/ictrobot/aoc-go/internal/util/parse"
	"github.com/samber/lo"
	"slices"
	"strings"
)

//go:embed example
var Example string

type Day04 struct {
	cards []card
}

type card struct {
	Num     int
	Winning []int
	Have    []int
}

func (d *Day04) Parse(input string) {
	d.cards = nil
	for _, l := range parse.Lines(input) {
		s := strings.SplitN(l, "|", 2)
		i1 := parse.ExtractInts(s[0])
		d.cards = append(d.cards, card{
			Num:     i1[0],
			Winning: i1[1:],
			Have:    parse.ExtractInts(s[1]),
		})
	}
}

func (d *Day04) ParseExample() {
	d.Parse(Example)
}

func (d *Day04) Part1() any {
	total := 0
	for _, c := range d.cards {
		value := 0
		for _, n := range c.Winning {
			if slices.Contains(c.Have, n) {
				if value == 0 {
					value = 1
				} else {
					value *= 2
				}
			}
		}
		total += value
	}

	return total
}

func (d *Day04) Part2() any {
	cards := make([]int, len(d.cards))
	collections.Fill(cards, 1)

	for idx, count := range cards {
		if count == 0 {
			continue
		}

		c := d.cards[idx]

		matching := 0
		for _, n := range c.Winning {
			if slices.Contains(c.Have, n) {
				matching++
			}
		}

		for i := 1; i <= matching; i++ {
			cards[idx+i] += count
		}
	}

	return lo.Sum(cards)
}
