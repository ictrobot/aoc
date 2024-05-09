package day07

import (
	_ "embed"
	"github.com/ictrobot/aoc-go/internal/util/parse"
	"slices"
	"strings"
)

//go:embed example
var Example string

type Day07 struct {
	hands []hand
}

type hand struct {
	cards string
	bid   int
}

const cardTypes = 14

var strength = [255]int8{
	'A': 1, 'K': 2, 'Q': 3, 'J': 4, 'T': 5,
	'9': 6, '8': 7, '7': 8, '6': 9, '5': 10,
	'4': 11, '3': 12, '2': 13, '1': 14,
}

var strengthWildcards = [255]int8{
	'A': 1, 'K': 2, 'Q': 3, 'T': 4, '9': 5,
	'8': 6, '7': 7, '6': 8, '5': 9, '4': 10,
	'3': 11, '2': 12, '1': 13, 'J': 14,
}

func (d *Day07) Parse(input string) {
	lines := parse.Lines(input)
	d.hands = make([]hand, len(lines))
	for i, line := range lines {
		idx := strings.IndexByte(line, ' ')
		d.hands[i].cards = line[:idx]
		d.hands[i].bid = parse.Int(line[idx+1:])
	}
}

func (d *Day07) ParseExample() {
	d.Parse(Example)
}

func (d *Day07) Part1() any {
	return d.winnings(false)
}

func (d *Day07) Part2() any {
	return d.winnings(true)
}

func (d *Day07) winnings(jAreWildcards bool) int {
	hands := slices.Clone(d.hands)

	slices.SortFunc(hands, func(a, b hand) int {
		// compare hand type, then each card strength
		aType := a.Type(jAreWildcards)
		bType := b.Type(jAreWildcards)
		if aType != bType {
			return bType - aType
		}

		for i := 0; i < len(a.cards); i++ {
			if a.cards[i] != b.cards[i] {
				if jAreWildcards {
					return int(strengthWildcards[b.cards[i]] - strengthWildcards[a.cards[i]])
				}
				return int(strength[b.cards[i]] - strength[a.cards[i]])
			}
		}

		return 0
	})

	sum := 0
	for r, h := range hands {
		sum += (r + 1) * h.bid
	}
	return sum
}

// Type returns the type of hand, lower is better
func (h *hand) Type(jAreWildcards bool) int {
	// find count of each card type
	jCount := int8(0)
	var typeCounts [cardTypes]int8
	for i := 0; i < len(h.cards); i++ {
		if jAreWildcards && h.cards[i] == 'J' {
			jCount++
			continue
		}

		// -1 will cause invalid cards to panic
		typeCounts[strength[h.cards[i]]-1]++
	}

	// find largest two counts
	var counts [2]int8
	for _, c := range typeCounts {
		if c > counts[0] {
			counts[1] = counts[0]
			counts[0] = c
		} else if c > counts[1] {
			counts[1] = c
		}
	}

	// find first type we have enough jokers to fulfil
	for t, req := range [][2]int8{{5}, {4}, {3, 2}, {3}, {2, 2}, {2}, {1}} {
		jUsed0 := req[0] - counts[0]
		if jUsed0 > jCount {
			continue
		}

		jUsed1 := req[1] - counts[1]
		if jUsed1 > jCount-jUsed0 {
			continue
		}

		return t
	}

	panic("no type")
}
