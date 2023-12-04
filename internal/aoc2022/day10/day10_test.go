package day10

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const Part1 = 13140
const Part2 = `##..##..##..##..##..##..##..##..##..##..
###...###...###...###...###...###...###.
####....####....####....####....####....
#####.....#####.....#####.....#####.....
######......######......######......####
#######.......#######.......#######.....`

func TestDay10_ParseExample(t *testing.T) {
	d1 := Day10{}
	d1.ParseExample()

	d2 := Day10{}
	d2.ParseExample()
	d2.ParseExample()

	assert.Equal(t, d1, d2, "should be idempotent")
}

func BenchmarkDay10_ParseExample(b *testing.B) {
	d := Day10{}
	for i := 0; i < b.N; i++ {
		d.ParseExample()
	}
}

func TestDay10_Part1(t *testing.T) {
	d := Day10{}
	d.ParseExample()

	assert.EqualValues(t, Part1, d.Part1())
}

func BenchmarkDay10_Part1(b *testing.B) {
	d := Day10{}
	d.ParseExample()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		assert.EqualValues(b, Part1, d.Part1())
	}
}

func TestDay10_Part2(t *testing.T) {
	d := Day10{}
	d.ParseExample()

	assert.EqualValues(t, Part2, d.Part2())
}

func BenchmarkDay10_Part2(b *testing.B) {
	d := Day10{}
	d.ParseExample()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		assert.EqualValues(b, Part2, d.Part2())
	}
}
