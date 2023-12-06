package day17

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const Part1 = 3068
const Part2 int64 = 1514285714288

func TestDay17_ParseExample(t *testing.T) {
	d1 := Day17{}
	d1.ParseExample()

	d2 := Day17{}
	d2.ParseExample()
	d2.ParseExample()

	assert.Equal(t, d1, d2, "should be idempotent")
}

func BenchmarkDay17_ParseExample(b *testing.B) {
	d := Day17{}
	for i := 0; i < b.N; i++ {
		d.ParseExample()
	}
}

func TestDay17_Part1(t *testing.T) {
	d := Day17{}
	d.ParseExample()

	assert.EqualValues(t, Part1, d.Part1())
}

func BenchmarkDay17_Part1(b *testing.B) {
	d := Day17{}
	d.ParseExample()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		assert.EqualValues(b, Part1, d.Part1())
	}
}

func TestDay17_Part2(t *testing.T) {
	d := Day17{}
	d.ParseExample()

	assert.EqualValues(t, Part2, d.Part2())
}

func BenchmarkDay17_Part2(b *testing.B) {
	d := Day17{}
	d.ParseExample()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		assert.EqualValues(b, Part2, d.Part2())
	}
}
