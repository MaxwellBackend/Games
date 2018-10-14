package ratelimit

import (
	"time"
)

// 令牌桶算法
type TokenBucket struct {
	capacity       int           // bucket总容量
	interval       time.Duration // bucket满需要花费的时间
	remain         int           // 剩余数量
	lastAccessTime time.Time     // 上次访问时间
}

func NewTokenBucket(interval time.Duration, capacity int) *TokenBucket {
	return &TokenBucket{
		interval:       interval,
		capacity:       capacity,
		remain:         capacity,
		lastAccessTime: time.Now(),
	}
}

func (l *TokenBucket) Access() bool {
	now := time.Now()

	since := now.Sub(l.lastAccessTime)

	if since >= l.interval {
		// 充满
		l.remain = l.capacity
	} else {
		// 补充
		add := int(float64(since) / float64(l.interval) * float64(l.capacity))

		if add+l.remain >= l.capacity {
			l.remain = l.capacity
		} else {
			l.remain = l.remain + add
		}
	}

	// 如果不够了，则直接返回失败
	if l.remain <= 0 {
		return false
	}

	// 扣除
	l.remain--
	return true
}
