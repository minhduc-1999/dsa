package ratelimiter

import "time"

type FixedWindow struct {
	windowSize  int
	windowStart time.Time
	maxRequest  int
	requests    int
}

func NewFixedWindow(windowSizeInSec int, maxRequest int) FixedWindow {
	return FixedWindow{
		windowSize:  windowSizeInSec,
		maxRequest:  maxRequest,
		requests:    0,
		windowStart: time.Now(),
	}
}

func (r *FixedWindow) Allow() bool {
	elapsedInSecond := time.Since(r.windowStart).Seconds()
	if elapsedInSecond > float64(r.windowSize) {
		r.requests = 0
		r.windowStart = time.Now()
	}
	if r.requests >= r.maxRequest {
		return false
	}
	r.requests++
	return true
}
