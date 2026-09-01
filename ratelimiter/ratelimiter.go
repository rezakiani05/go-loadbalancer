package ratelimiter

import (
	"sync"
	"time"
)

type client struct {
	count     int
	lastReset time.Time
}

type RateLimiter struct {
	mu      sync.Mutex
	limit   int
	window  time.Duration
	clients map[string]*client
}

func NewRateLimiter(limit int, window time.Duration) *RateLimiter {
	return &RateLimiter{
		limit:   limit,
		window:  window,
		clients: make(map[string]*client),
	}
}

func (rl *RateLimiter) Allow(ip string) bool {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	now := time.Now()
	c, exists := rl.clients[ip]

	if !exists || now.Sub(c.lastReset) > rl.window {
		rl.clients[ip] = &client{
			count:     1,
			lastReset: now,
		}
		return true
	}

	if c.count < rl.limit {
		c.count++
		return true
	}

	return false
}
