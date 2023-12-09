package vec

import (
	"fmt"
	"math"
)

type Grid[T comparable] struct {
	init       bool
	xMin, yMin int
	xMax, yMax int
	s          [][]T
}

// NewGrid returns a Grid initially sized to hold xMin <= x <= xMax
// & yMin <= y <= yMax. It will automatically grow as needed
func NewGrid[T comparable](xMin, yMin, xMax, yMax int) *Grid[T] {
	return (&Grid[T]{}).Resize(xMin, yMin, xMax, yMax)
}

// Clone returns a copy of this Grid. Elements are shallow copied
func (g *Grid[T]) Clone() *Grid[T] {
	if !g.init {
		return &Grid[T]{}
	}
	return g.Resize(g.xMin, g.yMin, g.xMax, g.yMax)
}

// Get returns the element at the given coords
func (g *Grid[T]) Get(v I2[int]) (t T) {
	if g.inBounds(v) {
		t = g.s[v.X-g.xMin][v.Y-g.yMin]
	}
	return
}

// GetInts returns the element at the given coords
// Provided for convenience when not using I2
func (g *Grid[T]) GetInts(x, y int) T {
	return g.Get(I2[int]{x, y})
}

// Set the element at the given coords, returning the old element
func (g *Grid[T]) Set(v I2[int], t T) T {
	if !g.inBounds(v) {
		var zero T
		if t == zero {
			// avoid resizing grid when setting element outside bounds to zero
			return zero
		}
		g.resizeToInclude(v.X, v.Y)
	}

	old := g.s[v.X-g.xMin][v.Y-g.yMin]
	g.s[v.X-g.xMin][v.Y-g.yMin] = t
	return old
}

// SetInts sets the element at the given coords, returning the old element
// Provided for convenience when not using I2
func (g *Grid[T]) SetInts(x, y int, t T) T {
	return g.Set(I2[int]{x, y}, t)
}

// SetIfZero sets the element at the given coords if it is currently zero,
// returning whether the element was zero
func (g *Grid[T]) SetIfZero(v I2[int], t T) bool {
	var zero T
	if !g.inBounds(v) {
		if t == zero {
			// avoid resizing grid when setting element outside bounds to zero
			return false
		}
		g.resizeToInclude(v.X, v.Y)
	}

	if g.s[v.X-g.xMin][v.Y-g.yMin] != zero {
		return false
	}

	g.s[v.X-g.xMin][v.Y-g.yMin] = t
	return true
}

// SetIfZeroInts sets the element at the given coords if it is currently zero,
// returning whether the element was zero
// Provided for convenience when not using I2
func (g *Grid[T]) SetIfZeroInts(x, y int, t T) bool {
	return g.SetIfZero(I2[int]{x, y}, t)
}

// Contains returns whether the element at the given coords is non-zero
func (g *Grid[T]) Contains(v I2[int]) bool {
	var zero T
	return g.inBounds(v) && g.s[v.X-g.xMin][v.Y-g.yMin] != zero
}

// ContainsInts returns whether the element at the given coords is non-zero
// Provided for convenience when not using I2
func (g *Grid[T]) ContainsInts(x, y int) bool {
	return g.Contains(I2[int]{x, y})
}

// Count returns the count of a given element in the grid, compared using ==
func (g *Grid[T]) Count(t T) int {
	if !g.init {
		return 0
	}

	count := 0
	for x := 0; x < len(g.s); x++ {
		for y := 0; y < len(g.s[x]); y++ {
			if g.s[x][y] == t {
				count++
			}
		}
	}
	return count
}

// CountNotZero returns the number of non-zero elements in the grid, compared using ==
func (g *Grid[T]) CountNotZero() int {
	if !g.init {
		return 0
	}

	var zero T
	count := 0
	for x := 0; x < len(g.s); x++ {
		for y := 0; y < len(g.s[x]); y++ {
			if g.s[x][y] != zero {
				count++
			}
		}
	}
	return count
}

