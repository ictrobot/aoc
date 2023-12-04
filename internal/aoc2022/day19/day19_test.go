package day19

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const Part1 = 33
const Part2 = 56 * 62

func TestDay19_Part1(t *testing.T) {
	d := Day19{}
	d.ParseExample()

	assert.EqualValues(t, Part1, d.Part1())
}

func BenchmarkDay19_Part1(b *testing.B) {
	d := Day19{}
	d.ParseExample()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		assert.EqualValues(b, Part1, d.Part1())
	}
}

func TestDay19_Part2(t *testing.T) {
	d := Day19{}
	d.ParseExample()

	assert.EqualValues(t, Part2, d.Part2())
}

func BenchmarkDay19_Part2(b *testing.B) {
	d := Day19{}
	d.ParseExample()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		assert.EqualValues(b, Part2, d.Part2())
	}
}
