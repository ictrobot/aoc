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

// IntPow computes base**exp using exponentiation by squaring
func IntPow[T constraints.Integer](base, exp T) T {
	if exp < 0 {
		panic("exp must be non-negative")
	}

	result := T(1)
	for {
		if exp%2 == 1 {
			result *= base
		}
		exp /= 2
		if exp == 0 {
			return result
		}
		base *= base
	}
}

// IntRoundedDiv divides two integers rounding to the nearest integer, rounding
// half away from zero. This should be equivalent to
//
//	math.Round(float64(a) / float64(b))
//
// which uses the same half rounding method
func IntRoundedDiv[T constraints.Integer](a, b T) T {
	if (a < 0) != (b < 0) {
		return (a - b/2) / b
	}
	return (a + b/2) / b
}

// GCD returns the greatest common divisor of two integers using the Euclidean
// algorithm, e.g. GCD(12, 18) = 6
func GCD[T constraints.Integer](a, b T) T {
	if b < 0 {
		b = -b
	}

	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// LCM returns the least common multiple of two integers, e.g. LCM(12, 18) = 36
func LCM[T constraints.Integer](a, b T) T {
	if a == 0 || b == 0 {
		return 0
	}

	if a < 0 {
		a = -a
	}
	if b < 0 {
		b = -b
	}

	return a * b / GCD(a, b)
}
