package day25

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const Part1 = "2=-1=0"

var snafuTestCases = []struct {
	dec   int64
	snafu string
}{
	{1, "1"},
	{2, "2"},
	{3, "1="},
	{4, "1-"},
	{5, "10"},
	{6, "11"},
	{7, "12"},
	{8, "2="},
	{9, "2-"},
	{10, "20"},
	{11, "21"},
	{15, "1=0"},
	{20, "1-0"},
	{31, "111"},
	{32, "112"},
	{37, "122"},
	{107, "1-12"},
	{198, "2=0="},
	{201, "2=01"},
	{353, "1=-1="},
	{906, "12111"},
	{1257, "20012"},
	{1747, "1=-0-2"},
	{2022, "1=11-2"},
	{12345, "1-0---0"},
	{314159265, "1121-1110-1=0"},
}

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

func TestSnafuDecode(t *testing.T) {
	for _, c := range snafuTestCases {
		assert.Equal(t, c.dec, snafuDecode(c.snafu))
	}
}

func TestSnafuEncode(t *testing.T) {
	for _, c := range snafuTestCases {
		assert.Equal(t, c.snafu, snafuEncode(c.dec))
	}
}
