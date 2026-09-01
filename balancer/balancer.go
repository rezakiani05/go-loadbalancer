package balancer

import (
	"go-loadbalancer/server"
	"sync/atomic"
)

type LoadBalancer struct {
	backends []*server.Backend
	current  uint64
}

func NewLoadBalancer(backends []*server.Backend) *LoadBalancer {
	return &LoadBalancer{
		backends: backends,
	}
}

func (lb *LoadBalancer) GetNextValidBackend() *server.Backend {
	n := len(lb.backends)
	if n == 0 {
		return nil
	}

	next := atomic.AddUint64(&lb.current, 1)

	for i := 0; i < n; i++ {
		idx := int((next + uint64(i)) % uint64(n))
		if lb.backends[idx].IsAlive() {
			return lb.backends[idx]
		}
	}

	return nil
}
