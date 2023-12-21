package day21

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const Part1 = 16
const Part2 = 16733044

func TestDay21_ParseExample(t *testing.T) {
	d1 := Day21{}
	d1.ParseExample()

	d2 := Day21{}
	d2.ParseExample()
	d2.ParseExample()

	assert.Equal(t, d1, d2, "should be idempotent")
}

func BenchmarkDay21_ParseExample(b *testing.B) {
	d := Day21{}
	for i := 0; i < b.N; i++ {
		d.ParseExample()
	}
}

func TestDay21_Part1(t *testing.T) {
	d := Day21{}
	d.ParseExample()

	assert.EqualValues(t, Part1, d.Part1())
}

func BenchmarkDay21_Part1(b *testing.B) {
	d := Day21{}
	d.ParseExample()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		assert.EqualValues(b, Part1, d.Part1())
	}
}

func TestDay21_Part2(t *testing.T) {
	d := Day21{}
	d.ParseExample()

	assert.EqualValues(t, Part2, d.Part2())
}

func BenchmarkDay21_Part2(b *testing.B) {
	d := Day21{}
	d.ParseExample()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		assert.EqualValues(b, Part2, d.Part2())
	}
}

func TestDay21_reachablePlots(t *testing.T) {
	d := Day21{}
	d.ParseExample()
	assert.EqualValues(t, 16, d.reachablePlots(6))
	assert.EqualValues(t, 50, d.reachablePlots(10))
	assert.EqualValues(t, 1594, d.reachablePlots(50))
	assert.EqualValues(t, 6536, d.reachablePlots(100))
	assert.EqualValues(t, 167004, d.reachablePlots(500))
	assert.EqualValues(t, 668697, d.reachablePlots(1000))
	assert.EqualValues(t, 16733044, d.reachablePlots(5000))
}

func Test_extrapolate(t *testing.T) {
	assert.EqualValues(t, 1, extrapolate(1, 1, 1, 1))
	assert.EqualValues(t, 1, extrapolate(1, 1, 1, 2))
	assert.EqualValues(t, 1, extrapolate(1, 1, 1, 3))
	assert.EqualValues(t, 1, extrapolate(1, 1, 1, 4))
	assert.EqualValues(t, 1, extrapolate(1, 1, 1, 5))

	assert.EqualValues(t, 9, extrapolate(3, 5, 7, 1))
	assert.EqualValues(t, 11, extrapolate(3, 5, 7, 2))
	assert.EqualValues(t, 13, extrapolate(3, 5, 7, 3))
	assert.EqualValues(t, 15, extrapolate(3, 5, 7, 4))
	assert.EqualValues(t, 17, extrapolate(3, 5, 7, 5))

	assert.EqualValues(t, 49, extrapolate(16, 25, 36, 1))
	assert.EqualValues(t, 64, extrapolate(16, 25, 36, 2))
	assert.EqualValues(t, 81, extrapolate(16, 25, 36, 3))
	assert.EqualValues(t, 100, extrapolate(16, 25, 36, 4))
	assert.EqualValues(t, 121, extrapolate(16, 25, 36, 5))
}
