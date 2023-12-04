package numbers

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFloatConstants(t *testing.T) {
	assert.Equal(t, []int64{
		Float32MaxInt - 2,
		Float32MaxInt - 1,
		Float32MaxInt,
		Float32MaxInt, // loss of precision
		Float32MaxInt + 2,
	}, []int64{
		int64(float32(Float32MaxInt - 2)),
		int64(float32(Float32MaxInt - 1)),
		int64(float32(Float32MaxInt)),
		int64(float32(Float32MaxInt + 1)),
		int64(float32(Float32MaxInt + 2)),
	})

	assert.Equal(t, []int64{
		Float64MaxInt - 2,
		Float64MaxInt - 1,
		Float64MaxInt,
		Float64MaxInt, // loss of precision
		Float64MaxInt + 2,
	}, []int64{
		int64(float64(Float64MaxInt - 2)),
		int64(float64(Float64MaxInt - 1)),
		int64(float64(Float64MaxInt)),
		int64(float64(Float64MaxInt + 1)),
		int64(float64(Float64MaxInt + 2)),
	})

	assert.Equal(t, []int64{
		-Float32MaxInt - 2,
		-Float32MaxInt, // loss of precision
		-Float32MaxInt,
		-Float32MaxInt + 1,
		-Float32MaxInt + 2,
	}, []int64{
		int64(float32(-Float32MaxInt - 2)),
		int64(float32(-Float32MaxInt - 1)),
		int64(float32(-Float32MaxInt)),
		int64(float32(-Float32MaxInt + 1)),
		int64(float32(-Float32MaxInt + 2)),
	})

	assert.Equal(t, []int64{
		-Float64MaxInt - 2,
		-Float64MaxInt, // loss of precision
		-Float64MaxInt,
		-Float64MaxInt + 1,
		-Float64MaxInt + 2,
	}, []int64{
		int64(float64(-Float64MaxInt - 2)),
		int64(float64(-Float64MaxInt - 1)),
		int64(float64(-Float64MaxInt)),
		int64(float64(-Float64MaxInt + 1)),
		int64(float64(-Float64MaxInt + 2)),
	})
}
