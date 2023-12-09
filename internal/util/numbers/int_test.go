package numbers

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMod(t *testing.T) {
	assert.Equal(t, 1, IntMod(6, 5))
	assert.Equal(t, 3, IntMod(8, 5))
	assert.Equal(t, 9, IntMod(-1, 10))
	assert.Equal(t, 17, IntMod(-3, 20))
}

func TestIntAbs(t *testing.T) {
	assert.Equal(t, 0, IntAbs(0))
	assert.Equal(t, 2, IntAbs(2))
	assert.Equal(t, 14, IntAbs(-14))
}

func TestIntAbsDiff(t *testing.T) {
	assert.Equal(t, 9, IntAbsDiff(0, 9))
	assert.Equal(t, 2, IntAbsDiff(43, 41))
	assert.Equal(t, 4, IntAbsDiff(-121, -117))
	assert.Equal(t, 1, IntAbsDiff(-63, -64))
}

func TestIntSign(t *testing.T) {
	assert.Equal(t, 1, IntSign(498))
	assert.Equal(t, 1, IntSign(1))
	assert.Equal(t, 0, IntSign(0))
	assert.Equal(t, -1, IntSign(-1))
	assert.Equal(t, -1, IntSign(-255))
}

func TestIntPow(t *testing.T) {
	assert.EqualValues(t, 1, IntPow(5, 0))
	assert.EqualValues(t, 5, IntPow(5, 1))
	assert.EqualValues(t, 25, IntPow(5, 2))
	assert.EqualValues(t, 125, IntPow(5, 3))
	assert.EqualValues(t, 625, IntPow(5, 4))
	assert.EqualValues(t, 15625, IntPow(5, 6))
	assert.EqualValues(t, 256, IntPow(2, 8))

	assert.Panics(t, func() {
		IntPow(10, -1)
	})
}

func TestIntRoundedDiv(t *testing.T) {
	assert.EqualValues(t, 10, IntRoundedDiv(104, 10))
	assert.EqualValues(t, 11, IntRoundedDiv(105, 10))
	assert.EqualValues(t, -10, IntRoundedDiv(-104, 10))
	assert.EqualValues(t, -11, IntRoundedDiv(-105, 10))
	assert.EqualValues(t, -10, IntRoundedDiv(104, -10))
	assert.EqualValues(t, -11, IntRoundedDiv(105, -10))
	assert.EqualValues(t, 10, IntRoundedDiv(-104, -10))
	assert.EqualValues(t, 11, IntRoundedDiv(-105, -10))
}

func TestGCD(t *testing.T) {
	assert.EqualValues(t, 4, GCD[int](4, 8))
	assert.EqualValues(t, 6, GCD[int8](18, 12))
	assert.EqualValues(t, 5, GCD[uint8](15, 25))
	assert.EqualValues(t, 7, GCD[int32](7, 0))
	assert.EqualValues(t, 9, GCD[uint64](0, 9))
	assert.EqualValues(t, 0, GCD[int64](0, 0))
	assert.EqualValues(t, 12, GCD[int](24, 36))
	assert.EqualValues(t, 12, GCD[int](24, -36))
	assert.EqualValues(t, 12, GCD[int](-24, 36))
	assert.EqualValues(t, 12, GCD[int](-24, -36))
}

func TestLCM(t *testing.T) {
	assert.EqualValues(t, 15, LCM[int](3, 5))
	assert.EqualValues(t, 16, LCM[int](8, 16))
	assert.EqualValues(t, 91, LCM[int](13, 7))
	assert.EqualValues(t, 0, LCM[int](0, 3))
	assert.EqualValues(t, 0, LCM[int](4, 0))
	assert.EqualValues(t, 0, LCM[int](0, 0))
	assert.EqualValues(t, 6, LCM[int](2, 3))
	assert.EqualValues(t, 6, LCM[int](-2, 3))
	assert.EqualValues(t, 6, LCM[int](2, -3))
	assert.EqualValues(t, 6, LCM[int](-2, -3))
}
