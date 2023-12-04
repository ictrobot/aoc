package day01

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const Part1 = 24000
const Part2 = 45000

func TestDay01_ParseExample(t *testing.T) {
	d1 := Day01{}
	d1.ParseExample()

	d2 := Day01{}
	d2.ParseExample()
	d2.ParseExample()

	assert.Equal(t, d1, d2, "should be idempotent")
}

func BenchmarkDay01_ParseExample(b *testing.B) {
	d := Day01{}
	for i := 0; i < b.N; i++ {
		d.ParseExample()
	}
}

func TestDay01_Part1(t *testing.T) {
	d := Day01{}
	d.ParseExample()

	assert.EqualValues(t, Part1, d.Part1())
}

func BenchmarkDay01_Part1(b *testing.B) {
	d := Day01{}
	d.ParseExample()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		assert.EqualValues(b, Part1, d.Part1())
	}
}

func TestDay01_Part2(t *testing.T) {
	d := Day01{}
	d.ParseExample()

	assert.EqualValues(t, Part2, d.Part2())
}

func BenchmarkDay01_Part2(b *testing.B) {
	d := Day01{}
	d.ParseExample()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		assert.EqualValues(b, Part2, d.Part2())
	}
}
