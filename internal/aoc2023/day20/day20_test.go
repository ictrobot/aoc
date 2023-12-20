package day20

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const Part1 = 11687500

func TestDay20_ParseExample(t *testing.T) {
	d1 := Day20{}
	d1.ParseExample()

	d2 := Day20{}
	d2.ParseExample()
	d2.ParseExample()

	assert.Equal(t, d1, d2, "should be idempotent")
}

func BenchmarkDay20_ParseExample(b *testing.B) {
	d := Day20{}
	for i := 0; i < b.N; i++ {
		d.ParseExample()
	}
}

func TestDay20_Part1(t *testing.T) {
	d := Day20{}
	d.ParseExample()

	assert.EqualValues(t, Part1, d.Part1())
}

func BenchmarkDay20_Part1(b *testing.B) {
	d := Day20{}
	d.ParseExample()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		assert.EqualValues(b, Part1, d.Part1())
	}
}
