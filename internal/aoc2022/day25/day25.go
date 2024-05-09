package day25

import (
	_ "embed"
	"fmt"
	"github.com/ictrobot/aoc-go/internal/util/numbers"
	"github.com/ictrobot/aoc-go/internal/util/parse"
	"strings"
)

//go:embed example
var Example string

type Day25 struct {
	numbers []string
}

func (d *Day25) Parse(input string) {
	d.numbers = parse.Lines(input)
}

func (d *Day25) ParseExample() {
	d.Parse(Example)
}

func (d *Day25) Part1() any {
	sum := int64(0)
	for _, num := range d.numbers {
		sum += snafuDecode(num)
	}
	return snafuEncode(sum)
}

func (d *Day25) Part2() any {
	// no part 2
	return "🎄"
}

func snafuDecode(s string) int64 {
	sum := int64(0)
	pow := int64(1)
	for i := 0; i < len(s); i++ {
		var v int64
		switch s[len(s)-1-i] {
		case '2':
			v = 2
		case '1':
			v = 1
		case '0':
			v = 0
		case '-':
			v = -1
		case '=':
			v = -2
		default:
			panic(fmt.Errorf("unknown character: %c", s[len(s)-1-i]))
		}

		sum += pow * v
		pow *= 5
	}
	return sum
}

func snafuEncode(n int64) string {
	nAbs := numbers.IntAbs(n)
	divisor := int64(1)
	for divisor < nAbs {
		divisor *= 5
	}

	var b strings.Builder
	for divisor >= 1 {
		digit := numbers.IntRoundedDiv(n, divisor)

		switch digit {
		case 2:
			b.WriteByte('2')
		case 1:
			b.WriteByte('1')
		case 0:
			if b.Len() > 0 { // don't add leading zeroes
				b.WriteByte('0')
			}
		case -1:
			b.WriteByte('-')
		case -2:
			b.WriteByte('=')
		default:
			panic(fmt.Errorf("invalid digit for n=%d, digit=%d", n, digit))
		}

		n -= divisor * digit
		divisor /= 5
	}

	return b.String()
}
