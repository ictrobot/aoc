package numbers

import "golang.org/x/exp/constraints"

// IntMod returns the smallest non-negative remainder (e.g. IntMod(-1, 5) = 4)
// For non-negative a values this is equivalent to the % operator
func IntMod[T constraints.Integer](a, b T) T {
	return (a%b + b) % b
}

func IntAbs[T constraints.Integer](x T) T {
	if x < 0 {
		return -x
	}
	return x
}

func IntAbsDiff[T constraints.Integer](x, y T) T {
	if x > y {
		return x - y
	}
	return y - x
}

func IntSign[T constraints.Integer](x T) int {
	if x > 0 {
		return 1
	} else if x < 0 {
		return -1
	}
	return 0
}
