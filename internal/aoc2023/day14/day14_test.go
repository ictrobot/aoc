package day14

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const Part1 = 136
const Part2 = 64

func TestDay14_ParseExample(t *testing.T) {
	d1 := Day14{}
	d1.ParseExample()

	d2 := Day14{}
	d2.ParseExample()
	d2.ParseExample()

	assert.Equal(t, d1, d2, "should be idempotent")
}

func BenchmarkDay14_ParseExample(b *testing.B) {
	d := Day14{}
	for i := 0; i < b.N; i++ {
		d.ParseExample()
	}
}

func TestDay14_Part1(t *testing.T) {
	d := Day14{}
	d.ParseExample()

	assert.EqualValues(t, Part1, d.Part1())
}

func BenchmarkDay14_Part1(b *testing.B) {
	d := Day14{}
	d.ParseExample()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		assert.EqualValues(b, Part1, d.Part1())
	}
}

func TestDay14_Part2(t *testing.T) {
	d := Day14{}
	d.ParseExample()

	assert.EqualValues(t, Part2, d.Part2())
}

func BenchmarkDay14_Part2(b *testing.B) {
	d := Day14{}
	d.ParseExample()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		assert.EqualValues(b, Part2, d.Part2())
	}
}
