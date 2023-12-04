package day11

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const Part1 = 10605
const Part2 = 2713310158

func TestDay11_ParseExample(t *testing.T) {
	d1 := Day11{}
	d1.ParseExample()

	d2 := Day11{}
	d2.ParseExample()
	d2.ParseExample()

	assert.Equal(t, d1, d2, "should be idempotent")
}

func BenchmarkDay11_ParseExample(b *testing.B) {
	d := Day11{}
	for i := 0; i < b.N; i++ {
		d.ParseExample()
	}
}

func TestDay11_Part1(t *testing.T) {
	d := Day11{}
	d.ParseExample()

	assert.EqualValues(t, Part1, d.Part1())
}

func BenchmarkDay11_Part1(b *testing.B) {
	d := Day11{}
	d.ParseExample()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		assert.EqualValues(b, Part1, d.Part1())
	}
}

func TestDay11_Part2(t *testing.T) {
	d := Day11{}
	d.ParseExample()

	assert.EqualValues(t, Part2, d.Part2())
}

func BenchmarkDay11_Part2(b *testing.B) {
	d := Day11{}
	d.ParseExample()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		assert.EqualValues(b, Part2, d.Part2())
	}
}
