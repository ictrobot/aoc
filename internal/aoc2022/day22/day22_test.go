package day22

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const Part1 = 6032

func TestDay22_ParseExample(t *testing.T) {
	d1 := Day22{}
	d1.ParseExample()

	d2 := Day22{}
	d2.ParseExample()
	d2.ParseExample()

	assert.Equal(t, d1, d2, "should be idempotent")
}

func BenchmarkDay22_ParseExample(b *testing.B) {
	d := Day22{}
	for i := 0; i < b.N; i++ {
		d.ParseExample()
	}
}

func TestDay22_Part1(t *testing.T) {
	d := Day22{}
	d.ParseExample()

	assert.EqualValues(t, Part1, d.Part1())
}

func BenchmarkDay22_Part1(b *testing.B) {
	d := Day22{}
	d.ParseExample()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		assert.EqualValues(b, Part1, d.Part1())
	}
}

// Part 2 unsupported on example input
