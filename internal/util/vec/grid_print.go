package vec

import (
	"fmt"
	"reflect"
	"strings"
)

// PrettyPrint returns a pretty printed string representing the grid
//
// Each element will be padded to valueWidth characters wide. If the value is
// already wider than this, it will be truncated and the last character
// replaced with !. Setting width to zero will just print non-zero elements,
// and setting width to one on a Grid[uint8] will automatically interpret the
// bytes as characters (any characters < 32 will be replaced to avoid printing
// control characters)
//
// By default, the bottom left is (minX, minY)
func (g *Grid[T]) PrettyPrint(valueWidth int, flipX, flipY bool) string {
	if valueWidth < 0 {
		panic(fmt.Sprintf("invalid value width: %d", valueWidth))
	}

	if !g.init || g.xMin == g.xMax || g.yMin == g.yMax {
		return "[empty]"
	}

	var zero T

	printAsChar := false
	if valueWidth == 1 {
		printAsChar = reflect.ValueOf(zero).Kind() == reflect.Uint8
	}

	// for corner labels
	var x1, x2, y1, y2 int
	if flipX {
		x1 = g.xMax
		x2 = g.xMin
	} else {
		x1 = g.xMin
		x2 = g.xMax
	}
	if flipY {
		y1 = g.yMin
		y2 = g.yMax
	} else {
		y1 = g.yMax
		y2 = g.yMin
	}

	var b strings.Builder
	actualWidth := valueWidth
	if actualWidth < 1 {
		actualWidth = 1
	}
	gridWidth := actualWidth * (g.xMax - g.xMin + 1)

	gridLabel(&b, true, x1, y1, x2, y1, gridWidth)
	b.WriteByte('\n')

	for i := 0; i <= g.yMax-g.yMin; i++ {
		y := g.yMax - i
		if flipY {
			y = g.yMin + i
		}

		for j := 0; j <= g.xMax-g.xMin; j++ {
			x := g.xMin + j
			if flipX {
				x = g.xMax - j
			}

			v := g.s[x-g.xMin][y-g.yMin]
			if valueWidth == 0 {
				// Print if non-zero
				if v != zero {
					b.WriteByte('#')
				} else {
					b.WriteByte('.')
				}
				continue
			}

			if printAsChar {
				// previously checked type of zero T is uint8, so should be safe
				u := uint8(reflect.ValueOf(v).Uint())
				if u < ' ' {
					b.WriteByte('.')
				} else {
					b.WriteByte(u)
				}
				continue
			}

			str := fmt.Sprintf("%v", v)

			padding := valueWidth - len(str)
			if padding < 0 {
				// value too big, cut at width and replace last char
				str = str[:valueWidth-1] + "!"
			} else if padding > 0 {
				for p := 0; p < padding; p++ {
					b.WriteByte(' ')
				}
			}

			b.WriteString(str)
		}
		b.WriteByte('\n')
	}

	gridLabel(&b, false, x1, y2, x2, y2, gridWidth)

	return b.String()
}

func gridLabel(b *strings.Builder, preferRight bool, x1, y1, x2, y2, width int) {
	left := fmt.Sprintf("%d, %d", x1, y1)
	right := fmt.Sprintf("%d, %d", x2, y2)

	if len(left)+len(right)+2 > width {
		// require two spaces between labels
		if preferRight {
			left = ""
		} else {
			// don't print padding if only printing left
			b.WriteString(left)
			return
		}
	}

	b.WriteString(left)

	// ensure at least one space is printed to show that label is right label,
	// especially when only printing right
	for i := 0; i < max(1, width-len(left)-len(right)); i++ {
		b.WriteByte(' ')
	}

	b.WriteString(right)
}
