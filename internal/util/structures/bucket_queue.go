package structures

import (
	"fmt"
	"strconv"
	"strings"
)

// BucketQueue implements a priority queue using a doubly linked list where
// each node contains a list of items with a given priority. Lower priority
// value items are higher priority i.e. are returned by Pop first. Items with
// the same priority may be returned in any order.
type BucketQueue[T comparable] struct {
	head *bucketNode[T]
	tail *bucketNode[T]
}

type bucketNode[T comparable] struct {
	priority   int
	values     []T
	next, prev *bucketNode[T]
}

// Push adds the given item with the given priority to the queue
// Lower priority values are returned by Pop first
func (b *BucketQueue[T]) Push(t T, priority int) {
	if b.head == nil {
		b.head = &bucketNode[T]{
			priority: priority,
			values:   []T{t},
		}
		b.tail = b.head
		return
	}
	if priority < b.head.priority {
		newNode := &bucketNode[T]{
			priority: priority,
			values:   []T{t},
			next:     b.head,
		}
		b.head.prev = newNode
		b.head = newNode
		return
	}
	if priority > b.tail.priority {
		newNode := &bucketNode[T]{
			priority: priority,
			values:   []T{t},
			prev:     b.tail,
		}
		b.tail.next = newNode
		b.tail = newNode
		return
	}

	// priority must be between or equal to one of head.priority & tail.priority
	// iterate forwards/backwards from whichever is numerically closer
	if priority-b.head.priority <= b.tail.priority-priority {
		// iterate forwards from b.head
		current := b.head
		for current.next != nil && priority >= current.next.priority {
			current = current.next
		}

		// current.next can be nil when the queue only has one priority
		// and pushing an item of the same priority. However, in that case the
		// following if statement should always be triggered, so the code below
		// it doesn't need to handle this case
		if priority == current.priority {
			current.values = append(current.values, t)
			return
		}

		newNode := &bucketNode[T]{
			priority: priority,
			values:   []T{t},
			next:     current.next,
			prev:     current,
		}
		current.next.prev = newNode
		current.next = newNode
	} else {
		// iterate backwards from b.tail
		current := b.tail
		for current.prev != nil && priority <= current.prev.priority {
			current = current.prev
		}

		if priority == current.priority {
			current.values = append(current.values, t)
			return
		}

		newNode := &bucketNode[T]{
			priority: priority,
			values:   []T{t},
			next:     current,
			prev:     current.prev,
		}
		current.prev.next = newNode
		current.prev = newNode
	}
}

// Pop returns and removes an item with the lowest priority value
func (b *BucketQueue[T]) Pop() T {
	if b.head == nil {
		panic("empty queue")
	}

	length := len(b.head.values)
	if length == 1 {
		v := b.head.values[0]
		b.head = b.head.next
		if b.head == nil {
			b.tail = nil
		} else {
			b.head.prev = nil
		}
		return v
	}

	// taking last item from slice instead of first avoids reducing its
	// capacity which helps to reduce allocations
	var zero T
	v := b.head.values[length-1]
	b.head.values[length-1] = zero // prevent memory leaks
	b.head.values = b.head.values[:length-1]
	return v
}

// Peek returns (without removing) the next item which would be returned by Pop
func (b *BucketQueue[T]) Peek() T {
	if b.head == nil {
		panic("empty queue")
	}

	// should always return the same item that would be popped next
	length := len(b.head.values)
	return b.head.values[length-1]
}

// IsEmpty returns if the queue contains zero items
func (b *BucketQueue[T]) IsEmpty() bool {
	return b.head == nil
}

// ToSlice returns a slice of the items currently in the queue, sorted by
// priority. Intended for debugging only
func (b *BucketQueue[T]) ToSlice() []T {
	var s []T
	current := b.head
	for current != nil {
		s = append(s, current.values...)
		current = current.next
	}
	return s
}

// String returns a pretty printed string representing the queue
func (b *BucketQueue[T]) String() string {
	var s strings.Builder
	s.WriteByte('[')

	current := b.head
	for current != nil {
		priority := strconv.Itoa(current.priority)

		// iterate backwards to match Pop order
		for i := len(current.values) - 1; i >= 0; i-- {
			if s.Len() > 1 {
				s.WriteString(", ")
			}
			s.WriteString(priority)
			s.WriteByte(':')
			_, _ = fmt.Fprint(&s, current.values[i])
		}

		current = current.next
	}

	s.WriteByte(']')
	return s.String()
}
