package day21

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const Part1 = 16

//const Part2 = 0

func TestDay21_ParseExample(t *testing.T) {
	d1 := Day21{}
	d1.ParseExample()

	d2 := Day21{}
	d2.ParseExample()
	d2.ParseExample()

	assert.Equal(t, d1, d2, "should be idempotent")
}

func BenchmarkDay21_ParseExample(b *testing.B) {
	d := Day21{}
	for i := 0; i < b.N; i++ {
		d.ParseExample()
	}
}

func TestDay21_Part1(t *testing.T) {
	d := Day21{}
	d.ParseExample()

	assert.EqualValues(t, Part1, d.Part1())
}

func BenchmarkDay21_Part1(b *testing.B) {
	d := Day21{}
	d.ParseExample()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		assert.EqualValues(b, Part1, d.Part1())
	}
}

//func TestDay21_Part2(t *testing.T) {
//	d := Day21{}
//	d.ParseExample()
//
//	assert.EqualValues(t, Part2, d.Part2())
//}
//
//func BenchmarkDay21_Part2(b *testing.B) {
//	d := Day21{}
//	d.ParseExample()
//
//	b.ResetTimer()
//	for i := 0; i < b.N; i++ {
//		assert.EqualValues(b, Part2, d.Part2())
//	}
//}
