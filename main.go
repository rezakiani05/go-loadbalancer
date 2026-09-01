package main

import (
	"fmt"
	"go-loadbalancer/balancer"
	"go-loadbalancer/config"
	"go-loadbalancer/proxy"
	"go-loadbalancer/ratelimiter"
	"go-loadbalancer/server"
	"net/http"
	"time"
)

func main() {
	// TODO:
	// ۱. بارگذاری تنظیمات با config.LoadConfig().
	// ۲. ساخت لیست backendها و مقداردهی اولیه Health Check دوره‌ای (با time.Ticker روی یک Goroutine مجزا).
	// ۳. ساخت RateLimiter و LoadBalancer.
	// ۴. ایجاد GatewayHandler.
	// ۵. اجرای سرور HTTP روی پورت تنظیم شده با http.ListenAndServe.
	fmt.Println("Starting API Gateway & Load Balancer...")
}
