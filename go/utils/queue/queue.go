package queue

import (
	"cmp"
	"fmt"
)

type Queue[T cmp.Ordered] struct {
  arr []T
  frontIndex int
  backIndex int
  cap int
}

func NewQueue[T cmp.Ordered](cap int) Queue[T] {
	return Queue[T]{
		arr:  make([]T, 0, cap),
		frontIndex: 0,
    backIndex: 0,
		cap:  cap,
	}
}

func (q Queue[T]) IsEmpty() bool {
  return q.Size() == 0
}

func (q Queue[T]) Size() int {
  if q.backIndex > q.frontIndex {
    return q.backIndex - q.frontIndex + 1
  }
  return q.cap - q.frontIndex + q.backIndex + 1
}

func (q Queue[T]) Front() (T, error) {
  if q.Size() == 0 {
    var result T
    return result, fmt.Errorf("Queue is empty")
  }
  return q.arr[q.frontIndex], nil
}

func (q Queue[T]) Back() (T, error) {
  if q.Size() == 0 {
    var result T
    return result, fmt.Errorf("Queue is empty")
  }
  return q.arr[q.backIndex], nil
}

func (q *Queue[T]) Enqueue(item T) {
  if q.Size() >= q.cap {
    return
  }
  q.backIndex = (q.backIndex + 1) % q.cap
  q.arr[q.backIndex] = item
}

func (q *Queue[T]) Dequeue(item T) {
  if q.Size() == 0 {
    return
  }
  q.frontIndex = (q.frontIndex + 1 ) % q.cap
}
