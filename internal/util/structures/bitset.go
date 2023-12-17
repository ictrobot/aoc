package structures

const (
	b32L1Shift = 6
	b32L1Bits  = 12
	b32L1Size  = 1 << b32L1Bits
	b32L1Mask  = b32L1Size - 1

	b32L2Shift = b32L1Shift + b32L1Bits
	b32L2Bits  = 7
	b32L2Size  = 1 << b32L2Bits
	b32L2Mask  = b32L2Size - 1

	b32L3Shift = b32L2Shift + b32L2Bits
	b32L3Bits  = 7
	b32L3Size  = 1 << b32L3Bits
	b32L3Mask  = b32L3Size - 1
)

type BitSet32[T ~int32 | ~uint32] [b32L2Size]*level2
type level2 [b32L2Size]*level1
type level1 [b32L1Size]uint64

func (b *BitSet32[T]) Has(t T) bool {
	x := uint32(t)

	lvl2 := b[(x>>b32L3Shift)&b32L3Mask]
	if lvl2 == nil {
		return false
	}

	lvl1 := lvl2[(x>>b32L2Shift)&b32L2Mask]
	if lvl1 == nil {
		return false
	}

	return (lvl1[(x>>b32L1Shift)&b32L1Mask] & (1 << (x & 63))) != 0
}

func (b *BitSet32[T]) Set(t T) {
	x := uint32(t)

	lvl2 := b[(x>>b32L3Shift)&b32L3Mask]
	if lvl2 == nil {
		lvl2 = new(level2)
		b[(x>>b32L3Shift)&b32L3Mask] = lvl2
	}

	lvl1 := lvl2[(x>>b32L2Shift)&b32L2Mask]
	if lvl1 == nil {
		lvl1 = new(level1)
		lvl2[(x>>b32L2Shift)&b32L2Mask] = lvl1
	}

	lvl1[(x>>b32L1Shift)&b32L1Mask] |= 1 << (x & 63)
}

func (b *BitSet32[T]) Clear(t T) {
	x := uint32(t)

	lvl2 := b[(x>>b32L3Shift)&b32L3Mask]
	if lvl2 == nil {
		return
	}

	lvl1 := lvl2[(x>>b32L2Shift)&b32L2Mask]
	if lvl1 == nil {
		return
	}

	lvl1[(x>>b32L1Shift)&b32L1Mask] &= ^(1 << (x & 63))
}

func (b *BitSet32[T]) Clone() *BitSet32[T] {
	c := new(BitSet32[T])
	for i := 0; i < b32L3Size; i++ {
		if b[i] == nil {
			continue
		}

		c[i] = new(level2)

		for j := 0; j < b32L2Size; j++ {
			if b[i][j] == nil {
				continue
			}

			c[i][j] = new(level1)
			*c[i][j] = *b[i][j]
		}
	}
	return c
}
