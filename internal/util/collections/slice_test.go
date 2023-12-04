package collections

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestClone2D(t *testing.T) {
	i1 := [][]int{{1, 2}, {3, 4, 5}, {6, 7, 8, 9}}
	i2 := Clone2D(i1)
	assert.Equal(t, i1, i2)

	i2[0][1] = 100
	i2[1][0] = 200
	i2[2][2] = 300

	assert.Equal(t, [][]int{{1, 2}, {3, 4, 5}, {6, 7, 8, 9}}, i1)
	assert.Equal(t, [][]int{{1, 100}, {200, 4, 5}, {6, 7, 300, 9}}, i2)
}

func TestFill(t *testing.T) {
	for _, i := range []int{0, 1, 2, 8, 16, 32, 64, 128} {
		s := make([]int64, i)
		Fill(s, int64(1+i))

		assert.Equal(t, i, len(s))
		for _, v := range s {
			assert.Equal(t, int64(1+i), v)
		}
	}
}
