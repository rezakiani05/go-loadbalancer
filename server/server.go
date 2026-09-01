package server

import (
	"net/http"
	"net/url"
	"sync"
	"time"
)

type Backend struct {
	URL   *url.URL
	Alive bool
	mux   sync.RWMutex
}

func (b *Backend) SetAlive(alive bool) {
	b.mux.Lock()
	defer b.mux.Unlock()
	b.Alive = alive
}

func (b *Backend) IsAlive() bool {
	b.mux.RLock()
	defer b.mux.RUnlock()
	return b.Alive
}

func (b *Backend) CheckHealth() {
	client := http.Client{
		Timeout: 2 * time.Second,
	}

	resp, err := client.Get(b.URL.String())
	if err != nil {
		b.SetAlive(false)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 200 && resp.StatusCode <= 299 {
		b.SetAlive(true)
	} else {
		b.SetAlive(false)
	}
}
