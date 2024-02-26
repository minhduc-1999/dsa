package queue

import (
	"fmt"
)

type Queue[T any] struct {
	arr        []T
	frontIndex int
	backIndex  int
	cap        int
}

func NewQueue[T any](cap int) Queue[T] {
	return Queue[T]{
		arr:        make([]T, 0, cap),
		frontIndex: 0,
		backIndex:  0,
		cap:        cap,
	}
}

func (q Queue[T]) IsEmpty() bool {
	return q.Size() == 0
}

func (q Queue[T]) Size() int {
	return q.backIndex - q.frontIndex
}

func (q Queue[T]) Front() (T, error) {
	if q.Size() == 0 {
		var result T
		return result, fmt.Errorf("Queue is empty")
	}
	return q.arr[q.frontIndex], nil
}

func (q *Queue[T]) Enqueue(item T) {
	if q.Size() >= q.cap {
		return
	}
  q.arr = append(q.arr, item)
	q.backIndex++
}

func (q *Queue[T]) Dequeue() {
	if q.Size() == 0 {
		return
	}
  copy(q.arr, q.arr[q.frontIndex+1:q.backIndex])
  q.backIndex--
  q.arr = q.arr[:q.backIndex]
}
