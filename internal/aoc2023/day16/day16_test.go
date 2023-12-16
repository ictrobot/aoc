package day16

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const Part1 = 46
const Part2 = 51

func TestDay16_ParseExample(t *testing.T) {
	d1 := Day16{}
	d1.ParseExample()

	d2 := Day16{}
	d2.ParseExample()
	d2.ParseExample()

	assert.Equal(t, d1, d2, "should be idempotent")
}

func BenchmarkDay16_ParseExample(b *testing.B) {
	d := Day16{}
	for i := 0; i < b.N; i++ {
		d.ParseExample()
	}
}

func TestDay16_Part1(t *testing.T) {
	d := Day16{}
	d.ParseExample()

	assert.EqualValues(t, Part1, d.Part1())
}

func BenchmarkDay16_Part1(b *testing.B) {
	d := Day16{}
	d.ParseExample()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		assert.EqualValues(b, Part1, d.Part1())
	}
}

func TestDay16_Part2(t *testing.T) {
	d := Day16{}
	d.ParseExample()

	assert.EqualValues(t, Part2, d.Part2())
}

func BenchmarkDay16_Part2(b *testing.B) {
	d := Day16{}
	d.ParseExample()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		assert.EqualValues(b, Part2, d.Part2())
	}
}
