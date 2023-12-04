package parse

import (
	"encoding/json"
	"errors"
	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
	"io"
	"strconv"
	"strings"
	"testing"
)

type reflectTest struct {
	Bool bool
	Ints struct {
		I   int
		I8  int8
		I16 int16
		I32 int32
		I64 int64
	}
	Uints struct {
		U   uint
		U8  uint8
		U16 uint16
		U32 uint32
		U64 uint64
	}
	Floats struct {
		F float32
		D float64
	}
	Str      string
	Optional *int
	Present  *string
	Arr      [3]string
	Coords   []struct{ X, Y int }
}

var reflectTestStrings = []string{
	"true",
	"-1",
	"-2",
	"-4",
	"-8",
	"-16",
	"0",
	"10",
	"100",
	"1000",
	"10000",
	"1234.5",
	"9876.5",
	"Hello World",
	// Optional skipped
	"Testing123",
	"X", "Y", "Z",
	"0", "1",
	"-1", "5",
	"1", "3",
}

var reflectTestExpected = reflectTest{
	Bool: true,
	Ints: struct {
		I   int
		I8  int8
		I16 int16
		I32 int32
		I64 int64
	}{-1, -2, -4, -8, -16},
	Uints: struct {
		U   uint
		U8  uint8
		U16 uint16
		U32 uint32
		U64 uint64
	}{0, 10, 100, 1000, 10000},
	Floats: struct {
		F float32
		D float64
	}{1234.5, 9876.5},
	Str:      "Hello World",
	Optional: nil,
	Present:  lo.ToPtr("Testing123"),
	Arr:      [...]string{"X", "Y", "Z"},
	Coords: []struct{ X, Y int }{
		{0, 1},
		{-1, 5},
		{1, 3},
	},
}

type reflectPtrTest struct {
	Coords *[2]int
	Final  int
}

type reflectUnexportedTest struct {
	unexported int
}

type reflectSliceTest struct {
	Opcode   string
	Operands []int
	Extra    *string
}

