package day07

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const Part1 = 95437
const Part2 = 24933642

func TestDay07_ParseExample(t *testing.T) {
	d1 := Day07{}
	d1.ParseExample()

	d2 := Day07{}
	d2.ParseExample()
	d2.ParseExample()

	assert.Equal(t, d1, d2, "should be idempotent")
}

func BenchmarkDay07_ParseExample(b *testing.B) {
	d := Day07{}
	for i := 0; i < b.N; i++ {
		d.ParseExample()
	}
}

func TestDay07_Part1(t *testing.T) {
	d := Day07{}
	d.ParseExample()

	assert.EqualValues(t, Part1, d.Part1())
}

func BenchmarkDay07_Part1(b *testing.B) {
	d := Day07{}
	d.ParseExample()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		assert.EqualValues(b, Part1, d.Part1())
	}
}

func TestDay07_Part2(t *testing.T) {
	d := Day07{}
	d.ParseExample()

	assert.EqualValues(t, Part2, d.Part2())
}

func BenchmarkDay07_Part2(b *testing.B) {
	d := Day07{}
	d.ParseExample()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		assert.EqualValues(b, Part2, d.Part2())
	}
}
