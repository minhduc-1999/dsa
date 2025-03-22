package queue

type Queue[T any] struct {
	arr        []T
	frontIndex int
	cap        int
	len        int
}

func NewQueue[T any](cap int) Queue[T] {
	return Queue[T]{
		arr:        make([]T, cap),
		frontIndex: 0,
		len:        0,
		cap:        cap,
	}
}

func (q Queue[T]) IsEmpty() bool {
	return q.len == 0
}

func (q Queue[T]) Size() int {
	return q.len
}

func (q Queue[T]) Front() (T, bool) {
	if q.Size() == 0 {
		var result T
		return result, false
	}
	return q.arr[q.frontIndex], true
}

func (q *Queue[T]) Enqueue(item T) bool {
	if q.len >= q.cap {
		return false
	}
	q.arr[q.len] = item
	q.len++
	return true
}

func (q *Queue[T]) Dequeue() (T, bool) {
	var v T
	if q.len == 0 {
		return v, false
	}
	v = q.arr[q.frontIndex]
	if q.frontIndex++; q.frontIndex == q.cap {
		q.frontIndex = 0
	}
	if q.len--; q.len < 0 {
		q.len = 0
	}
	return v, true
}
