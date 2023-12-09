package day04

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const Part1 = 4512
const Part2 = 1924

func TestDay04_ParseExample(t *testing.T) {
	d1 := Day04{}
	d1.ParseExample()

	d2 := Day04{}
	d2.ParseExample()
	d2.ParseExample()

	assert.Equal(t, d1, d2, "should be idempotent")
}

func BenchmarkDay04_ParseExample(b *testing.B) {
	d := Day04{}
	for i := 0; i < b.N; i++ {
		d.ParseExample()
	}
}

func TestDay04_Part1(t *testing.T) {
	d := Day04{}
	d.ParseExample()

	assert.EqualValues(t, Part1, d.Part1())
}

func BenchmarkDay04_Part1(b *testing.B) {
	d := Day04{}
	d.ParseExample()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		assert.EqualValues(b, Part1, d.Part1())
	}
}

func TestDay04_Part2(t *testing.T) {
	d := Day04{}
	d.ParseExample()

	assert.EqualValues(t, Part2, d.Part2())
}

func BenchmarkDay04_Part2(b *testing.B) {
	d := Day04{}
	d.ParseExample()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		assert.EqualValues(b, Part2, d.Part2())
	}
}
