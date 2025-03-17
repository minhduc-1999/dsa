package ratelimiter_test

import (
	"dsa/ratelimiter"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestTokenBucket(t *testing.T) {
	// Create a new token bucket with capacity 5 and rate 1 token per second
	tb := ratelimiter.NewTokenBucket(5, 5)

	// Initially, the bucket should have 5 tokens
	assert.True(t, tb.Allow())
	assert.True(t, tb.Allow())
	assert.True(t, tb.Allow())
	assert.True(t, tb.Allow())
	assert.True(t, tb.Allow())
	assert.False(t, tb.Allow())
	time.Sleep(500 * time.Millisecond)
	assert.True(t, tb.Allow())
	assert.True(t, tb.Allow())
	assert.False(t, tb.Allow())
}
