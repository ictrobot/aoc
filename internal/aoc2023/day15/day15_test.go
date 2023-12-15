package day15

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const Part1 = 1320
const Part2 = 145

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
