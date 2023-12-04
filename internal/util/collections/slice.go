package collections

import (
	"slices"
)

func Clone2D[S ~[][]E, E any](v S) S {
	result := make(S, len(v))
	for i := range v {
		result[i] = slices.Clone(v[i])
	}
	return result
}

// Fill replaces the contents of the given slice a shallow copy of the given value
func Fill[S ~[]E, E any](s S, value E) {
	// for small arrays it is faster to just use a simple for loop
	// for large arrays populating 32 elements and then copying blocks is faster
	// than starting with only one element
	for i := 0; i < min(len(s), 32); i++ {
		s[i] = value
	}

	if len(s) <= 32 {
		return
	}

	for j := 32; j < len(s); j *= 2 {
		copy(s[j:], s[:j])
	}
}
