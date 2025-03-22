package queue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQueue(t *testing.T) {
	queue := NewQueue[int](10)

	assert.True(t, queue.IsEmpty())
	assert.Equal(t, 0, queue.Size())
	_, ok := queue.Dequeue()
	assert.False(t, ok)
	_, ok = queue.Front()

	assert.False(t, ok)
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)

	assert.Equal(t, false, queue.IsEmpty())
	assert.Equal(t, 3, queue.Size())

	frontValue, ok := queue.Front()
	assert.True(t, ok)
	assert.Equal(t, 1, frontValue)
	frontValue, ok = queue.Dequeue()
	assert.True(t, ok)
	assert.Equal(t, 1, frontValue)

	frontValue, ok = queue.Front()
	assert.True(t, ok)
	assert.Equal(t, 2, frontValue)
	frontValue, ok = queue.Dequeue()
	assert.True(t, ok)
	assert.Equal(t, 2, frontValue)

	frontValue, ok = queue.Front()
	assert.True(t, ok)
	assert.Equal(t, 3, frontValue)
	frontValue, ok = queue.Dequeue()
	assert.True(t, ok)
	assert.Equal(t, 3, frontValue)

	assert.Equal(t, true, queue.IsEmpty())
}
