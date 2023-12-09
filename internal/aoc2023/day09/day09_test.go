package day09

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const Part1 = 114
const Part2 = 2

func TestDay09_ParseExample(t *testing.T) {
	d1 := Day09{}
	d1.ParseExample()

	d2 := Day09{}
	d2.ParseExample()
	d2.ParseExample()

	assert.Equal(t, d1, d2, "should be idempotent")
}

func BenchmarkDay09_ParseExample(b *testing.B) {
	d := Day09{}
	for i := 0; i < b.N; i++ {
		d.ParseExample()
	}
}

func TestDay09_Part1(t *testing.T) {
	d := Day09{}
	d.ParseExample()

	assert.EqualValues(t, Part1, d.Part1())
}

func BenchmarkDay09_Part1(b *testing.B) {
	d := Day09{}
	d.ParseExample()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		assert.EqualValues(b, Part1, d.Part1())
	}
}

func TestDay09_Part2(t *testing.T) {
	d := Day09{}
	d.ParseExample()

	assert.EqualValues(t, Part2, d.Part2())
}

func BenchmarkDay09_Part2(b *testing.B) {
	d := Day09{}
	d.ParseExample()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		assert.EqualValues(b, Part2, d.Part2())
	}
}
