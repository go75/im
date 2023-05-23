package utils

import (
	"time"
)

// 限流器
type Limiter struct {
	// 令牌桶
	bucket chan struct{}
	duration time.Duration
}

func NewLimiter(duration time.Duration, bucketCapacity uint32) *Limiter {
	limiter := &Limiter{
		bucket: make(chan struct{}, bucketCapacity),
		duration: duration,
	}
	go func() {
		for {
			time.Sleep(duration)
			limiter.bucket <- struct{}{}
		}
	}()
	return limiter
}

func (l *Limiter) Allow() bool {
	select {
	case <- l.bucket:
		return true
	default:
		return false
	}
}
