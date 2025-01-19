package main

import (
	"sync"
	"time"
)

type RateLimiter struct {
	rate      int
	interval  time.Duration
	tokens    int
	lastCheck time.Time
	mu        sync.Mutex
}

func NewRateLimiter(rate int, interval time.Duration) *RateLimiter {
	return &RateLimiter{
		rate:      rate,
		interval:  interval,
		tokens:    rate,
		lastCheck: time.Now(),
	}
}

func (rl *RateLimiter) Allow() bool {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	now := time.Now()
	elapsed := now.Sub(rl.lastCheck)
	rl.lastCheck = now

	rl.tokens += int(elapsed / rl.interval)
	if rl.tokens > rl.rate {
		rl.tokens = rl.rate
	}

	if rl.tokens > 0 {
		rl.tokens--
		return true
	}

	return false
}
