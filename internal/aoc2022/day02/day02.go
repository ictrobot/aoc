package day02

import (
	_ "embed"
	"github.com/ictrobot/aoc/internal/util/parse"
	"github.com/samber/lo"
)

//go:embed example
var Example string

const (
	rock     = 1
	paper    = 2
	scissors = 3
)

type Day02 struct {
	moves [][]string
}

func (d *Day02) Parse(input string) {
	d.moves = lo.Chunk(parse.Whitespace(input), 2)
	//*d = parse.MustReflect[Day02](parse.Whitespace(input))
}

func (d *Day02) ParseExample() {
	d.Parse(Example)
}

var opponentMap = map[string]int{"A": rock, "B": paper, "C": scissors}
var youMap = map[string]int{"X": rock, "Y": paper, "Z": scissors}

func (d *Day02) Part1() any {
	total := 0
	for _, move := range d.moves {
		opponent := opponentMap[move[0]]
		you := youMap[move[1]]

		outcome := 0
		if (you == rock && opponent == scissors) || (you == paper && opponent == rock) || (you == scissors && opponent == paper) {
			outcome = 6
		} else if you == opponent {
			outcome = 3
		}

		total += you + outcome
	}
	return total
}

var winMap = map[string]int{"A": paper, "B": scissors, "C": rock}
var loseMap = map[string]int{"A": scissors, "B": rock, "C": paper}

func (d *Day02) Part2() any {
	total := 0
	for _, move := range d.moves {
		opponent := opponentMap[move[0]]

		var you, outcome int
		if move[1] == "Z" { // win
			you = winMap[move[0]]
			outcome = 6
		} else if move[1] == "Y" { // draw
			you = opponent
			outcome = 3
		} else { // lose
			you = loseMap[move[0]]
			outcome = 0
		}

		total += you + outcome
	}
	return total
}
