package parse

import (
	"github.com/samber/lo"
	"strconv"
)

func Int(s string) int {
	v, err := strconv.ParseInt(s, 10, strconv.IntSize)
	return int(lo.Must(v, err, "parsing int"))
}

func Int8(s string) int8 {
	v, err := strconv.ParseInt(s, 10, 8)
	return int8(lo.Must(v, err, "parsing int8"))
}

func Int16(s string) int16 {
	v, err := strconv.ParseInt(s, 10, 16)
	return int16(lo.Must(v, err, "parsing int16"))
}

func Int32(s string) int32 {
	v, err := strconv.ParseInt(s, 10, 32)
	return int32(lo.Must(v, err, "parsing int32"))
}

func Int64(s string) int64 {
	v, err := strconv.ParseInt(s, 10, 64)
	return lo.Must(v, err, "parsing int64")
}

func Uint(s string) uint {
	v, err := strconv.ParseUint(s, 10, strconv.IntSize)
	return uint(lo.Must(v, err, "parsing uint"))
}

func Uint8(s string) uint8 {
	v, err := strconv.ParseUint(s, 10, 8)
	return uint8(lo.Must(v, err, "parsing uint8"))
}

func Uint16(s string) uint16 {
	v, err := strconv.ParseUint(s, 10, 16)
	return uint16(lo.Must(v, err, "parsing uint16"))
}

func Uint32(s string) uint32 {
	v, err := strconv.ParseUint(s, 10, 32)
	return uint32(lo.Must(v, err, "parsing uint32"))
}

func Uint64(s string) uint64 {
	v, err := strconv.ParseUint(s, 10, 64)
	return lo.Must(v, err, "parsing uint64")
}

func Ints(s []string) []int {
	r := make([]int, len(s))
	for i, v := range s {
		r[i] = Int(v)
	}
	return r
}

func Int8s(s []string) []int8 {
	r := make([]int8, len(s))
	for i, v := range s {
		r[i] = Int8(v)
	}
	return r
}

func Int16s(s []string) []int16 {
	r := make([]int16, len(s))
	for i, v := range s {
		r[i] = Int16(v)
	}
	return r
}

func Int32s(s []string) []int32 {
	r := make([]int32, len(s))
	for i, v := range s {
		r[i] = Int32(v)
	}
	return r
}

func Int64s(s []string) []int64 {
	r := make([]int64, len(s))
	for i, v := range s {
		r[i] = Int64(v)
	}
	return r
}

func Uints(s []string) []uint {
	r := make([]uint, len(s))
	for i, v := range s {
		r[i] = Uint(v)
	}
	return r
}

func Uint8s(s []string) []uint8 {
	r := make([]uint8, len(s))
	for i, v := range s {
		r[i] = Uint8(v)
	}
	return r
}

func Uint16s(s []string) []uint16 {
	r := make([]uint16, len(s))
	for i, v := range s {
		r[i] = Uint16(v)
	}
	return r
}

func Uint32s(s []string) []uint32 {
	r := make([]uint32, len(s))
	for i, v := range s {
		r[i] = Uint32(v)
	}
	return r
}

func Uint64s(s []string) []uint64 {
	r := make([]uint64, len(s))
	for i, v := range s {
		r[i] = Uint64(v)
	}
	return r
}

func IntStrings(s string) []string {
	return intStrings(s, true)
}

func ExtractInts(s string) []int {
	return Ints(IntStrings(s))
}

func ExtractInt64s(s string) []int64 {
	return Int64s(IntStrings(s))
}

func UintStrings(s string) []string {
	return intStrings(s, false)
}

func ExtractUints(s string) []uint {
	return Uints(UintStrings(s))
}

func ExtractUint64s(s string) []uint64 {
	return Uint64s(UintStrings(s))
}

func ExtractDigits(s string) []int {
	r := make([]int, 0, len(s))
	for _, c := range s {
		if c >= '0' && c <= '9' {
			r = append(r, int(c-'0'))
		}
	}
	return r
}

func intStrings(s string, signed bool) []string {
	if s == "" {
		return nil
	}

	i := 0
	results := make([]string, 0, len(s)/2)
	for i < len(s) {
		for i < len(s) && (s[i] < '0' || s[i] > '9') {
			i++
		}

		if i >= len(s) {
			break
		}

		j := i + 1
		for j < len(s) && s[j] >= '0' && s[j] <= '9' {
			j++
		}

		if signed && i > 0 && (s[i-1] == '-' || s[i-1] == '+') {
			i--
		}

		results = append(results, s[i:j])
		i = j
	}

	return results
}
