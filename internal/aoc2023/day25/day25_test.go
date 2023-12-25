package day25

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const Part1 = 54

func TestDay25_ParseExample(t *testing.T) {
	d1 := Day25{}
	d1.ParseExample()

	d2 := Day25{}
	d2.ParseExample()
	d2.ParseExample()

	assert.Equal(t, d1, d2, "should be idempotent")
}

func BenchmarkDay25_ParseExample(b *testing.B) {
	d := Day25{}
	for i := 0; i < b.N; i++ {
		d.ParseExample()
	}
}

func TestDay25_Part1(t *testing.T) {
	d := Day25{}
	d.ParseExample()

	assert.EqualValues(t, Part1, d.Part1())
}

func BenchmarkDay25_Part1(b *testing.B) {
	d := Day25{}
	d.ParseExample()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		assert.EqualValues(b, Part1, d.Part1())
	}
}
