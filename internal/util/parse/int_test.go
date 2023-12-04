package parse

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInt(t *testing.T) {
	assert.Equal(t, -2147483648, Int("-2147483648"))
	assert.Equal(t, 2147483647, Int("2147483647"))

	assert.Panics(t, func() {
		Int("-9223372036854775809")
	})
	assert.Panics(t, func() {
		Int("9223372036854775808")
	})
	assert.Panics(t, func() {
		Int("1.0")
	})
	assert.Panics(t, func() {
		Int("a")
	})
}

func TestInt8(t *testing.T) {
	assert.Equal(t, int8(-128), Int8("-128"))
	assert.Equal(t, int8(127), Int8("127"))

	assert.Panics(t, func() {
		Int8("-129")
	})
	assert.Panics(t, func() {
		Int8("128")
	})
	assert.Panics(t, func() {
		Int8("1.0")
	})
	assert.Panics(t, func() {
		Int8("a")
	})
}

func TestInt16(t *testing.T) {
	assert.Equal(t, int16(-32768), Int16("-32768"))
	assert.Equal(t, int16(32767), Int16("32767"))

	assert.Panics(t, func() {
		Int16("-32769")
	})
	assert.Panics(t, func() {
		Int16("32768")
	})
	assert.Panics(t, func() {
		Int16("1.0")
	})
	assert.Panics(t, func() {
		Int16("a")
	})
}

func TestInt32(t *testing.T) {
	assert.Equal(t, int32(-2147483648), Int32("-2147483648"))
	assert.Equal(t, int32(2147483647), Int32("2147483647"))

	assert.Panics(t, func() {
		Int32("-2147483649")
	})
	assert.Panics(t, func() {
		Int32("2147483648")
	})
	assert.Panics(t, func() {
		Int32("1.0")
	})
	assert.Panics(t, func() {
		Int32("a")
	})
}

func TestInt64(t *testing.T) {
	assert.Equal(t, int64(-9223372036854775808), Int64("-9223372036854775808"))
	assert.Equal(t, int64(9223372036854775807), Int64("9223372036854775807"))

	assert.Panics(t, func() {
		Int64("-9223372036854775809")
	})
	assert.Panics(t, func() {
		Int64("9223372036854775808")
	})
	assert.Panics(t, func() {
		Int64("1.0")
	})
	assert.Panics(t, func() {
		Int64("a")
	})
}

func TestUint(t *testing.T) {
	assert.Equal(t, uint(0), Uint("0"))
	assert.Equal(t, uint(4294967295), Uint("4294967295"))

	assert.Panics(t, func() {
		Uint("-1")
	})
	assert.Panics(t, func() {
		Uint("18446744073709551616")
	})
	assert.Panics(t, func() {
		Uint("1.0")
	})
	assert.Panics(t, func() {
		Uint("a")
	})
}

func TestUint8(t *testing.T) {
	assert.Equal(t, uint8(0), Uint8("0"))
	assert.Equal(t, uint8(255), Uint8("255"))

	assert.Panics(t, func() {
		Uint8("-1")
	})
	assert.Panics(t, func() {
		Uint8("256")
	})
	assert.Panics(t, func() {
		Uint8("1.0")
	})
	assert.Panics(t, func() {
		Uint8("a")
	})
}

func TestUint16(t *testing.T) {
	assert.Equal(t, uint16(0), Uint16("0"))
	assert.Equal(t, uint16(65535), Uint16("65535"))

	assert.Panics(t, func() {
		Uint16("-1")
	})
	assert.Panics(t, func() {
		Uint16("65536")
	})
	assert.Panics(t, func() {
		Uint16("1.0")
	})
	assert.Panics(t, func() {
		Uint16("a")
	})
}

func TestUint32(t *testing.T) {
	assert.Equal(t, uint32(0), Uint32("0"))
	assert.Equal(t, uint32(4294967295), Uint32("4294967295"))

	assert.Panics(t, func() {
		Uint32("-1")
	})
	assert.Panics(t, func() {
		Uint32("4294967296")
	})
	assert.Panics(t, func() {
		Uint32("1.0")
	})
	assert.Panics(t, func() {
		Uint32("a")
	})
}

func TestUint64(t *testing.T) {
	assert.Equal(t, uint64(0), Uint64("0"))
	assert.Equal(t, uint64(18446744073709551615), Uint64("18446744073709551615"))

	assert.Panics(t, func() {
		Uint64("-1")
	})
	assert.Panics(t, func() {
		Uint64("18446744073709551616")
	})
	assert.Panics(t, func() {
		Uint64("1.0")
	})
	assert.Panics(t, func() {
		Uint64("a")
	})
}

func TestInts(t *testing.T) {
	assert.Equal(t, []int{}, Ints(nil))
	assert.Equal(t, []int{}, Ints([]string{}))
	assert.Equal(t, []int{-2147483648, 18, 2147483647}, Ints([]string{"-2147483648", "18", "2147483647"}))

	assert.Panics(t, func() {
		Ints([]string{"a"})
	})
}

