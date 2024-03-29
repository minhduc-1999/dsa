package stack

import (
	"fmt"
)

type Stack[T any] struct {
	arr  []T
	size int
	cap  int
}

func NewStack[T any](cap int) Stack[T] {
	return Stack[T]{
		arr:  make([]T, 0, cap),
		size: 0,
		cap:  cap,
	}
}

func (s Stack[T]) Size() int {
  return s.size
}

func (s *Stack[T]) Push(item T) error {
	if s.size >= s.cap {
		return fmt.Errorf("Stack is full")
	}
	s.size++
  s.arr = s.arr[:s.size]
	s.arr[s.size-1] = item
	return nil
}

func (s *Stack[T]) Pop() {
	if s.size == 0 {
    return
	}
	s.size--
  s.arr = s.arr[:s.size]
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

func (s Stack[T]) Data() []T {
  return s.arr
}
