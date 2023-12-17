package structures

import "container/heap"

// PriorityQueue implements a basic generic PriorityQueue using Heap
type PriorityQueue[T HeapItem[T]] struct {
	h Heap[T]
}

func (q *PriorityQueue[T]) Push(t T) {
	heap.Push(&q.h, t)
}

func (q *PriorityQueue[T]) Pop() T {
	return heap.Pop(&q.h).(T)
}

func (q *PriorityQueue[T]) Len() int {
	return len(q.h)
}

func (q *PriorityQueue[T]) IsEmpty() bool {
	return len(q.h) == 0
}
