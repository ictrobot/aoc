package structures

// Deque implements a double ended queue ("deque") using a circular buffer.
// This enables it to be used as a queue or stack, but is mainly beneficial
// for queue usage as it avoids reducing the slice's capacity when popping
// elements from the front
type Deque[T any] struct {
	s          []T
	head, tail int
	len        int
}

// PushFront adds an element at the front
func (d *Deque[T]) PushFront(t T) {
	if len(d.s) == d.len {
		d.grow()
	}

	if d.len > 0 {
		if d.head == 0 {
			d.head = len(d.s) - 1
		} else {
			d.head--
		}
	}

	d.s[d.head] = t
	d.len++
}

// PushBack adds an element at the end
func (d *Deque[T]) PushBack(t T) {
	if len(d.s) == d.len {
		d.grow()
	}

	if d.len > 0 {
		if d.tail == len(d.s)-1 {
			d.tail = 0
		} else {
			d.tail++
		}
	}

	d.s[d.tail] = t
	d.len++
}

// PopFront removes the first element
func (d *Deque[T]) PopFront() T {
	if d.len == 0 {
		panic("deque empty")
	}

	var zero T
	t := d.s[d.head]
	d.s[d.head] = zero

	if d.len > 1 {
		if d.head == len(d.s)-1 {
			d.head = 0
		} else {
			d.head++
		}
	}

	d.len--
	return t
}

// PopBack removes the last element
func (d *Deque[T]) PopBack() T {
	if d.len == 0 {
		panic("deque empty")
	}

	var zero T
	t := d.s[d.tail]
	d.s[d.tail] = zero

	if d.len > 1 {
		if d.tail == 0 {
			d.tail = len(d.s) - 1
		} else {
			d.tail--
		}
	}

	d.len--
	return t
}

// PeekFront returns the first element without removing it
func (d *Deque[T]) PeekFront() T {
	if d.IsEmpty() {
		panic("deque empty")
	}

	return d.s[d.head]
}

// PeekBack returns the last element without removing it
func (d *Deque[T]) PeekBack() T {
	if d.IsEmpty() {
		panic("deque empty")
	}

	return d.s[d.tail]
}

func (d *Deque[T]) Len() int {
	return d.len
}

func (d *Deque[T]) IsEmpty() bool {
	return d.len == 0
}

func (d *Deque[T]) grow() {
	if len(d.s) == 0 {
		d.s = make([]T, 32)
		return
	}

	s := make([]T, len(d.s)*2)
	if d.head <= d.tail {
		copy(s, d.s[d.head:d.tail+1])
	} else {
		copy(s, d.s[d.head:])
		copy(s[len(d.s)-d.head:], d.s[:d.tail+1])
	}
	d.s = s

	d.head = 0
	d.tail = d.len - 1
}
