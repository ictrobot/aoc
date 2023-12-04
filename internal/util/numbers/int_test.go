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
