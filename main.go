package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"time"

	"go-loadbalancer/balancer"
	"go-loadbalancer/config"
	"go-loadbalancer/proxy"
	"go-loadbalancer/ratelimiter"
	"go-loadbalancer/server"
)

func main() {
	cfg := config.LoadConfig()

	var backends []*server.Backend
	for _, urlStr := range cfg.BackendURLs {
		parsedURL, err := url.Parse(urlStr)
		if err != nil {
			log.Fatalf("Invalid backend URL %s: %v", urlStr, err)
		}
		backends = append(backends, &server.Backend{
			URL:   parsedURL,
			Alive: true,
		})
	}

	ticker := time.NewTicker(10 * time.Second)
	go func() {
		for range ticker.C {
			for _, b := range backends {
				b.CheckHealth()
			}
		}
	}()

	rl := ratelimiter.NewRateLimiter(cfg.RateLimit, cfg.RateLimitWindow)
	lb := balancer.NewLoadBalancer(backends)
	handler := proxy.NewGatewayHandler(lb, rl)

	fmt.Printf("API Gateway & Load Balancer is running on %s\n", cfg.Port)
	if err := http.ListenAndServe(cfg.Port, handler); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
