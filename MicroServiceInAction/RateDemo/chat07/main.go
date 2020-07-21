package main

import (
	"fmt"
	"sync"
	"time"
)

type TokenBucket struct {
	rate         int64 //固定token放入速率
	capacity     int64 //桶的容量
	tokens       int64 //当前桶中的token数
	lastTokenSec int64 //上次放入token到桶的时间戳
	sync.Mutex
}

func (l *TokenBucket) Allow() bool {
	l.Lock()
	defer l.Unlock()

	now := time.Now().Unix()
	l.tokens += (now - l.lastTokenSec) * l.rate
	if l.tokens > l.capacity {
		l.tokens = l.capacity
	}
	l.lastTokenSec = now
	if l.tokens > 0 {
		l.tokens--
		return true
	}
	return false
}

func (l *TokenBucket) Set(r, c int64) {
	l.rate = r
	l.capacity = c
	l.lastTokenSec = time.Now().Unix()
}

func main() {

	var wg sync.WaitGroup
	var rate TokenBucket
	rate.Set(3, 2)
	time.Sleep(time.Second)
	for i := 0; i < 10; i++ {
		wg.Add(1)
		fmt.Print(" Create Req:", i, time.Now().Format(time.RFC3339))
		go func(i int) {
			defer wg.Done()
			if rate.Allow() {
				fmt.Println(" Resp succeed.", i, time.Now().Format(time.RFC3339))
			} else {
				fmt.Println(" Resp failed.")
			}
		}(i)
		time.Sleep(200 * time.Millisecond)
	}
	wg.Wait()
}
