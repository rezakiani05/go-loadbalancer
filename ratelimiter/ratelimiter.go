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
	clients map[string]*client // نگهداری وضعیت هر IP
}

func NewRateLimiter(limit int, window time.Duration) *RateLimiter {
	// TODO:
	// مقداردهی اولیه ساختار RateLimiter به همراه یک map خالی برای clients.
	return nil
}

func (rl *RateLimiter) Allow(ip string) bool {
	// TODO:
	// ۱. با rl.mu.Lock() قفل کن.
	// ۲. چک کن آیا این IP در map وجود دارد یا نه. اگر نبود، بسازش.
	// ۳. اگر از زمان lastReset بیشتر از rl.window گذشته بود، count را ۱ کن و lastReset را بروزرسانی کن.
	// ۴. اگر هنوز درون بازه زمانی هستیم:
	//    - اگر count < limit بود: یکی به count اضافه کن و true برگردان (اجازه دسترسی).
	//    - در غیر این صورت: false برگردان (بلاک کردن درخواست).
	return false
}