func TestInt8s(t *testing.T) {
	assert.Equal(t, []int8{}, Int8s(nil))
	assert.Equal(t, []int8{}, Int8s([]string{}))
	assert.Equal(t, []int8{-128, 18, 127}, Int8s([]string{"-128", "18", "127"}))

	assert.Panics(t, func() {
		Int8s([]string{"a"})
	})
}

func TestInt16s(t *testing.T) {
	assert.Equal(t, []int16{}, Int16s(nil))
	assert.Equal(t, []int16{}, Int16s([]string{}))
	assert.Equal(t, []int16{-32768, 18, 32767}, Int16s([]string{"-32768", "18", "32767"}))

	assert.Panics(t, func() {
		Int16s([]string{"a"})
	})
}

func TestInt32s(t *testing.T) {
	assert.Equal(t, []int32{}, Int32s(nil))
	assert.Equal(t, []int32{}, Int32s([]string{}))
	assert.Equal(t, []int32{-2147483648, 18, 2147483647}, Int32s([]string{"-2147483648", "18", "2147483647"}))

	assert.Panics(t, func() {
		Int32s([]string{"a"})
	})
}

func TestInt64s(t *testing.T) {
	assert.Equal(t, []int64{}, Int64s(nil))
	assert.Equal(t, []int64{}, Int64s([]string{}))
	assert.Equal(t, []int64{-9223372036854775808, 18, 9223372036854775807}, Int64s([]string{"-9223372036854775808", "18", "9223372036854775807"}))

	assert.Panics(t, func() {
		Int64s([]string{"a"})
	})
}

func TestUints(t *testing.T) {
	assert.Equal(t, []uint{}, Uints(nil))
	assert.Equal(t, []uint{}, Uints([]string{}))
	assert.Equal(t, []uint{0, 18, 2147483648}, Uints([]string{"0", "18", "2147483648"}))

	assert.Panics(t, func() {
		Uints([]string{"a"})
	})
}

func TestUint8s(t *testing.T) {
	assert.Equal(t, []uint8{}, Uint8s(nil))
	assert.Equal(t, []uint8{}, Uint8s([]string{}))
	assert.Equal(t, []uint8{0, 18, 255}, Uint8s([]string{"0", "18", "255"}))

	assert.Panics(t, func() {
		Uint8s([]string{"a"})
	})
}

func TestUint16s(t *testing.T) {
	assert.Equal(t, []uint16{}, Uint16s(nil))
	assert.Equal(t, []uint16{}, Uint16s([]string{}))
	assert.Equal(t, []uint16{0, 18, 65535}, Uint16s([]string{"0", "18", "65535"}))

	assert.Panics(t, func() {
		Uint16s([]string{"a"})
	})
}

func TestUint32s(t *testing.T) {
	assert.Equal(t, []uint32{}, Uint32s(nil))
	assert.Equal(t, []uint32{}, Uint32s([]string{}))
	assert.Equal(t, []uint32{0, 18, 4294967295}, Uint32s([]string{"0", "18", "4294967295"}))

	assert.Panics(t, func() {
		Uint32s([]string{"a"})
	})
}

func TestUint64s(t *testing.T) {
	assert.Equal(t, []uint64{}, Uint64s(nil))
	assert.Equal(t, []uint64{}, Uint64s([]string{}))
	assert.Equal(t, []uint64{0, 18, 18446744073709551615}, Uint64s([]string{"0", "18", "18446744073709551615"}))

	assert.Panics(t, func() {
		Uint64s([]string{"a"})
	})
}

func TestExtractInts(t *testing.T) {
	assert.Equal(t, []int{1, -2147483648, 2, 2147483647}, ExtractInts("1: -2147483648 \n 2 abc 2147483647\n"))
}

func TestExtractInt64s(t *testing.T) {
	assert.Equal(t, []int64{2, -9223372036854775808, 2, 9223372036854775807}, ExtractInt64s("2: -9223372036854775808 \n 2 abc 9223372036854775807\n"))
}

func TestExtractUints(t *testing.T) {
	assert.Equal(t, []uint{3, 0, 2, 4294967295}, ExtractUints("3: 0 \n 2 abc 4294967295\n"))
}

func TestExtractUint64s(t *testing.T) {
	assert.Equal(t, []uint64{4, 0, 2, 18446744073709551615}, ExtractUint64s("4: 0 \n 2 abc 18446744073709551615\n"))
}

func TestExtractDigits(t *testing.T) {
	assert.Equal(t, []int{1, 2, 3}, ExtractDigits("123"))
	assert.Equal(t, []int{1, 2, 3, 4, 5, 6, 7}, ExtractDigits("1\na23b456c7"))
}
