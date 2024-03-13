package lrucache

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_LRUCache(t *testing.T) {
	cache := Constructor(2)
	cache.Put(1, 1)
	cache.Put(2, 2)
	v := cache.Get(1)
	assert.Equal(t, 1, v)
	cache.Put(3, 3)
	v = cache.Get(2)
	assert.Equal(t, -1, v)
	cache.Put(4, 4)
	v = cache.Get(1)
	assert.Equal(t, -1, v)
	v = cache.Get(3)
	assert.Equal(t, 3, v)
	v = cache.Get(4)
	assert.Equal(t, 4, v)
}
