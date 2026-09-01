package config

import "time"

type Config struct {
	Port            string        // پورتی که لودبالانسر روی آن اجرا می‌شود (مثلا ":8080")
	BackendURLs     []string      // آدرس سرورهای مقصد (مثلا ["http://localhost:8081", "http://localhost:8082"])
	RateLimit       int           // حداکثر تعداد درخواست مجاز در بازه زمانی
	RateLimitWindow time.Duration // بازه زمانی Rate Limit (مثلا ۱ ثانیه)
}

func LoadConfig() *Config {
	// TODO:
	// ۱. یک نمونه از Struct بالا با مقادیر پیش‌فرض ایجاد کن.
	// ۲. آدرس ۲ یا ۳ سرور مجازی (مثلاً http://localhost:8081 و http://localhost:8082) را در BackendURLs بگذار.
	// ۳. RateLimit را برابر ۱۰ و RateLimitWindow را برابر ۱ ثانیه قرار بده.
	return nil
}
