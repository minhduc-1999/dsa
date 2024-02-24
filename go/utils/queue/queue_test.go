package queue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQueue(t *testing.T) {
	queue := NewQueue[int](10)
	want := Queue[int]{
		arr:        make([]int, 0, 10),
		frontIndex: 0,
		backIndex:  0,
		cap:        10,
	}
	assert.Equal(t, want, queue)

	queue.Enqueue(1)
  queue.Enqueue(2)
  queue.Enqueue(3)

	assert.Equal(t, false, queue.IsEmpty())
  assert.Equal(t, 3, queue.Size())

  frontValue, _ := queue.Front()
	assert.Equal(t, 1, frontValue)

  queue.Dequeue()
  frontValue, _ = queue.Front()
  assert.Equal(t, 2, frontValue)

  queue.Dequeue()
  frontValue, _ = queue.Front()
  assert.Equal(t, 3, frontValue)

  queue.Dequeue()
	assert.Equal(t, true, queue.IsEmpty())
}
