type Queue[T any] struct {
	data []T
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{data: make([]T, 0)}
}

func (q *Queue[T]) Empty() bool {
	return len(q.data) == 0
}

func (q *Queue[T]) Len() int {
	return len(q.data)
}

func (q *Queue[T]) Push(v T) {
	q.data = append(q.data, v)
}

func (q *Queue[T]) Pop() T {
	if q.Empty() {
		panic("cannot dequeue: queue is empty")
	}
	v := q.data[0]
	q.data = q.data[1:]
	return v
}

func (q *Queue[T]) Front() T {
	if q.Empty() {
		panic("cannot access front: queue is empty")
	}
	return q.data[0]
}

func (q *Queue[T]) Back() T {
	if q.Empty() {
		panic("cannot access back: queue is empty")
	}
	return q.data[len(q.data)-1]
}