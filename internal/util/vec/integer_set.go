package vec

import (
	"golang.org/x/exp/constraints"
)

type I2Set[T constraints.Integer] map[I2[T]]struct{}

type I3Set[T constraints.Integer] map[I3[T]]struct{}

// Add the given vector to the set, returning if the vector was not present before
func (s I2Set[T]) Add(v I2[T]) bool {
	if _, ok := s[v]; ok {
		return false
	}

	s[v] = struct{}{}
	return true
}

// Add the given vector to the set, returning if the vector was not present before
func (s I3Set[T]) Add(v I3[T]) bool {
	if _, ok := s[v]; ok {
		return false
	}

	s[v] = struct{}{}
	return true
}

func (s I2Set[T]) AddInts(x, y T) bool {
	return s.Add(I2[T]{x, y})
}

func (s I3Set[T]) AddInts(x, y, z T) bool {
	return s.Add(I3[T]{x, y, z})
}

// Remove the given vector from the set, returning if it was present
func (s I2Set[T]) Remove(v I2[T]) bool {
	if _, ok := s[v]; !ok {
		return false
	}

	delete(s, v)
	return true
}

// Remove the given vector from the set, returning if it was present
func (s I3Set[T]) Remove(v I3[T]) bool {
	if _, ok := s[v]; !ok {
		return false
	}

	delete(s, v)
	return true
}

func (s I2Set[T]) RemoveInts(x, y T) bool {
	return s.Remove(I2[T]{x, y})
}

func (s I3Set[T]) RemoveInts(x, y, z T) bool {
	return s.Remove(I3[T]{x, y, z})
}

func (s I2Set[T]) Contains(v I2[T]) bool {
	_, present := s[v]
	return present
}

func (s I3Set[T]) Contains(v I3[T]) bool {
	_, present := s[v]
	return present
}

func (s I2Set[T]) ContainsInts(x, y T) bool {
	return s.Contains(I2[T]{x, y})
}

func (s I3Set[T]) ContainsInts(x, y, z T) bool {
	return s.Contains(I3[T]{x, y, z})
}
