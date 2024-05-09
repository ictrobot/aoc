package vec

import (
	"github.com/ictrobot/aoc-go/internal/util/numbers"
	"golang.org/x/exp/constraints"
	"math"
)

type I2[T constraints.Integer] struct {
	X, Y T
}

type I3[T constraints.Integer] struct {
	X, Y, Z T
}

func (i I2[T]) Add(j I2[T]) I2[T] {
	return I2[T]{i.X + j.X, i.Y + j.Y}
}

func (i I3[T]) Add(j I3[T]) I3[T] {
	return I3[T]{i.X + j.X, i.Y + j.Y, i.Z + j.Z}
}

func (i I2[T]) Sub(j I2[T]) I2[T] {
	return I2[T]{i.X - j.X, i.Y - j.Y}
}

func (i I3[T]) Sub(j I3[T]) I3[T] {
	return I3[T]{i.X - j.X, i.Y - j.Y, i.Z - j.Z}
}

func (i I2[T]) Mul(c T) I2[T] {
	return I2[T]{i.X * c, i.Y * c}
}

func (i I3[T]) Mul(c T) I3[T] {
	return I3[T]{i.X * c, i.Y * c, i.Z * c}
}

func (i I2[T]) EuclideanDist2(j I2[T]) int64 {
	dX := int64(i.X) - int64(j.X)
	dY := int64(i.Y) - int64(j.Y)
	return (dX * dX) + (dY * dY)
}

func (i I3[T]) EuclideanDist2(j I3[T]) int64 {
	dX := int64(i.X) - int64(j.X)
	dY := int64(i.Y) - int64(j.Y)
	dZ := int64(i.Z) - int64(j.Z)
	return (dX * dX) + (dY * dY) + (dZ * dZ)
}

func (i I2[T]) EuclideanDist(j I2[T]) float64 {
	dX := float64(i.X) - float64(j.X)
	dY := float64(i.Y) - float64(j.Y)
	return math.Sqrt((dX * dX) + (dY * dY))
}

func (i I3[T]) EuclideanDist(j I3[T]) float64 {
	dX := float64(i.X) - float64(j.X)
	dY := float64(i.Y) - float64(j.Y)
	dZ := float64(i.Z) - float64(j.Z)
	return math.Sqrt((dX * dX) + (dY * dY) + (dZ * dZ))
}

func (i I2[T]) ManhattanDist(j I2[T]) int64 {
	dX := numbers.IntAbsDiff(int64(i.X), int64(j.X))
	dY := numbers.IntAbsDiff(int64(i.Y), int64(j.Y))
	return dX + dY
}

func (i I3[T]) ManhattanDist(j I3[T]) int64 {
	dX := numbers.IntAbsDiff(int64(i.X), int64(j.X))
	dY := numbers.IntAbsDiff(int64(i.Y), int64(j.Y))
	dZ := numbers.IntAbsDiff(int64(i.Z), int64(j.Z))
	return dX + dY + dZ
}

func (i I2[T]) MaxComponent() T {
	return max(i.X, i.Y)
}

func (i I3[T]) MaxComponent() T {
	return max(i.X, i.Y, i.Z)
}

func (i I2[T]) MinComponent() T {
	return min(i.X, i.Y)
}

func (i I3[T]) MinComponent() T {
	return min(i.X, i.Y, i.Z)
}

func (i I2[T]) Abs() I2[T] {
	return I2[T]{numbers.IntAbs(i.X), numbers.IntAbs(i.Y)}
}

func (i I3[T]) Abs() I3[T] {
	return I3[T]{numbers.IntAbs(i.X), numbers.IntAbs(i.Y), numbers.IntAbs(i.Z)}
}

func (i I2[T]) Sign() I2[int] {
	return I2[int]{numbers.IntSign(i.X), numbers.IntSign(i.Y)}
}

func (i I3[T]) Sign() I3[int] {
	return I3[int]{numbers.IntSign(i.X), numbers.IntSign(i.Y), numbers.IntSign(i.Z)}
}

func (i I2[T]) WithZ(z T) I3[T] {
	return I3[T]{i.X, i.Y, z}
}

func (i I3[T]) XY() I2[T] {
	return I2[T]{i.X, i.Y}
}
