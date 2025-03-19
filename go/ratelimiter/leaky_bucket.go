package ratelimiter

import (
	"sync/atomic"
	"time"
)

type LeakyBucket struct {
	leakRate float64
	token    atomic.Int32
	cap      int
	lastLeak time.Time
}

func NewLeakyBucket(rate float64, cap int) LeakyBucket {
	return LeakyBucket{
		leakRate: rate,
		token:    atomic.Int32{},
		cap:      cap,
		lastLeak: time.Now(),
	}
}

func (b *LeakyBucket) Allow() bool {
	elapsed := time.Since(b.lastLeak).Seconds()
	leakedToken := int32(elapsed * float64(b.leakRate))
	if leakedToken > int32(b.token.Load()) {
		b.token.Swap(0)
	} else {
		b.token.Add(-leakedToken)
	}
	b.lastLeak = time.Now()
	if int(b.token.Load()+1) > b.cap {
		return false
	}
	b.token.Add(1)
	return true
}
