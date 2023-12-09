package day04

import (
	_ "embed"
	"github.com/ictrobot/aoc/internal/util/parse"
	"slices"
)

//go:embed example
var Example string

type Day04 struct {
	numbers []int
	boards  [][]int
}

func (d *Day04) Parse(input string) {
	chunks := parse.Chunks(input)

	d.numbers = parse.ExtractInts(chunks[0])

	d.boards = make([][]int, len(chunks)-1)
	for i, chunk := range chunks[1:] {
		d.boards[i] = parse.ExtractInts(chunk)
	}
}

func (d *Day04) ParseExample() {
	d.Parse(Example)
}

func (d *Day04) Part1() any {
	marked := make([][25]bool, len(d.boards))
	for _, num := range d.numbers {
		for i, b := range d.boards {
			idx := slices.Index(b, num)
			if idx < 0 {
				continue
			}
			marked[i][idx] = true

			rowStart := (idx / 5) * 5
			colStart := idx % 5
			if (marked[i][rowStart] && marked[i][rowStart+1] && marked[i][rowStart+2] && marked[i][rowStart+3] && marked[i][rowStart+4]) ||
				(marked[i][colStart] && marked[i][colStart+5] && marked[i][colStart+10] && marked[i][colStart+15] && marked[i][colStart+20]) {
				// won, calculate final score
				unmarkedSum := 0
				for j, n := range b {
					if !marked[i][j] {
						unmarkedSum += n
					}
				}
				return unmarkedSum * num
			}
		}
	}

	panic("no winning boards")
}

func (d *Day04) Part2() any {
	marked := make([][25]bool, len(d.boards))
	won := make([]bool, len(d.boards))
	remaining := len(d.boards)
	for _, num := range d.numbers {
		for i, b := range d.boards {
			if won[i] {
				continue
			}

			idx := slices.Index(b, num)
			if idx < 0 {
				continue
			}
			marked[i][idx] = true

			rowStart := (idx / 5) * 5
			colStart := idx % 5
			if (marked[i][rowStart] && marked[i][rowStart+1] && marked[i][rowStart+2] && marked[i][rowStart+3] && marked[i][rowStart+4]) ||
				(marked[i][colStart] && marked[i][colStart+5] && marked[i][colStart+10] && marked[i][colStart+15] && marked[i][colStart+20]) {

				won[i] = true
				remaining--

				if remaining == 0 {
					// last board to win, calculate final score
					unmarkedSum := 0
					for j, n := range b {
						if !marked[i][j] {
							unmarkedSum += n
						}
					}
					return unmarkedSum * num
				}
			}
		}
	}

	panic("not all boards win")
}
