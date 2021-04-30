package ratelimit

import (
	"sync/atomic"
	"time"

	grpc_ratelimit "github.com/grpc-ecosystem/go-grpc-middleware/ratelimit"
)

func NewLimiter(limit int64) grpc_ratelimit.Limiter {
	return &rateLimiter{
		limit: limit,
		count: 0,
	}
}

type rateLimiter struct {
	count int64
	limit int64
}

func (rl *rateLimiter) Limit() bool {
	if atomic.LoadInt64(&rl.count) > rl.limit {
		return true
	}
	atomic.StoreInt64(&rl.count, atomic.AddInt64(&rl.count, 1)) // count++
	_ = time.AfterFunc(time.Second, func() {
		atomic.StoreInt64(&rl.count, atomic.AddInt64(&rl.count, -1)) // count--
	})
	return false
}
