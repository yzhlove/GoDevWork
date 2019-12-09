package main

import (
	"fmt"
	"os"
	"runtime/trace"
	"sync"
	"time"
)

//RLock

var lock sync.RWMutex
var wg sync.WaitGroup

func main() {
	_ = trace.Start(os.Stderr)
	defer trace.Stop()
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go gets()
	}
	wg.Wait()
}

func gets() {
	for i := 0; i < 100000; i++ {
		get(i)
	}
	wg.Done()
}

func get(_ int) {
	begin := time.Now()
	lock.RLock()
	defer lock.RUnlock()
	t := time.Since(begin).Nanoseconds() / 1000000
	if t > 100 { //如果大于100ms
		fmt.Println("fuck hear")
	}
}
