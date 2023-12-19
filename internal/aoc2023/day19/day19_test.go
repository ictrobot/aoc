package day19

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const Part1 = 19114
const Part2 int64 = 167409079868000

func TestDay19_ParseExample(t *testing.T) {
	d1 := Day19{}
	d1.ParseExample()

	d2 := Day19{}
	d2.ParseExample()
	d2.ParseExample()

	assert.Equal(t, d1, d2, "should be idempotent")
}

func BenchmarkDay19_ParseExample(b *testing.B) {
	d := Day19{}
	for i := 0; i < b.N; i++ {
		d.ParseExample()
	}
}

func TestDay19_Part1(t *testing.T) {
	d := Day19{}
	d.ParseExample()

	assert.EqualValues(t, Part1, d.Part1())
}

func BenchmarkDay19_Part1(b *testing.B) {
	d := Day19{}
	d.ParseExample()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		assert.EqualValues(b, Part1, d.Part1())
	}
}

func TestDay19_Part2(t *testing.T) {
	d := Day19{}
	d.ParseExample()

	assert.EqualValues(t, Part2, d.Part2())
}

func BenchmarkDay19_Part2(b *testing.B) {
	d := Day19{}
	d.ParseExample()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		assert.EqualValues(b, Part2, d.Part2())
	}
}
