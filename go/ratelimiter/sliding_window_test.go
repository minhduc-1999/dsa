package ratelimiter_test

import (
	"dsa/ratelimiter"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestSlidingWindow(t *testing.T) {
	r := ratelimiter.NewSlidingWindow(2, 5)

	assert.True(t, r.Allow())
	time.Sleep(200 * time.Millisecond)
	assert.True(t, r.Allow())
	time.Sleep(200 * time.Millisecond)
	assert.True(t, r.Allow())
	time.Sleep(200 * time.Millisecond)
	assert.True(t, r.Allow())
	assert.True(t, r.Allow())
	assert.False(t, r.Allow())
	time.Sleep(1500 * time.Millisecond)
	assert.True(t, r.Allow())
	assert.False(t, r.Allow())
}
