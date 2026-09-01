package balancer

import (
	"go-loadbalancer/server"
)

type LoadBalancer struct {
	backends []*server.Backend
	current  uint64 // ایندکس چرخشی
}

func NewLoadBalancer(backends []*server.Backend) *LoadBalancer {
	// TODO: مقداردهی اولیه LoadBalancer با اسلایس backends.
	return nil
}

func (lb *LoadBalancer) GetNextValidBackend() *server.Backend {
	// TODO:
	// ۱. با atomic.AddUint64(&lb.current, 1) مقدار current را به صورت Safe یکی اضافه کن.
	// ۲. با استفاده از باقیمانده تقسیم (index = current % len(lb.backends)) سرور بعدی را انتخاب کن.
	// ۳. یک حلقه بزن تا سروری را پیدا کنی که IsAlive() آن true باشد.
	// ۴. اگر هیچ سروری سالم نبود، nil برگردان.
	return nil
}
