package day18

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const Part1 = 64
const Part2 = 58

func TestDay18_ParseExample(t *testing.T) {
	d1 := Day18{}
	d1.ParseExample()

	d2 := Day18{}
	d2.ParseExample()
	d2.ParseExample()

	assert.Equal(t, d1, d2, "should be idempotent")
}

func BenchmarkDay18_ParseExample(b *testing.B) {
	d := Day18{}
	for i := 0; i < b.N; i++ {
		d.ParseExample()
	}
}

func TestDay18_Part1(t *testing.T) {
	d := Day18{}
	d.ParseExample()

	assert.EqualValues(t, Part1, d.Part1())
}

func BenchmarkDay18_Part1(b *testing.B) {
	d := Day18{}
	d.ParseExample()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		assert.EqualValues(b, Part1, d.Part1())
	}
}

func TestDay18_Part2(t *testing.T) {
	d := Day18{}
	d.ParseExample()

	assert.EqualValues(t, Part2, d.Part2())
}

func BenchmarkDay18_Part2(b *testing.B) {
	d := Day18{}
	d.ParseExample()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		assert.EqualValues(b, Part2, d.Part2())
	}
}
