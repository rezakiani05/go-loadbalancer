package proxy

import (
	"go-loadbalancer/balancer"
	"go-loadbalancer/ratelimiter"
	"net/http"
)

type GatewayHandler struct {
	lb          *balancer.LoadBalancer
	rateLimiter *ratelimiter.RateLimiter
}

func NewGatewayHandler(lb *balancer.LoadBalancer, rl *ratelimiter.RateLimiter) *GatewayHandler {
	// TODO: مقداردهی ساختار GatewayHandler.
	return nil
}

func (g *GatewayHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// TODO:
	// ۱. استخراج IP کاربر از r.RemoteAddr.
	// ۲. چک کردن با g.rateLimiter.Allow(ip). اگر false بود:
	//    http.Error(w, "Too Many Requests", http.StatusTooManyRequests) بده و return کن.
	// ۳. گرفتن سرور بعدی با backend := g.lb.GetNextValidBackend().
	//    اگر backend nil بود: http.Error(w, "Service Unavailable", http.StatusServiceUnavailable) بده.
	// ۴. ساخت یک ReverseProxy با httputil.NewSingleHostReverseProxy(backend.URL) و اجرای ServeHTTP(w, r) روی آن.
}
