package day05

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const Part1 = 35
const Part2 = 46

func TestDay05_ParseExample(t *testing.T) {
	d1 := Day05{}
	d1.ParseExample()

	d2 := Day05{}
	d2.ParseExample()
	d2.ParseExample()

	assert.Equal(t, d1, d2, "should be idempotent")
}

func BenchmarkDay05_ParseExample(b *testing.B) {
	d := Day05{}
	for i := 0; i < b.N; i++ {
		d.ParseExample()
	}
}

func TestDay05_Part1(t *testing.T) {
	d := Day05{}
	d.ParseExample()

	assert.EqualValues(t, Part1, d.Part1())
}

func BenchmarkDay05_Part1(b *testing.B) {
	d := Day05{}
	d.ParseExample()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		assert.EqualValues(b, Part1, d.Part1())
	}
}

func TestDay05_Part2(t *testing.T) {
	d := Day05{}
	d.ParseExample()

	assert.EqualValues(t, Part2, d.Part2())
}

func BenchmarkDay05_Part2(b *testing.B) {
	d := Day05{}
	d.ParseExample()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		assert.EqualValues(b, Part2, d.Part2())
	}
}