// Counts returns the number of each distinct value of T in the grid
func (g *Grid[T]) Counts() map[T]int {
	if !g.init {
		return nil
	}

	var zero T
	counts := make(map[T]int)
	for x := 0; x < len(g.s); x++ {
		for y := 0; y < len(g.s[x]); y++ {
			if g.s[x][y] != zero {
				counts[g.s[x][y]]++
			}
		}
	}
	return counts
}

// Bounds returns the grid's xMin, yMin, xMax, yMax, which is likely to be
// slightly larger than the actual bounds containing non-zero values.
//
// Can be used with NewGrid to make a new grid with the same size as this grid:
//
//	vec.NewGrid[T](grid.Bounds())
func (g *Grid[T]) Bounds() (xMin int, yMin int, xMax int, yMax int) {
	return g.xMin, g.yMin, g.xMax, g.yMax
}

// NonZeroBounds returns the min x, min y, max x & max y with non-zero elements.
// If there are no non-zero elements, zeroes are returned.
func (g *Grid[T]) NonZeroBounds() (xMin int, yMin int, xMax int, yMax int) {
	if !g.init {
		return 0, 0, 0, 0
	}

	xMin = math.MaxInt
	yMin = math.MaxInt
	xMax = math.MinInt
	yMax = math.MinInt

	var zero T
	found := false
	for x := 0; x < len(g.s); x++ {
		for y := 0; y < len(g.s[x]); y++ {
			if g.s[x][y] != zero {
				found = true
				xMin = min(xMin, x)
				yMin = min(yMin, y)
				xMax = max(xMax, x)
				yMax = max(yMax, y)
			}
		}
	}

	if !found {
		return 0, 0, 0, 0
	}

	return xMin + g.xMin, yMin + g.yMin, xMax + g.xMin, yMax + g.yMin
}

// Resize returns a copy of the Grid resized to the specified size.
// Any elements outside the new size will be lost
func (g *Grid[T]) Resize(xMin, yMin, xMax, yMax int) *Grid[T] {
	if xMin > xMax {
		panic(fmt.Errorf("x min %d larger than x max %d", xMin, xMax))
	} else if yMin > yMax {
		panic(fmt.Errorf("y min %d larger than y max %d", xMin, xMax))
	}

	n := Grid[T]{init: true, xMin: xMin, yMin: yMin, xMax: xMax, yMax: yMax}

	xLen := xMax - xMin + 1
	n.s = make([][]T, xLen)

	// allocate each slice from one large slice to ensure slices are contiguous
	// which slightly improves both allocation & access performance
	yLen := yMax - yMin + 1
	ts := make([]T, xLen*yLen)
	for x := 0; x < xMax-xMin+1; x++ {
		n.s[x] = ts[x*yLen : (x+1)*yLen : (x+1)*yLen]
	}

	if g.init {
		// copy region in bounds on both
		x1 := max(xMin, g.xMin)
		x2 := min(xMax, g.xMax)
		y1 := max(yMin, g.yMin)
		y2 := min(yMax, g.yMax)
		for x := x1; x <= x2; x++ {
			copy(n.s[x-xMin][y1-yMin:y2-yMin+1], g.s[x-g.xMin][y1-g.yMin:y2-g.yMin+1])
		}
	}

	return &n
}

func (g *Grid[T]) inBounds(v I2[int]) bool {
	return g.init &&
		v.X >= g.xMin &&
		v.X <= g.xMax &&
		v.Y >= g.yMin &&
		v.Y <= g.yMax
}

func (g *Grid[T]) resizeToInclude(x, y int) {
	if g.init {
		*g = *g.Resize(
			min(g.xMin, x-8),
			min(g.yMin, y-8),
			max(g.xMax, x+8),
			max(g.yMax, y+8),
		)
	} else {
		*g = *g.Resize(x-8, y-8, x+8, y+8)
	}
}

func (g *Grid[T]) Format(f fmt.State, _ rune) {
	// ensure the grid is printed when printing parsed input
	// only works if the grid is stored in an exported field
	_, _ = f.Write([]byte{'&'})
	if f.Flag('+') {
		_, _ = fmt.Fprintf(f, "%+v", *g)
	} else if f.Flag('#') {
		_, _ = fmt.Fprintf(f, "%#v", *g)
	} else {
		_, _ = fmt.Fprintf(f, "%v", *g)
	}
}
