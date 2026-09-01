package config

import "time"

type Config struct {
	Port            string
	BackendURLs     []string
	RateLimit       int
	RateLimitWindow time.Duration
}

func LoadConfig() *Config {
	return &Config{
		Port: ":8080",
		BackendURLs: []string{
			"http://localhost:8081",
			"http://localhost:8082",
		},
		RateLimit:       10,
		RateLimitWindow: time.Second,
	}
}