func TestReflect(t *testing.T) {
	// Int
	b0, err := Reflect[bool]([]string{"1"})
	assert.Equal(t, true, b0)
	assert.NoError(t, err)
	b1, err := Reflect[bool]([]string{"false"})
	assert.Equal(t, false, b1)
	assert.NoError(t, err)
	b2, err := Reflect[bool]([]string{})
	assert.Zero(t, b2)
	assert.ErrorIs(t, err, NotEnoughStrings)
	b3, err := Reflect[bool]([]string{"a"})
	assert.Zero(t, b3)
	assert.ErrorIs(t, err, UnknownValueError)

	// Int
	i0, err := Reflect[int]([]string{"1"})
	assert.Equal(t, 1, i0)
	assert.NoError(t, err)
	i1, err := Reflect[int]([]string{})
	assert.Zero(t, i1)
	assert.ErrorIs(t, err, NotEnoughStrings)
	i2, err := Reflect[int]([]string{"a"})
	assert.Zero(t, i2)
	assert.ErrorIs(t, err, strconv.ErrSyntax)
	i3, err := Reflect[int8]([]string{"128"})
	assert.Zero(t, i3)
	assert.ErrorIs(t, err, strconv.ErrRange)

	// Uint
	u0, err := Reflect[uint]([]string{"1"})
	assert.Equal(t, uint(1), u0)
	assert.NoError(t, err)
	u1, err := Reflect[uint]([]string{})
	assert.Zero(t, u1)
	assert.ErrorIs(t, err, NotEnoughStrings)
	u2, err := Reflect[uint]([]string{"-1"})
	assert.Zero(t, u2)
	assert.ErrorIs(t, err, strconv.ErrSyntax)
	u3, err := Reflect[uint8]([]string{"256"})
	assert.Zero(t, u3)
	assert.ErrorIs(t, err, strconv.ErrRange)

	// Float
	f0, err := Reflect[float32]([]string{"1.125"})
	assert.Equal(t, float32(1.125), f0)
	assert.NoError(t, err)
	f1, err := Reflect[float64]([]string{"12.875"})
	assert.Equal(t, 12.875, f1)
	assert.NoError(t, err)
	f2, err := Reflect[float64]([]string{})
	assert.Zero(t, f2)
	assert.ErrorIs(t, err, NotEnoughStrings)
	f3, err := Reflect[float64]([]string{"2E+308"})
	assert.Zero(t, f3)
	assert.ErrorIs(t, err, strconv.ErrRange)

	// Array
	a0, err := Reflect[[2]uint]([]string{"1", "2"})
	assert.Equal(t, [2]uint{1, 2}, a0)
	assert.NoError(t, err)
	a1, err := Reflect[[2]int]([]string{"1"})
	assert.Zero(t, a1)
	assert.ErrorIs(t, err, NotEnoughStrings)

	// Pointer
	p0, err := Reflect[*int]([]string{"123"})
	assert.Equal(t, 123, *p0)
	assert.NoError(t, err)
	p1, err := Reflect[*int]([]string{})
	assert.Equal(t, (*int)(nil), p1)
	assert.NoError(t, err)
	p2, err := Reflect[reflectPtrTest]([]string{"10", "20", "30"})
	assert.Equal(t, reflectPtrTest{lo.ToPtr([2]int{10, 20}), 30}, p2)
	assert.NoError(t, err)
	p3, err := Reflect[reflectPtrTest]([]string{"10"})
	assert.Equal(t, reflectPtrTest{nil, 10}, p3)
	assert.NoError(t, err)
	p4, err := Reflect[reflectPtrTest]([]string{})
	assert.Zero(t, p4)
	assert.ErrorIs(t, err, NotEnoughStrings)
	p5, err := Reflect[reflectPtrTest]([]string{"10", "20"})
	assert.Zero(t, p5)
	assert.ErrorIs(t, err, NotEnoughStrings)
	p6, err := Reflect[reflectPtrTest]([]string{"10", "20", "30", "40"})
	assert.Zero(t, p6)
	assert.ErrorIs(t, err, TooManyStrings)
	p7, err := Reflect[*int]([]string{"a"})
	assert.Zero(t, p7)
	assert.ErrorIs(t, err, TooManyStrings)

	// Slice
	s1, err := Reflect[[]int]([]string{"1", "2", "3"})
	assert.Equal(t, []int{1, 2, 3}, s1)
	assert.Zero(t, err)
	s2, err := Reflect[[][2]int]([]string{"4", "5", "6", "7"})
	assert.Equal(t, [][2]int{{4, 5}, {6, 7}}, s2)
	assert.Zero(t, err)
	s3, err := Reflect[[][2]int]([]string{"4", "5", "a", "7"})
	assert.Zero(t, s3)
	assert.ErrorIs(t, err, TooManyStrings)
	s4, err := Reflect[[]*int]([]string{"a"})
	assert.Zero(t, s4)
	assert.ErrorIs(t, err, TooManyStrings)

	// String
	t1, err := Reflect[string]([]string{"ABC"})
	assert.Equal(t, "ABC", t1)
	assert.Zero(t, err)
	t2, err := Reflect[string]([]string{})
	assert.Zero(t, t2)
	assert.ErrorIs(t, err, NotEnoughStrings)

	// Struct
	x1, err := Reflect[reflectTest](reflectTestStrings)
	assert.Equal(t, reflectTestExpected, x1)
	assert.NoError(t, err)
	x2, err := Reflect[reflectSliceTest]([]string{"add", "1"})
	assert.Equal(t, reflectSliceTest{Opcode: "add", Operands: []int{1}}, x2)
	assert.NoError(t, err)
	x3, err := Reflect[reflectSliceTest]([]string{"add", "1", "2", "3"})
	assert.Equal(t, reflectSliceTest{Opcode: "add", Operands: []int{1, 2, 3}}, x3)
	assert.NoError(t, err)
	x4, err := Reflect[reflectSliceTest]([]string{"add", "1", "2", "3", "x"})
	assert.Equal(t, reflectSliceTest{Opcode: "add", Operands: []int{1, 2, 3}, Extra: lo.ToPtr("x")}, x4)
	assert.NoError(t, err)
	x5, err := Reflect[reflectSliceTest]([]string{"add", "y"})
	assert.Equal(t, reflectSliceTest{Opcode: "add", Extra: lo.ToPtr("y")}, x5)
	assert.NoError(t, err)
	x6, err := Reflect[reflectSliceTest]([]string{"add", "x", "y"})
	assert.Zero(t, x6)
	assert.ErrorIs(t, err, TooManyStrings)

	assert.PanicsWithValue(t, "unexported field: field 0 `unexported` in `reflectUnexportedTest`", func() {
		_, _ = Reflect[reflectUnexportedTest]([]string{"1"})
	})

	// Since slices greedily consume strings, there will never be an int left for Never
	x7, err := Reflect[struct {
		Ints  []int
		Never int
	}]([]string{"1", "2", "3"})
	assert.Zero(t, x7)
	assert.ErrorIs(t, err, NotEnoughStrings)

	// Unsupported
	assert.PanicsWithValue(t, "unsupported type: chan string", func() {
		_, _ = Reflect[chan string]([]string{""})
	})
}

