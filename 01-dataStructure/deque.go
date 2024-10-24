type Deque[T any] struct {
	l, r []T
}

func NewDeque[T any]() *Deque[T] {
	return &Deque[T]{}
}

func (q *Deque[T]) Empty() bool {
	return len(q.l) == 0 && len(q.r) == 0
}

func (q *Deque[T]) Size() int {
	return len(q.l) + len(q.r)
}

func (q *Deque[T]) PushFront(v T) {
	q.l = append(q.l, v)
}

func (q *Deque[T]) PushBack(v T) {
	q.r = append(q.r, v)
}

func (q *Deque[T]) PopFront() T {
	if q.Empty() {
		panic("cannot pop from front: deque is Empty")
	}
	if len(q.l) > 0 {
		v, q.l := q.l[len(q.l)-1], q.l[:len(q.l)-1]
		return v
	}
	v, q.r := q.r[0], q.r[1:]
	return v
}

func (q *Deque[T]) PopBack() T {
	if q.Empty() {
		panic("cannot pop from back: deque is Empty")
	}
	if len(q.r) > 0 {
		v, q.r := q.r[len(q.r)-1], q.r[:len(q.r)-1]
		return v
	}
	v, q.l := q.l[0], q.l[1:]
	return v
}

func (q *Deque[T]) Front() T {
	if q.Empty() {
		panic("cannot access front: deque is Empty")
	}
	if len(q.l) > 0 {
		return q.l[len(q.l)-1]
	}
	return q.r[0]
}

func (q *Deque[T]) Back() T {
	if q.Empty() {
		panic("cannot access back: deque is Empty")
	}
	if len(q.r) > 0 {
		return q.r[len(q.r)-1]
	}
	return q.l[0]
}

func (q *Deque[T]) Get(i int) T {
	if i < 0 || i >= q.Size() {
		panic("index out of bounds")
	}
	if i < len(q.l) {
		return q.l[len(q.l)-1-i]
	}
	return q.r[i-len(q.l)]
}