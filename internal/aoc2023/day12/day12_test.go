package day12

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const Part1 = 21
const Part2 = 525152

func TestDay12_ParseExample(t *testing.T) {
	d1 := Day12{}
	d1.ParseExample()

	d2 := Day12{}
	d2.ParseExample()
	d2.ParseExample()

	assert.Equal(t, d1, d2, "should be idempotent")
}

func BenchmarkDay12_ParseExample(b *testing.B) {
	d := Day12{}
	for i := 0; i < b.N; i++ {
		d.ParseExample()
	}
}

func TestDay12_Part1(t *testing.T) {
	d := Day12{}
	d.ParseExample()

	assert.EqualValues(t, Part1, d.Part1())
}

func BenchmarkDay12_Part1(b *testing.B) {
	d := Day12{}
	d.ParseExample()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		assert.EqualValues(b, Part1, d.Part1())
	}
}

func TestDay12_Part2(t *testing.T) {
	d := Day12{}
	d.ParseExample()

	assert.EqualValues(t, Part2, d.Part2())
}

func BenchmarkDay12_Part2(b *testing.B) {
	d := Day12{}
	d.ParseExample()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		assert.EqualValues(b, Part2, d.Part2())
	}
}
