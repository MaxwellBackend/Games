package ratelimit

import (
	"testing"
	"time"
	"fmt"
)

func TestLeakyBucket_Access(t *testing.T) {
	lb := NewLeakyBucket(time.Second, 10)

	// 前10次操作
	for i := 0; i < 10; i++ {
		if !lb.Access() {
			t.Error(fmt.Sprintf("%v, Access fail", i))
		}
	}

	// 延时999毫秒
	time.Sleep(time.Millisecond * 999)

	// 第11次操作
	if !lb.Access() {
		t.Error("11 Access fail")
	}
}
