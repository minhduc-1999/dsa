package stack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateStack(t *testing.T) {
	stack := NewStack[int](10)
	want := Stack[int]{
		arr:  make([]int, 0, 10),
		size: 0,
		cap:  10,
	}
	assert.Equal(t, want, stack)
}

func TestStackPushItem(t *testing.T) {
	stack := NewStack[int](10)
	want := Stack[int]{
		arr:  make([]int, 0, 10),
		size: 0,
		cap:  10,
	}
	assert.Equal(t, want, stack)
  assert.Equal(t, true, stack.IsEmpty())

	stack.Push(1)
	stack.Push(2)

  assert.Equal(t, false, stack.IsEmpty())
  assert.Equal(t, 2, stack.Size())

  top, _ := stack.Top()
  assert.Equal(t, 2, top)

  stack.Pop()
  top, _ = stack.Top()
  assert.Equal(t, 1, top)

  stack.Pop()
  assert.Equal(t, true, stack.IsEmpty())
}
