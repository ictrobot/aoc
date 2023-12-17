package structures

// Heap implements a binary heap using generics. A LessThan function must be
// set when creating a Heap
type Heap[T comparable] struct {
	s        []T
	LessThan func(a, b T) bool
}

func (h *Heap[T]) Len() int {
	return len(h.s)
}

func (h *Heap[T]) IsEmpty() bool {
	return len(h.s) == 0
}

func (h *Heap[T]) Peek() T {
	if len(h.s) == 0 {
		panic("empty heap")
	}

	return h.s[0]
}

func (h *Heap[T]) Pop() T {
	if len(h.s) == 0 {
		panic("empty heap")
	}

	x := h.s[0]
	h.s[0] = h.s[len(h.s)-1]

	var zero T
	h.s[len(h.s)-1] = zero
	h.s = h.s[:len(h.s)-1]

	h.heapifyDown(len(h.s))
	return x
}

func (h *Heap[T]) Push(x T) {
	if h.LessThan == nil {
		panic("no comparator")
	}

	h.s = append(h.s, x)
	h.heapifyUp(len(h.s) - 1)
}

func (h *Heap[T]) heapifyUp(i int) {
	for {
		parent := (i - 1) / 2
		if i == parent || !h.LessThan(h.s[i], h.s[parent]) {
			break
		}

		h.s[i], h.s[parent] = h.s[parent], h.s[i]
		i = parent
	}
}

func (h *Heap[T]) heapifyDown(n int) {
	parent := 0
	for {
		left := 2*parent + 1
		right := 2*parent + 2
		if left >= n {
			break
		}

		swap := left
		if right < n && h.LessThan(h.s[right], h.s[left]) {
			swap = right
		}

		if !h.LessThan(h.s[swap], h.s[parent]) {
			break
		}

		h.s[parent], h.s[swap] = h.s[swap], h.s[parent]
		parent = swap
	}
}
