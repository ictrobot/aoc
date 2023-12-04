package day06

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const Part1 = 7
const Part2 = 19

func TestDay06_ParseExample(t *testing.T) {
	d1 := Day06{}
	d1.ParseExample()

	d2 := Day06{}
	d2.ParseExample()
	d2.ParseExample()

	assert.Equal(t, d1, d2, "should be idempotent")
}

func BenchmarkDay06_ParseExample(b *testing.B) {
	d := Day06{}
	for i := 0; i < b.N; i++ {
		d.ParseExample()
	}
}

func TestDay06_Part1(t *testing.T) {
	d := Day06{}
	d.ParseExample()

	assert.EqualValues(t, Part1, d.Part1())
}

func BenchmarkDay06_Part1(b *testing.B) {
	d := Day06{}
	d.ParseExample()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		assert.EqualValues(b, Part1, d.Part1())
	}
}

func TestDay06_Part2(t *testing.T) {
	d := Day06{}
	d.ParseExample()

	assert.EqualValues(t, Part2, d.Part2())
}

func BenchmarkDay06_Part2(b *testing.B) {
	d := Day06{}
	d.ParseExample()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		assert.EqualValues(b, Part2, d.Part2())
	}
}
