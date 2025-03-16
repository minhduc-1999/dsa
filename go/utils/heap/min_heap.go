package heap

import "cmp"

type MinHeap[T cmp.Ordered] struct {
	arr []T
	cap int
}

func NewMinHeap[T cmp.Ordered](cap int) MinHeap[T] {
	return MinHeap[T]{
		arr: make([]T, 0, cap),
		cap: cap,
	}
}

func (h MinHeap[T]) left(i int) int {
	return 2*i + 1
}

func (h MinHeap[T]) right(i int) int {
	return 2*i + 2
}

func (h MinHeap[T]) parent(i int) int {
	return (i - 1) / 2
}

func (h *MinHeap[T]) ExtractMin() (T, bool) {
	var result T
	if len(h.arr) == 0 {
		return result, false
	}
	min := h.arr[0]
	h.arr[0] = h.arr[len(h.arr)-1]
	h.arr = h.arr[:len(h.arr)-1]
	h.heapify(0)
	return min, true
}

func (h MinHeap[T]) GetMin() (T, bool) {
	var result T
	if len(h.arr) == 0 {
		return result, false
	}
	return h.arr[0], true
}

func (h *MinHeap[T]) Insert(value T) bool {
	if len(h.arr) == h.cap {
		return false
	}
	h.arr = append(h.arr, value)
	i := len(h.arr) - 1
	for i != 0 && h.arr[h.parent(i)] > h.arr[i] {
		h.swap(i, h.parent(i))
		i = h.parent(i)
	}
	return true
}

func (h *MinHeap[T]) heapify(i int) {
	if i >= len(h.arr) {
		return
	}
	l := h.left(i)
	r := h.right(i)
	iMin := i
	if l < len(h.arr) && h.arr[iMin] > h.arr[l] {
		iMin = l
	}
	if r < len(h.arr) && h.arr[iMin] > h.arr[r] {
		iMin = r
	}
	if iMin != i {
		h.swap(iMin, i)
		h.heapify(iMin)
	}
}

func (h *MinHeap[T]) swap(i, j int) {
	h.arr[i], h.arr[j] = h.arr[j], h.arr[i]
}
