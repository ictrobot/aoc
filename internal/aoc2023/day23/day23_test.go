package day23

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const Part1 = 94
const Part2 = 154

func TestDay23_ParseExample(t *testing.T) {
	d1 := Day23{}
	d1.ParseExample()

	d2 := Day23{}
	d2.ParseExample()
	d2.ParseExample()

	assert.Equal(t, d1, d2, "should be idempotent")
}

func BenchmarkDay23_ParseExample(b *testing.B) {
	d := Day23{}
	for i := 0; i < b.N; i++ {
		d.ParseExample()
	}
}

func TestDay23_Part1(t *testing.T) {
	d := Day23{}
	d.ParseExample()

	assert.EqualValues(t, Part1, d.Part1())
}

func BenchmarkDay23_Part1(b *testing.B) {
	d := Day23{}
	d.ParseExample()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		assert.EqualValues(b, Part1, d.Part1())
	}
}

func TestDay23_Part2(t *testing.T) {
	d := Day23{}
	d.ParseExample()

	assert.EqualValues(t, Part2, d.Part2())
}

func BenchmarkDay23_Part2(b *testing.B) {
	d := Day23{}
	d.ParseExample()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		assert.EqualValues(b, Part2, d.Part2())
	}
}
