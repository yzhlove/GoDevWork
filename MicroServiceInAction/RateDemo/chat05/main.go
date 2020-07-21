package main

import (
	"fmt"
	"sync"
	"time"
)

type LimitRate struct {
	rate  int           //计时器周期内最大允许的请求数
	begin time.Time     //计数开始时间
	cycle time.Duration //计数周期
	count int           //计数周期内累计收到的请求数
	sync.Mutex
}

func (l *LimitRate) Allow() bool {
	l.Lock()
	defer l.Unlock()
	if l.count == l.rate-1 {
		now := time.Now()
		if now.Sub(l.begin) >= l.cycle {
			l.Reset(now)
			return true
		}
		return false
	}
	l.count++
	return true
}

func (l *LimitRate) Allow2() bool {
	l.Lock()
	defer l.Unlock()
	if l.count == l.rate-1 {
		now := time.Now()
		if now.Sub(l.begin) > l.cycle {
			l.Reset(now)
			return true
		} else {
			time.Sleep(10 * time.Millisecond)
		}
	}
	l.count++
	return true
}

func (l *LimitRate) Set(r int, cyc time.Duration) {
	l.rate = r
	l.begin = time.Now()
	l.cycle = cyc
	l.count = 0
}

func (l *LimitRate) Reset(now time.Time) {
	l.begin = now
	l.count = 0
}

func main() {

	var wg sync.WaitGroup
	var rate LimitRate
	rate.Set(3, time.Second) //一秒内最多三次请求

	for i := 0; i < 10; i++ {
		wg.Add(1)
		fmt.Print("Create Req:", time.Now().Format(time.RFC3339))
		go func(i int) {
			defer wg.Done()
			//rate.Allow()
			if rate.Allow2() {
				fmt.Println(" ✓Response Req Succeed:", i, time.Now().Format(time.RFC3339))
			} else {
				fmt.Println(" ✗")
			}
		}(i)
		time.Sleep(200 * time.Millisecond)
	}
	wg.Wait()
}
