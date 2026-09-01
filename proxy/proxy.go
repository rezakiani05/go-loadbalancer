package proxy

import (
	"net"
	"net/http"
	"net/http/httputil"

	"go-loadbalancer/balancer"
	"go-loadbalancer/ratelimiter"
)

type GatewayHandler struct {
	lb          *balancer.LoadBalancer
	rateLimiter *ratelimiter.RateLimiter
}

func NewGatewayHandler(lb *balancer.LoadBalancer, rl *ratelimiter.RateLimiter) *GatewayHandler {
	return &GatewayHandler{
		lb:          lb,
		rateLimiter: rl,
	}
}

func (g *GatewayHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	host, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		host = r.RemoteAddr
	}

	if !g.rateLimiter.Allow(host) {
		http.Error(w, "Too Many Requests", http.StatusTooManyRequests)
		return
	}

	backend := g.lb.GetNextValidBackend()
	if backend == nil {
		http.Error(w, "Service Unavailable", http.StatusServiceUnavailable)
		return
	}

	proxy := httputil.NewSingleHostReverseProxy(backend.URL)
	proxy.ServeHTTP(w, r)
}
