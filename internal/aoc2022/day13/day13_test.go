package day13

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const Part1 = 13
const Part2 = 140

func TestDay13_ParseExample(t *testing.T) {
	d1 := Day13{}
	d1.ParseExample()

	d2 := Day13{}
	d2.ParseExample()
	d2.ParseExample()

	assert.Equal(t, d1, d2, "should be idempotent")
}

func BenchmarkDay13_ParseExample(b *testing.B) {
	d := Day13{}
	for i := 0; i < b.N; i++ {
		d.ParseExample()
	}
}

func TestDay13_Part1(t *testing.T) {
	d := Day13{}
	d.ParseExample()

	assert.EqualValues(t, Part1, d.Part1())
}

func BenchmarkDay13_Part1(b *testing.B) {
	d := Day13{}
	d.ParseExample()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		assert.EqualValues(b, Part1, d.Part1())
	}
}

func TestDay13_Part2(t *testing.T) {
	d := Day13{}
	d.ParseExample()

	assert.EqualValues(t, Part2, d.Part2())
}

func BenchmarkDay13_Part2(b *testing.B) {
	d := Day13{}
	d.ParseExample()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		assert.EqualValues(b, Part2, d.Part2())
	}
}
