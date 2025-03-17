package ratelimiter

import "time"

type TokenBucket struct {
	token int
	cap   int
	// req/s
	rate     int
	lastFill time.Time
}

func NewTokenBucket(rate int, cap int) TokenBucket {
	return TokenBucket{
		token:    cap,
		cap:      cap,
		rate:     rate,
		lastFill: time.Now(),
	}
}

func (b *TokenBucket) Allow() bool {
	elapsed := time.Since(b.lastFill)
	b.token = min(b.cap, b.token+int(elapsed.Seconds()*float64(b.rate)))
	b.lastFill = time.Now()
	if allow := b.token > 0; allow {
		b.token -= 1
		return true
	}
	return false
}
