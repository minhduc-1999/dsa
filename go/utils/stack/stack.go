package stack

import (
	"cmp"
	"fmt"
)

type Stack[T cmp.Ordered] struct {
	arr  []T
	size int
	cap  int
}

func NewStack[T cmp.Ordered](cap int) Stack[T] {
	return Stack[T]{
		arr:  make([]T, 0, cap),
		size: 0,
		cap:  cap,
	}
}

func (s *Stack[T]) Push(item T) error {
	if s.size >= s.cap {
		return fmt.Errorf("Stack is full")
	}
	s.size++
	s.arr[s.size-1] = item
	return nil
}

func (s *Stack[T]) Pop() {
	if s.size == 0 {
    return
	}
	s.size--
}

func (s Stack[T]) Top() (T, error) {
	if s.size == 0 {
		var result T
		return result, fmt.Errorf("Stack is empty")
	}
	return s.arr[s.size-1], nil
}

func (s Stack[T]) IsEmpty() bool {
	return s.size == 0
}
