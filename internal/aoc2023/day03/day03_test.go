package day03

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const Part1 = 4361
const Part2 = 467835

func TestDay03_ParseExample(t *testing.T) {
	d1 := Day03{}
	d1.ParseExample()

	d2 := Day03{}
	d2.ParseExample()
	d2.ParseExample()

	assert.Equal(t, d1, d2, "should be idempotent")
}

func BenchmarkDay03_ParseExample(b *testing.B) {
	d := Day03{}
	for i := 0; i < b.N; i++ {
		d.ParseExample()
	}
}

func TestDay03_Part1(t *testing.T) {
	d := Day03{}
	d.ParseExample()

	assert.EqualValues(t, Part1, d.Part1())
}

func BenchmarkDay03_Part1(b *testing.B) {
	d := Day03{}
	d.ParseExample()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		assert.EqualValues(b, Part1, d.Part1())
	}
}

func TestDay03_Part2(t *testing.T) {
	d := Day03{}
	d.ParseExample()

	assert.EqualValues(t, Part2, d.Part2())
}

func BenchmarkDay03_Part2(b *testing.B) {
	d := Day03{}
	d.ParseExample()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		assert.EqualValues(b, Part2, d.Part2())
	}
}
