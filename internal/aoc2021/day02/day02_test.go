package day02

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const Part1 = 150
const Part2 = 900

func TestDay02_ParseExample(t *testing.T) {
	d1 := Day02{}
	d1.ParseExample()

	d2 := Day02{}
	d2.ParseExample()
	d2.ParseExample()

	assert.Equal(t, d1, d2, "should be idempotent")
}

func BenchmarkDay02_ParseExample(b *testing.B) {
	d := Day02{}
	for i := 0; i < b.N; i++ {
		d.ParseExample()
	}
}

func TestDay02_Part1(t *testing.T) {
	d := Day02{}
	d.ParseExample()

	assert.EqualValues(t, Part1, d.Part1())
}

func BenchmarkDay02_Part1(b *testing.B) {
	d := Day02{}
	d.ParseExample()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		assert.EqualValues(b, Part1, d.Part1())
	}
}

func TestDay02_Part2(t *testing.T) {
	d := Day02{}
	d.ParseExample()

	assert.EqualValues(t, Part2, d.Part2())
}

func BenchmarkDay02_Part2(b *testing.B) {
	d := Day02{}
	d.ParseExample()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		assert.EqualValues(b, Part2, d.Part2())
	}
}
