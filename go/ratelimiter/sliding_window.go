package ratelimiter

import (
	"dsa/utils/queue"
	"time"
)

type SlidingWindow struct {
	requests        queue.Queue[int64]
	size            int
	lastWindowStart time.Time
	windowSizeInSec int
}

func NewSlidingWindow(sizeInSec int, cap int) SlidingWindow {
	return SlidingWindow{
		size:            0,
		lastWindowStart: time.Now(),
		requests:        queue.NewQueue[int64](cap),
		windowSizeInSec: sizeInSec,
	}
}

func (r *SlidingWindow) Allow() bool {
	newStart := time.Now().Add(time.Second * time.Duration(-r.windowSizeInSec))
	r.lastWindowStart = newStart
	for !r.requests.IsEmpty() {
		if request, ok := r.requests.Front(); ok && request <= newStart.UnixMilli() {
			r.requests.Dequeue()
		} else {
			break
		}
	}
	return r.requests.Enqueue(time.Now().UnixMilli())
}
