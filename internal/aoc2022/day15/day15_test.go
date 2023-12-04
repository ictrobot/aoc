package day15

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const Part1 = 26
const Part2 = 56000011

func TestDay15_ParseExample(t *testing.T) {
	d1 := Day15{}
	d1.ParseExample()

	d2 := Day15{}
	d2.ParseExample()
	d2.ParseExample()

	assert.Equal(t, d1, d2, "should be idempotent")
}

func BenchmarkDay15_ParseExample(b *testing.B) {
	d := Day15{}
	for i := 0; i < b.N; i++ {
		d.ParseExample()
	}
}

func TestDay15_Part1(t *testing.T) {
	d := Day15{}
	d.ParseExample()

	assert.EqualValues(t, Part1, d.Part1())
}

func BenchmarkDay15_Part1(b *testing.B) {
	d := Day15{}
	d.ParseExample()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		assert.EqualValues(b, Part1, d.Part1())
	}
}

func TestDay15_Part2(t *testing.T) {
	d := Day15{}
	d.ParseExample()

	assert.EqualValues(t, Part2, d.Part2())
}

func BenchmarkDay15_Part2(b *testing.B) {
	d := Day15{}
	d.ParseExample()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		assert.EqualValues(b, Part2, d.Part2())
	}
}

func TestCombineIntervals(t *testing.T) {
	cases := []struct {
		msg        string
		in1, in2   interval
		combinable bool
		out1, out2 interval
	}{
		{
			"overlapping, first starts before second",
			interval{1, 5},
			interval{4, 7},
			true,
			interval{1, 7},
			interval{4, 7},
		},
		{
			"overlapping, first starts after second",
			interval{-10, -2},
			interval{-15, -9},
			true,
			interval{-15, -2},
			interval{-15, -9},
		},
		{
			"first inside second",
			interval{1, 3},
			interval{0, 5},
			true,
			interval{0, 5},
			interval{0, 5},
		},
		{
			"second inside first",
			interval{0, 100},
			interval{10, 20},
			true,
			interval{0, 100},
			interval{10, 20},
		},
		{
			"adjacent, first before second",
			interval{0, 5},
			interval{6, 10},
			true,
			interval{0, 10},
			interval{6, 10},
		},
		{
			"adjacent, first after second",
			interval{6, 10},
			interval{11, 15},
			true,
			interval{6, 15},
			interval{11, 15},
		},
		{
			"no overlap & not adjacent, first starts before second",
			interval{10, 15},
			interval{17, 20},
			false,
			interval{10, 15},
			interval{17, 20},
		},
		{
			"no overlap & not adjacent, first starts after second",
			interval{20, 25},
			interval{0, 18},
			false,
			interval{20, 25},
			interval{0, 18},
		},
	}

	for _, c := range cases {
		a := c.in1
		b := c.in2
		assert.Equal(t, c.combinable, combineIntervals(&a, &b), c.msg)
		assert.Equal(t, c.out1, a, c.msg)
		assert.Equal(t, c.out2, b, c.msg)
	}
}
