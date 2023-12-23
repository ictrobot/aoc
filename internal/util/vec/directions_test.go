package vec

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDirectionConstants(t *testing.T) {
	assert.Equal(t, I2[int]{-1, 0}, I2Directions[NegX])
	assert.Equal(t, I2[int]{1, 0}, I2Directions[PosX])
	assert.Equal(t, I2[int]{0, -1}, I2Directions[NegY])
	assert.Equal(t, I2[int]{0, 1}, I2Directions[PosY])
	assert.Equal(t, I2Directions[:], I2DirectionsWithDiagonals[:len(I2Directions)])

	assert.Equal(t, I3[int]{-1, 0, 0}, I3Directions[NegX])
	assert.Equal(t, I3[int]{1, 0, 0}, I3Directions[PosX])
	assert.Equal(t, I3[int]{0, -1, 0}, I3Directions[NegY])
	assert.Equal(t, I3[int]{0, 1, 0}, I3Directions[PosY])
	assert.Equal(t, I3[int]{0, 0, -1}, I3Directions[NegZ])
	assert.Equal(t, I3[int]{0, 0, 1}, I3Directions[PosZ])
	assert.Equal(t, I3Directions[:], I3DirectionsWithDiagonals[:len(I3Directions)])
}

func TestDirectionOpposites(t *testing.T) {
	assert.Equal(t, len(I2DirectionsWithDiagonals), len(I2Opposites))
	for i := 0; i < len(I2Opposites); i++ {
		v1 := I2DirectionsWithDiagonals[i]
		v2 := I2DirectionsWithDiagonals[I2Opposites[i]]
		assert.Equal(t, v1.Mul(-1), v2)
	}

	assert.Equal(t, len(I3DirectionsWithDiagonals), len(I3Opposites))
	for i := 0; i < len(I3Opposites); i++ {
		v1 := I3DirectionsWithDiagonals[i]
		v2 := I3DirectionsWithDiagonals[I3Opposites[i]]
		assert.Equal(t, v1.Mul(-1), v2)
	}
}
