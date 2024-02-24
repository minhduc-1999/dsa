package heap

import "cmp"

type MaxHeap[T cmp.Ordered] struct {
	arr      []T
	size     int
	capacity int
}

func NewMaxHeap[T cmp.Ordered](cap int) MaxHeap[T] {
	return MaxHeap[T]{
		arr:      make([]T, 0, cap),
		size:     0,
		capacity: cap,
	}
}

func New[T cmp.Ordered](arr []T) MaxHeap[T] {
	heap := MaxHeap[T]{
		arr:      arr,
		size:     len(arr),
		capacity: len(arr),
	}
	n := len(arr)/2 - 1
	for i := n; i >= 0; i-- {
		heap.heapify(i)
	}
	return heap
}

func (h MaxHeap[T]) left(i int) int {
	return i*2 - 1
}

func (h MaxHeap[T]) right(i int) int {
	return i*2 + 1
}

func (h MaxHeap[T]) parent(i int) int {
	return (i - 1) / 2
}

func (h *MaxHeap[T]) Insert(item T) {
	if h.size >= h.capacity {
		return
	}
	h.size += 1
	i := h.size - 1
	h.arr[i] = item
	for i != 0 && h.arr[i] > h.arr[h.parent(i)] {
		h.swap(i, h.parent(i))
		i = h.parent(i)
	}
}

func (h *MaxHeap[T]) heapify(index int) {
	if index >= h.size {
		return
	}
	max := index

	if h.arr[max] < h.arr[h.left(index)] {
		max = h.left(index)
	}

	if h.arr[max] < h.arr[h.right(index)] {
		max = h.right(index)
	}

	if max != index {
		h.swap(index, max)
		h.heapify(max)
	}
}

func (h *MaxHeap[T]) swap(lhs int, rhs int) {
	temp := h.arr[lhs]
	h.arr[lhs] = h.arr[rhs]
	h.arr[rhs] = temp
}
