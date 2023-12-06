package vec

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func TestNewGrid(t *testing.T) {
	assert.Equal(t, &Grid[int]{
		init: true,
		xMin: 2,
		yMin: 2,
		xMax: 5,
		yMax: 5,
		s:    [][]int{{0, 0, 0, 0}, {0, 0, 0, 0}, {0, 0, 0, 0}, {0, 0, 0, 0}},
	}, NewGrid[int](2, 2, 5, 5))
}

func TestGrid_Clone(t *testing.T) {
	g1 := NewGrid[int](10, 10, 15, 15)
	g1.SetInts(11, 12, 1)

	g2 := g1.Clone()
	g2.SetInts(11, 12, 2)

	assert.Equal(t, 1, g1.GetInts(11, 12))
	assert.Equal(t, 2, g2.GetInts(11, 12))

	g3 := new(Grid[string])
	assert.Equal(t, &Grid[string]{init: false}, g3.Clone())
}

func TestGrid_Get(t *testing.T) {
	g1 := new(Grid[float32])
	for i := 0; i < 10; i++ {
		g1.SetInts(i*2, i*3, float32(i)/4)
	}

	for i := 0; i < 10; i++ {
		assert.Equal(t, float32(i)/4, g1.Get(I2[int]{i * 2, i * 3}))
	}
}

func TestGrid_GetInts(t *testing.T) {
	g1 := new(Grid[string])
	for i := 20; i >= 0; i-- {
		g1.SetInts(-i, i, strconv.Itoa(i))
	}

	for i := 0; i < 20; i++ {
		assert.Equal(t, strconv.Itoa(i), g1.GetInts(-i, i))
	}
}

func TestGrid_Set(t *testing.T) {
	g1 := NewGrid[rune](10, 10, 12, 12)
	assert.Equal(t, [][]rune{
		{0, 0, 0},
		{0, 0, 0},
		{0, 0, 0},
	}, g1.s)

	assert.EqualValues(t, 0, g1.Set(I2[int]{12, 10}, '#'))
	assert.Equal(t, [][]rune{
		{0, 0, 0},
		{0, 0, 0},
		{'#', 0, 0},
	}, g1.s)

	assert.EqualValues(t, 0, g1.Set(I2[int]{11, 12}, '@'))
	assert.Equal(t, [][]rune{
		{0, 0, 0},
		{0, 0, '@'},
		{'#', 0, 0},
	}, g1.s)

	assert.EqualValues(t, '#', g1.Set(I2[int]{12, 10}, '!'))
	assert.Equal(t, [][]rune{
		{0, 0, 0},
		{0, 0, '@'},
		{'!', 0, 0},
	}, g1.s)

	// setting zero outside current bounds shouldn't expand
	g1.Set(I2[int]{100, 200}, 0)
	assert.Equal(t, &Grid[rune]{
		init: true,
		xMin: 10, yMin: 10,
		xMax: 12, yMax: 12,
		s: [][]rune{
			{0, 0, 0},
			{0, 0, '@'},
			{'!', 0, 0},
		},
	}, g1)

	g1.SetIfZero(I2[int]{201, 16}, 103)
	assert.LessOrEqual(t, g1.xMin, 10)
	assert.LessOrEqual(t, g1.yMin, 10)
	assert.GreaterOrEqual(t, g1.xMax, 201)
	assert.GreaterOrEqual(t, g1.yMax, 16)
}

