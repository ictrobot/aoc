package day24

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const Part1 = 18
const Part2 = 54

func TestDay24_ParseExample(t *testing.T) {
	d1 := Day24{}
	d1.ParseExample()

	d2 := Day24{}
	d2.ParseExample()
	d2.ParseExample()

	assert.Equal(t, d1, d2, "should be idempotent")
}

func BenchmarkDay24_ParseExample(b *testing.B) {
	d := Day24{}
	for i := 0; i < b.N; i++ {
		d.ParseExample()
	}
}

func TestDay24_Part1(t *testing.T) {
	d := Day24{}
	d.ParseExample()

	assert.EqualValues(t, Part1, d.Part1())
}

func BenchmarkDay24_Part1(b *testing.B) {
	d := Day24{}
	d.ParseExample()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		assert.EqualValues(b, Part1, d.Part1())
	}
}

func TestDay24_Part2(t *testing.T) {
	d := Day24{}
	d.ParseExample()

	assert.EqualValues(t, Part2, d.Part2())
}

func BenchmarkDay24_Part2(b *testing.B) {
	d := Day24{}
	d.ParseExample()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		assert.EqualValues(b, Part2, d.Part2())
	}
}
