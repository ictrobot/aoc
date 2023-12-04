package parse

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFloat32(t *testing.T) {
	assert.Equal(t, float32(1.5), Float32("1.5"))
	assert.Equal(t, float32(0.25), Float32("0.25"))
	assert.Equal(t, float32(64), Float32("64"))
	assert.Equal(t, float32(10000000000), Float32("1e10"))

	assert.Panics(t, func() {
		Float32("")
	})
	assert.Panics(t, func() {
		Float32("a")
	})
}

func TestFloat64(t *testing.T) {
	assert.Equal(t, 1.5, Float64("1.5"))
	assert.Equal(t, 0.25, Float64("0.25"))
	assert.Equal(t, 64.0, Float64("64"))
	assert.Equal(t, 10000000000.0, Float64("1e10"))

	assert.Panics(t, func() {
		Float64("")
	})
	assert.Panics(t, func() {
		Float64("a")
	})
}

func TestFloat32s(t *testing.T) {
	assert.Equal(t, []float32{}, Float32s(nil))
	assert.Equal(t, []float32{}, Float32s([]string{}))
	assert.Equal(t, []float32{-0.25, 0.25, 0.5}, Float32s([]string{"-0.25", "0.25", "0.5"}))

	assert.Panics(t, func() {
		Float32s([]string{"a"})
	})
}

func TestFloat64s(t *testing.T) {
	assert.Equal(t, []float64{}, Float64s(nil))
	assert.Equal(t, []float64{}, Float64s([]string{}))
	assert.Equal(t, []float64{-0.25, 0.25, 0.5}, Float64s([]string{"-0.25", "0.25", "0.5"}))

	assert.Panics(t, func() {
		Float64s([]string{"a"})
	})
}

func TestExtractFloat32s(t *testing.T) {
	assert.Equal(t, []float32{0.5, 1.0, 1.5, 2.0}, ExtractFloat32s("0.5abc1\n  1.5:2"))
}

func TestExtractFloat64s(t *testing.T) {
	assert.Equal(t, []float64{0.5, 1.0, 1.5, 2.0}, ExtractFloat64s("0.5abc1\n  1.5:2"))
}
