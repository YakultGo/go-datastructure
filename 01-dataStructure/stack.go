type Stack[T any] struct {
	data []T
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{data: make([]T, 0)}
}

func (s *Stack[T]) Push(v T) {
	s.data = append(s.data, v)
}

func (s *Stack[T]) Pop() (v T) {
	if s.Empty() {
		panic("stack is empty")
	}
	v, s.data = s.data[len(s.data)-1], s.data[:len(s.data)-1]
	return
}

func (s *Stack[T]) Top() T {
	if s.Empty() {
		panic("stack is empty")
	}
	return s.data[len(s.data)-1]
}

func (s *Stack[T]) Empty() bool {
	return len(s.data) == 0
}

func (s *Stack[T]) Size() int {
	return len(s.data)
}