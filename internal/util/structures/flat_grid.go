package structures

import "slices"

// FlatGrid implements a generic 2d grid with a flat backing slice, which is
// exported for fast access & modification. This is ideal for simple grids
// where the minimum coordinate is (0, 0) and automatic resizing is not needed.
// Methods are simple and should be inlined in almost all cases.
//
// vec.Grid offers more functionality (including automatic resizing & bounds
// checks) at the cost of more overhead for each access/modification.
type FlatGrid[T any] struct {
	S             []T
	Width, Height int
}

func NewFlatGrid[T any](width, height int) *FlatGrid[T] {
	return &FlatGrid[T]{S: make([]T, width*height), Width: width, Height: height}
}

func (f *FlatGrid[T]) Get(x, y int) T {
	return f.S[(y*f.Height)+x]
}

func (f *FlatGrid[T]) Set(x, y int, t T) {
	f.S[(y*f.Height)+x] = t
}

func (f *FlatGrid[T]) Index(x, y int) int {
	return (y * f.Height) + x
}

func (f *FlatGrid[T]) Coords(idx int) (x, y int) {
	return idx / f.Height, idx % f.Height
}

func (f *FlatGrid[T]) InBounds(x, y int) bool {
	return x >= 0 && y >= 0 && x < f.Width && y < f.Height
}

func (f *FlatGrid[T]) Clone() FlatGrid[T] {
	return FlatGrid[T]{S: slices.Clone(f.S), Width: f.Width, Height: f.Height}
}

func (f *FlatGrid[T]) Clear() {
	clear(f.S)
}

// Size is a helper function which returns width & height, meant to be used
// when creating a new blank grid with the same size as an existing one
//
//	structures.NewFlatGrid(existing.Size())
func (f *FlatGrid[T]) Size() (width, height int) {
	return f.Width, f.Height
}
