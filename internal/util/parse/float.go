package parse

import (
	"github.com/samber/lo"
	"regexp"
	"strconv"
)

var decimalRegexp = regexp.MustCompile(`[+-]?[0-9]+(\.[0-9]+)?`)

func Float32(s string) float32 {
	v, err := strconv.ParseFloat(s, 32)
	return float32(lo.Must(v, err, "parsing float"))
}

func Float64(s string) float64 {
	v, err := strconv.ParseFloat(s, 64)
	return lo.Must(v, err, "parsing float")
}

func Float32s(s []string) []float32 {
	r := make([]float32, len(s))
	for i, v := range s {
		r[i] = Float32(v)
	}
	return r
}

func Float64s(s []string) []float64 {
	r := make([]float64, len(s))
	for i, v := range s {
		r[i] = Float64(v)
	}
	return r
}

func ExtractFloat32s(s string) []float32 {
	return Float32s(decimalRegexp.FindAllString(s, -1))
}

func ExtractFloat64s(s string) []float64 {
	return Float64s(decimalRegexp.FindAllString(s, -1))
}
