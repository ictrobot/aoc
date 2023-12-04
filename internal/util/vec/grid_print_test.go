package vec

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGrid_PrettyPrint(t *testing.T) {
	g := &Grid[uint8]{}
	assert.Equal(t, `[empty]`, g.PrettyPrint(1, false, false))

	for i := 0; i < 5; i++ {
		g.SetInts(-500+i, -10, '@')
	}

	g = g.Resize(-501, -11, -495, -9)
	assert.Equal(t, ` -495, -9
.......
.@@@@@.
.......
-501, -11`, g.PrettyPrint(1, false, false))

	// wide enough for both labels at top, but not bottom
	g = g.Resize(-507, -11, -489, -9)
	assert.Equal(t, `-507, -9   -489, -9
...................
.......@@@@@.......
...................
-507, -11`, g.PrettyPrint(1, false, false))

	g = g.Resize(-508, -11, -489, -9)
	assert.Equal(t, `-508, -9    -489, -9
....................
........@@@@@.......
....................
-508, -11  -489, -11`, g.PrettyPrint(1, false, false))

	g.SetInts(-499, -10, '#')
	for i := 0; i < 4; i++ {
		g.SetInts(-500+i, -10+i, '/')
	}

	g = g.Resize(-508, -11, -489, -6)
	assert.Equal(t, `-508, -6    -489, -6
....................
.........../........
........../.........
........./..........
......../#@@@.......
....................
-508, -11  -489, -11`, g.PrettyPrint(1, false, false))
	assert.Equal(t, `-489, -6    -508, -6
....................
......../...........
........./..........
........../.........
.......@@@#/........
....................
-489, -11  -508, -11`, g.PrettyPrint(1, true, false))
	assert.Equal(t, `-508, -11  -489, -11
....................
......../#@@@.......
........./..........
........../.........
.........../........
....................
-508, -6    -489, -6`, g.PrettyPrint(1, false, true))
	assert.Equal(t, `-489, -11  -508, -11
....................
.......@@@#/........
........../.........
........./..........
......../...........
....................
-489, -6    -508, -6`, g.PrettyPrint(1, true, true))

	g2 := &Grid[int]{}
	g2.SetInts(1, 1, 1)
	g2.SetInts(2, 1, 2)
	g2.SetInts(2, 2, 20)
	g2.SetInts(3, 1, 3)
	g2.SetInts(3, 3, 30)
	g2 = g2.Resize(0, 0, 4, 4)

	assert.Equal(t, ` 4, 4
.....
...#.
..#..
.###.
.....
0, 0`, g2.PrettyPrint(0, false, false))
	assert.Equal(t, ` 4, 4
00000
000!0
00!00
01230
00000
0, 0`, g2.PrettyPrint(1, false, false))
	assert.Equal(t, `0, 4  4, 4
 0 0 0 0 0
 0 0 030 0
 0 020 0 0
 0 1 2 3 0
 0 0 0 0 0
0, 0  4, 0`, g2.PrettyPrint(2, false, false))
	assert.Equal(t, `0, 4       4, 4
  0  0  0  0  0
  0  0  0 30  0
  0  0 20  0  0
  0  1  2  3  0
  0  0  0  0  0
0, 0       4, 0`, g2.PrettyPrint(3, false, false))

	assert.Panics(t, func() {
		g2.PrettyPrint(-1, false, false)
	})
}
