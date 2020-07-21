package main

import (
	"fmt"
	"math"
	"sync"
	"time"
)

//漏桶算法
type LeakyBucket struct {
	rate       float64 //固定每秒出水率
	capacity   float64 //桶的容量
	water      float64 //桶中当前的水量
	lastLeakMs int64   //桶上次漏水的时间戳 ms
	sync.Mutex
}

func (l *LeakyBucket) Allow() bool {
	l.Lock()
	defer l.Unlock()
	now := time.Now().UnixNano() / 1e6
	eclipse := float64(now-l.lastLeakMs) / 1000 * l.rate
	fmt.Print(" ===> water -> ", l.water, " eclipse => ", eclipse, "\t")
	l.water -= eclipse
	l.water = math.Max(0, l.water)
	l.lastLeakMs = now
	if l.water+1 < l.capacity {
		l.water++
		return true
	}
	return false
}

func (l *LeakyBucket) Set(r, c float64) {
	l.rate = r
	l.capacity = c
	l.lastLeakMs = time.Now().UnixNano() / 1e6
}

func main() {

	var wg sync.WaitGroup
	var rate LeakyBucket
	//每秒访问限制3哥请求，桶容量为3个
	rate.Set(3, 3)
	for i := 1; i <= 10; i++ {
		wg.Add(1)
		fmt.Print("Create Req:", i, time.Now().Format(time.RFC3339))
		go func(i int) {
			defer wg.Done()
			if rate.Allow() {
				fmt.Println(" Resp Succeed.", i, time.Now().Format(time.RFC3339))
			} else {
				fmt.Println(" Resp Failed .")
			}
		}(i)
		time.Sleep(100 * time.Millisecond)
	}
	wg.Wait()
}
