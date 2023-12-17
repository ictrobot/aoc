package structures

// HeapItem represents an item which can be placed into a Heap or PriorityQueue
type HeapItem[T any] interface {
	comparable

	// LessThan returns if the current element is less than the argument.
	// See heap.Interface for requirements
	LessThan(T) bool
}

// Heap implements heap.Interface and is used by PriorityQueue
type Heap[T HeapItem[T]] []T

func (h *Heap[T]) Len() int {
	return len(*h)
}

func (h *Heap[T]) Less(i, j int) bool {
	return (*h)[i].LessThan((*h)[j])
}

func (h *Heap[T]) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *Heap[T]) Push(x any) {
	*h = append(*h, x.(T))
}

func (h *Heap[T]) Pop() any {
	s := *h
	*h = s[0 : len(s)-1]
	return s[len(s)-1]
}