func TestMustReflect(t *testing.T) {
	actual := MustReflect[reflectTest](reflectTestStrings)
	assert.Equal(t, reflectTestExpected, actual)

	assert.Panics(t, func() {
		MustReflect[struct{ X int }]([]string{})
	})
}

func BenchmarkMustReflect(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MustReflect[reflectTest](reflectTestStrings)
	}
}

// real code should use json.Unmarshal !!!
type nestedArray struct {
	_        Placeholder `match:"["`
	Elements []nestedElement
	_        Placeholder `match:"]"`
	_        Placeholder `regex:"\\s" flags:"optional,multiple"`
}

type nestedElement struct {
	Int    *int
	Nested *nestedArray
	_      Placeholder `match:"," flags:"optional"`
	_      Placeholder `regex:"\\s" flags:"optional,multiple"`
}

var nestedArrays = "[[1], [2,3,4]]\n\n[1, [2, [3, [4, 5]]], 6, 7]\n"

func TestMustReflect_NestedArray(t *testing.T) {
	assert.Equal(t, []nestedArray{
		{ // [[1], [2,3,4]]
			Elements: []nestedElement{
				{ // [1]
					Nested: &nestedArray{
						Elements: []nestedElement{
							{Int: lo.ToPtr(1)},
						},
					},
				},
				{ // [2, 3, 4]
					Nested: &nestedArray{
						Elements: []nestedElement{
							{Int: lo.ToPtr(2)},
							{Int: lo.ToPtr(3)},
							{Int: lo.ToPtr(4)},
						},
					},
				},
			},
		},
		{ // [1, [2, [3, [4, 5]]], 6, 7]
			Elements: []nestedElement{
				{
					Int: lo.ToPtr(1),
				},
				{ // [2, [3, [4, 5]]]
					Nested: &nestedArray{
						Elements: []nestedElement{
							{Int: lo.ToPtr(2)},
							{ // [3, [4, 5]]
								Nested: &nestedArray{
									Elements: []nestedElement{
										{Int: lo.ToPtr(3)},
										{ // [4, 5]
											Nested: &nestedArray{
												Elements: []nestedElement{
													{Int: lo.ToPtr(4)},
													{Int: lo.ToPtr(5)},
												},
											},
										},
									},
								},
							},
						},
					},
				},
				{
					Int: lo.ToPtr(6),
				},
				{
					Int: lo.ToPtr(7),
				},
			},
		},
	}, MustReflect[[]nestedArray](Characters(nestedArrays)))
}

func BenchmarkMustReflect_NestedArray(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MustReflect[[]nestedArray](Characters(nestedArrays))
	}
}

func BenchmarkJsonDecoder_NestedArray(b *testing.B) {
	for i := 0; i < b.N; i++ {
		dec := json.NewDecoder(strings.NewReader(nestedArrays))

		for {
			var arr []interface{}
			if err := dec.Decode(&arr); errors.Is(err, io.EOF) {
				break
			} else if err != nil {
				b.Fatal(err)
			}
		}
	}
}
