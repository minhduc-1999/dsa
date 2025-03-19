package ratelimiter_test

import (
	"dsa/ratelimiter"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestFixedWindow(t *testing.T) {
	tb := ratelimiter.NewFixedWindow(2, 5)

	assert.True(t, tb.Allow())
	assert.True(t, tb.Allow())
	assert.True(t, tb.Allow())
	assert.True(t, tb.Allow())
	assert.True(t, tb.Allow())
	assert.False(t, tb.Allow())
	time.Sleep(2 * time.Second)
	assert.True(t, tb.Allow())
	assert.True(t, tb.Allow())
	assert.True(t, tb.Allow())
	assert.True(t, tb.Allow())
	assert.True(t, tb.Allow())
	assert.False(t, tb.Allow())
}
