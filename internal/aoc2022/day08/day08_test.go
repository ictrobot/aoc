package day08

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const Part1 = 21
const Part2 = 8

func TestDay08_ParseExample(t *testing.T) {
	d1 := Day08{}
	d1.ParseExample()

	d2 := Day08{}
	d2.ParseExample()
	d2.ParseExample()

	assert.Equal(t, d1, d2, "should be idempotent")
}

func BenchmarkDay08_ParseExample(b *testing.B) {
	d := Day08{}
	for i := 0; i < b.N; i++ {
		d.ParseExample()
	}
}

func TestDay08_Part1(t *testing.T) {
	d := Day08{}
	d.ParseExample()

	assert.EqualValues(t, Part1, d.Part1())
}

func BenchmarkDay08_Part1(b *testing.B) {
	d := Day08{}
	d.ParseExample()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		assert.EqualValues(b, Part1, d.Part1())
	}
}

func TestDay08_Part2(t *testing.T) {
	d := Day08{}
	d.ParseExample()

	assert.EqualValues(t, Part2, d.Part2())
}

func BenchmarkDay08_Part2(b *testing.B) {
	d := Day08{}
	d.ParseExample()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		assert.EqualValues(b, Part2, d.Part2())
	}
}
