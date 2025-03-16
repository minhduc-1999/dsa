package heap_test

import (
	"dsa/utils/heap"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinHeap(t *testing.T) {
	heap := heap.NewMinHeap[int](10)

	// Test Insert
	heap.Insert(10)
	heap.Insert(5)
	heap.Insert(8)
	heap.Insert(3)

	// Test GetMin
	min, success := heap.GetMin()
	assert.True(t, success)
	assert.Equal(t, 3, min)

	// Test ExtractMin
	min, success = heap.ExtractMin()
	assert.True(t, success)
	assert.Equal(t, 3, min)

	// After extraction, the next minimum should be 5
	min, success = heap.GetMin()
	assert.True(t, success)
	assert.Equal(t, 5, min)

	// Test ExtractMin again
	min, success = heap.ExtractMin()
	assert.True(t, success)
	assert.Equal(t, 5, min)

	// Test ExtractMin again
	min, success = heap.ExtractMin()
	assert.True(t, success)
	assert.Equal(t, 8, min)

	// Test ExtractMin again (heap should be left with only 10)
	min, success = heap.ExtractMin()
	assert.True(t, success)
	assert.Equal(t, 10, min)

	// Test ExtractMin on empty heap
	_, success = heap.ExtractMin()
	assert.False(t, success)

	// Test GetMin on empty heap
	_, success = heap.GetMin()
	assert.False(t, success)
}
