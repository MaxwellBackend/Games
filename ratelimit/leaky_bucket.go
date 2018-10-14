package ratelimit

import (
	"time"
)

// 漏桶算法
type LeakyBucket struct {
	capacity     int           // bucket总容量
	interval     time.Duration // 漏出水滴的时间
	inDrops      int           // 当前bucket中的水滴数量
	lastLeakTime time.Time     // 上次漏出时间
}

func NewLeakyBucket(interval time.Duration, capacity int) *LeakyBucket {
	return &LeakyBucket{
		interval:     interval,
		capacity:     capacity,
		inDrops:      0,
		lastLeakTime: time.Now(),
	}
}

func (l *LeakyBucket) Access() bool {
	now := time.Now()

	since := now.Sub(l.lastLeakTime)

	// 漏出的水滴数量
	leaks := int(float64(since) / float64(l.interval))

	if leaks > 0 {
		if l.inDrops <= leaks {
			// 重置漏桶中的请求数量
			l.inDrops = 0
		} else {
			// 减少漏桶中的请求数量
			l.inDrops -= leaks
		}
		l.lastLeakTime = now
	}

	// 漏桶未满
	if l.inDrops < l.capacity {
		l.inDrops++
		return true
	}

	// 漏桶已满
	return false
}
