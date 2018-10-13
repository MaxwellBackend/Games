package ratelimit

import (
	"testing"
	"time"
	"fmt"
)

func TestTokenBucket_Access(t *testing.T) {
	tb := NewTokenBucket(time.Second, 10)

	// 前10次操作
	for i := 0; i < 10; i ++ {
		if !tb.Access() {
			t.Error(fmt.Sprintf("%v, Access fail", i))
		}
	}

	// 延时90毫秒
	time.Sleep(time.Millisecond * 90)

	// 第11次操作
	if !tb.Access() {
		t.Error("11 Access fail")
	}

}
