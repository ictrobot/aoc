package deep

import (
	"github.com/ictrobot/aoc/internal/util/parse"
	"github.com/stretchr/testify/assert"
	"testing"
)

type cloneTest struct {
	Name string
	Arr  [3]int
	Self *cloneTest
}

type placeholderTest struct {
	X, Y int
	_    parse.Placeholder
	_    struct{}
}

func TestClone(t *testing.T) {
	assert.Nil(t, Clone[int](nil))

	// Array
	arrIn := [3]int{1, 2, 3}
	arrOut := *Clone(&arrIn)
	arrOut[2] = 10
	assert.Equal(t, [3]int{1, 2, 3}, arrIn)
	assert.Equal(t, [3]int{1, 2, 10}, arrOut)

	// Map
	mapIn := map[int]string{
		1: "one",
		2: "two",
		3: "three",
	}
	mapOut := *Clone(&mapIn)
	delete(mapOut, 2)
	mapOut[4] = "four"
	assert.Equal(t, map[int]string{
		1: "one",
		2: "two",
		3: "three",
	}, mapIn)
	assert.Equal(t, map[int]string{
		1: "one",
		3: "three",
		4: "four",
	}, mapOut)

	// Pointer
	target := 123
	ptrIn := &target
	ptrOut := Clone(ptrIn)
	assert.Equal(t, 123, *ptrIn)
	assert.Equal(t, 123, *ptrOut)
	*ptrOut = 456
	assert.Equal(t, 123, *ptrIn)
	assert.Equal(t, 456, *ptrOut)

	// Slice
	sliceIn := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	sliceOut := *Clone(&sliceIn)
	sliceOut[0][2] = 9
	sliceOut[2] = []int{10, 20, 30}
	assert.Equal(t, [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, sliceIn)
	assert.Equal(t, [][]int{{1, 2, 9}, {4, 5, 6}, {10, 20, 30}}, sliceOut)

	// Struct
	structIn := cloneTest{Name: "TestValue", Arr: [3]int{-1, 0, 1}}
	structOut := *Clone(&structIn)
	structOut.Name = "Cloned"
	structOut.Arr[2] *= 100
	assert.Equal(t, cloneTest{Name: "TestValue", Arr: [3]int{-1, 0, 1}}, structIn)
	assert.Equal(t, cloneTest{Name: "Cloned", Arr: [3]int{-1, 0, 100}}, structOut)

	// Shared pointers
	sharedIn := struct{ A, B *int }{&target, &target}
	sharedOut := *Clone(&sharedIn)
	*sharedOut.A = 100
	assert.Equal(t, 123, *sharedIn.A)
	assert.Equal(t, 123, *sharedIn.B)
	assert.Equal(t, 100, *sharedOut.A)
	assert.Equal(t, 100, *sharedOut.B)

	// Struct with recursive pointer
	recursiveIn := &cloneTest{Name: "Recursive", Arr: [3]int{8, 4, 2}}
	recursiveIn.Self = recursiveIn
	recursiveOut := Clone(recursiveIn)
	recursiveOut.Name = "Clone"
	recursiveOut.Self.Self.Arr[1] = 0

	recursiveInExpected := &cloneTest{Name: "Recursive", Arr: [3]int{8, 4, 2}}
	recursiveInExpected.Self = recursiveInExpected
	recursiveOutExpected := &cloneTest{Name: "Clone", Arr: [3]int{8, 0, 2}}
	recursiveOutExpected.Self = recursiveOutExpected
	assert.Equal(t, recursiveInExpected, recursiveIn)
	assert.Equal(t, recursiveOutExpected, recursiveOut)

	// Struct with placeholders
	placeholderIn := placeholderTest{X: 12, Y: 23}
	placeholderOut := *Clone(&placeholderIn)
	placeholderOut.Y = 100
	assert.Equal(t, placeholderTest{X: 12, Y: 23}, placeholderIn)
	assert.Equal(t, placeholderTest{X: 12, Y: 100}, placeholderOut)

	// Panics
	assert.PanicsWithValue(t, "unexported field: field 0 `a` in ``", func() {
		Clone[struct{ a int }](&struct{ a int }{10})
	})
	assert.PanicsWithValue(t, "unsupported type: func(string)", func() {
		var x func(string)
		Clone(&x)
	})
}