func TestGrid_SetInts(t *testing.T) {
	g1 := NewGrid[int](-100, -100, -97, -97)
	assert.Equal(t, [][]int{
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
	}, g1.s)

	assert.EqualValues(t, 0, g1.SetInts(-99, -100, 99100))
	assert.Equal(t, [][]int{
		{0, 0, 0, 0},
		{99100, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
	}, g1.s)

	assert.EqualValues(t, 0, g1.SetInts(-97, -97, -1))
	assert.Equal(t, [][]int{
		{0, 0, 0, 0},
		{99100, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, -1},
	}, g1.s)

	assert.EqualValues(t, 99100, g1.SetInts(-99, -100, 78))
	assert.Equal(t, [][]int{
		{0, 0, 0, 0},
		{78, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, -1},
	}, g1.s)
}

func TestGrid_SetIfZero(t *testing.T) {
	g1 := NewGrid[int8](1000, 1001, 1002, 1003)
	assert.Zero(t, g1.GetInts(1002, 1002))

	assert.True(t, g1.SetIfZero(I2[int]{1002, 1002}, 5))
	assert.EqualValues(t, 5, g1.GetInts(1002, 1002))

	assert.False(t, g1.SetIfZero(I2[int]{1002, 1002}, 89))
	assert.EqualValues(t, 5, g1.GetInts(1002, 1002))

	g1.Set(I2[int]{1002, 1002}, 89)
	assert.EqualValues(t, 89, g1.GetInts(1002, 1002))

	// setting zero outside current bounds shouldn't expand
	g1.SetIfZero(I2[int]{-1001, 2002}, 0)
	assert.Equal(t, &Grid[int8]{
		init: true,
		xMin: 1000, yMin: 1001,
		xMax: 1002, yMax: 1003,
		s: [][]int8{
			{0, 0, 0},
			{0, 0, 0},
			{0, 89, 0},
		},
	}, g1)

	g1.SetIfZero(I2[int]{1000, 1004}, 103)
	assert.LessOrEqual(t, g1.xMin, 1000)
	assert.LessOrEqual(t, g1.yMin, 1001)
	assert.GreaterOrEqual(t, g1.xMax, 1002)
	assert.GreaterOrEqual(t, g1.yMax, 1004)
}

func TestGrid_SetIfZeroInts(t *testing.T) {
	g1 := NewGrid[float32](0, 0, 5, 5)
	assert.Zero(t, g1.GetInts(2, 4))

	assert.True(t, g1.SetIfZeroInts(2, 4, 32.5))
	assert.EqualValues(t, 32.5, g1.GetInts(2, 4))

	assert.False(t, g1.SetIfZeroInts(2, 4, 8.25))
	assert.EqualValues(t, 32.5, g1.GetInts(2, 4))
}

func TestGrid_Contains(t *testing.T) {
	g1 := NewGrid[int](0, 0, 5, 5)
	g1.SetInts(1, 2, 33)
	g1.SetInts(3, 4, -44)
	g1.SetInts(0, 0, 0)

	assert.True(t, g1.Contains(I2[int]{1, 2}))
	assert.True(t, g1.Contains(I2[int]{3, 4}))
	assert.False(t, g1.Contains(I2[int]{1, 1}))
	assert.False(t, g1.Contains(I2[int]{0, 0}))
}

func TestGrid_ContainsInts(t *testing.T) {
	g1 := NewGrid[bool](0, 0, 5, 5)
	g1.SetInts(4, 1, true)
	g1.SetInts(1, 4, false)
	g1.SetInts(3, 3, true)

	assert.True(t, g1.ContainsInts(4, 1))
	assert.True(t, g1.ContainsInts(3, 3))
	assert.False(t, g1.ContainsInts(1, 4))
	assert.False(t, g1.ContainsInts(5, 5))
}

func TestGrid_Count(t *testing.T) {
	g1 := NewGrid[bool](0, 0, 5, 5)

	for i := 0; i <= 20; i++ {
		g1.SetInts(i/5, (i*i)%3, true)
	}
	assert.Equal(t, 9, g1.Count(true))

	g2 := new(Grid[string])
	assert.Equal(t, 0, g2.Count("anything"))
}

func TestGrid_CountNonZero(t *testing.T) {
	g1 := NewGrid[int64](0, 0, 5, 5)

	for i := 0; i < 6; i++ {
		g1.SetInts(i, i%2, int64(i%3))
	}
	assert.Equal(t, 4, g1.CountNotZero())

	g2 := new(Grid[int64])
	assert.Equal(t, 0, g2.CountNotZero())
}

func TestGrid_Counts(t *testing.T) {
	g1 := new(Grid[string])
	g1.SetInts(2, 5, "a1")
	g1.SetInts(1, 4, "a1")
	g1.SetInts(1, 6, "a1")
	g1.SetInts(5, 0, "a1")
	g1.SetInts(3, 8, "b2")
	g1.SetInts(3, 7, "b2")
	g1.SetInts(1, 4, "c0")

	assert.Equal(t, map[string]int{
		"c0": 1,
		"b2": 2,
		"a1": 3,
	}, g1.Counts())

	g2 := new(Grid[int])
	assert.Equal(t, map[int]int(nil), g2.Counts())
}

func TestGrid_Bounds(t *testing.T) {
	g1 := NewGrid[int32](10, 20, 30, 40)

	xMin, yMin, xMax, yMax := g1.Bounds()
	assert.EqualValues(t, 10, xMin)
	assert.EqualValues(t, 20, yMin)
	assert.EqualValues(t, 30, xMax)
	assert.EqualValues(t, 40, yMax)

	g1.SetInts(100, 50, 213)

	xMin, yMin, xMax, yMax = g1.Bounds()
	assert.LessOrEqual(t, xMin, 10)
	assert.LessOrEqual(t, yMin, 20)
	assert.GreaterOrEqual(t, xMax, 100)
	assert.GreaterOrEqual(t, yMax, 50)
}

func TestGrid_NonZeroBounds(t *testing.T) {
	g1 := new(Grid[int32])
	xMin, yMin, xMax, yMax := g1.NonZeroBounds()
	assert.Zero(t, xMin)
	assert.Zero(t, yMin)
	assert.Zero(t, xMax)
	assert.Zero(t, yMax)

	g1 = NewGrid[int32](100, 200, 110, 210)
	xMin, yMin, xMax, yMax = g1.NonZeroBounds()
	assert.Zero(t, xMin)
	assert.Zero(t, yMin)
	assert.Zero(t, xMax)
	assert.Zero(t, yMax)

	g1.SetInts(108, 204, 1)
	g1.SetInts(102, 205, 2)
	xMin, yMin, xMax, yMax = g1.NonZeroBounds()
	assert.Equal(t, 102, xMin)
	assert.Equal(t, 204, yMin)
	assert.Equal(t, 108, xMax)
	assert.Equal(t, 205, yMax)
}

func TestGrid_Resize(t *testing.T) {
	g1 := NewGrid[int](-2, -2, 2, 2)
	for x := 0; x < 5; x++ {
		for y := 0; y < 5; y++ {
			g1.SetInts(x-2, y-2, ((x+1)*101)+y*10)
		}
	}
	assert.Equal(t, &Grid[int]{init: true,
		xMin: -2, yMin: -2,
		xMax: 2, yMax: 2,
		s: [][]int{
			{101, 111, 121, 131, 141},
			{202, 212, 222, 232, 242},
			{303, 313, 323, 333, 343},
			{404, 414, 424, 434, 444},
			{505, 515, 525, 535, 545},
		},
	}, g1)

	g2 := g1.Resize(-3, -3, 3, 3)
	assert.Equal(t, &Grid[int]{init: true,
		xMin: -3, yMin: -3,
		xMax: 3, yMax: 3,
		s: [][]int{
			{0, 0, 0, 0, 0, 0, 0},
			{0, 101, 111, 121, 131, 141, 0},
			{0, 202, 212, 222, 232, 242, 0},
			{0, 303, 313, 323, 333, 343, 0},
			{0, 404, 414, 424, 434, 444, 0},
			{0, 505, 515, 525, 535, 545, 0},
			{0, 0, 0, 0, 0, 0, 0},
		},
	}, g2)

	g3 := g1.Resize(-1, -5, 1, 5)
	assert.Equal(t, &Grid[int]{init: true,
		xMin: -1, yMin: -5,
		xMax: 1, yMax: 5,
		s: [][]int{
			{0, 0, 0, 202, 212, 222, 232, 242, 0, 0, 0},
			{0, 0, 0, 303, 313, 323, 333, 343, 0, 0, 0},
			{0, 0, 0, 404, 414, 424, 434, 444, 0, 0, 0},
		},
	}, g3)

	g4 := g3.Resize(-4, -1, 4, -1)
	assert.Equal(t, &Grid[int]{init: true,
		xMin: -4, yMin: -1,
		xMax: 4, yMax: -1,
		s: [][]int{
			{0},
			{0},
			{0},
			{212},
			{313},
			{414},
			{0},
			{0},
			{0},
		},
	}, g4)

	g5 := g2.Resize(-2, -2, -2, -2)
	assert.Equal(t, &Grid[int]{init: true,
		xMin: -2, yMin: -2,
		xMax: -2, yMax: -2,
		s: [][]int{
			{101},
		},
	}, g5)

	assert.Panics(t, func() {
		g1.Resize(11, 0, 10, 10)
	})
	assert.Panics(t, func() {
		g1.Resize(-20, -10, -10, -20)
	})
}

func TestGrid_Format(t *testing.T) {
	g1 := NewGrid[int](1, 1, 3, 3)
	g1.SetInts(1, 1, 3)
	g1.SetInts(1, 2, 5)
	g1.SetInts(2, 3, 7)
	g1.SetInts(3, 1, 9)

	s2 := struct{ G *Grid[int] }{g1}
	assert.Equal(t,
		"{&{true 1 1 3 3 [[3 5 0] [0 0 7] [9 0 0]]}}",
		fmt.Sprintf("%v", s2),
	)
	assert.Equal(t,
		"{G:&{init:true xMin:1 yMin:1 xMax:3 yMax:3 s:[[3 5 0] [0 0 7] [9 0 0]]}}",
		fmt.Sprintf("%+v", s2),
	)
	assert.Equal(t,
		"struct { G *vec.Grid[int] }{G:&vec.Grid[int]{init:true, xMin:1, yMin:1, xMax:3, yMax:3, s:[][]int{[]int{3, 5, 0}, []int{0, 0, 7}, []int{9, 0, 0}}}}",
		fmt.Sprintf("%#v", s2),
	)
}
