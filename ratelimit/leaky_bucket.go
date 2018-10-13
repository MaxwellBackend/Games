package ratelimit

import (
	"time"
)

// 漏桶算法
type LeakyBucket struct {
	capacity   int           // bucket总容量
	interval   time.Duration // bucket满需要花费的时间
	remain     int           // 剩余数量
	lastAccess time.Time     // 上次访问时间
}

func NewLeakyBucket(interval time.Duration, capacity int) *LeakyBucket {
	return &LeakyBucket{
		interval:   interval,
		capacity:   capacity,
		remain:     capacity,
		lastAccess: time.Now(),
	}
}

func (l *LeakyBucket) Access() bool {
	now := time.Now()

	since := now.Sub(l.lastAccess)

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
