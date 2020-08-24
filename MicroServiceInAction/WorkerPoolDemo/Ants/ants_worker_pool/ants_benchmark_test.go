package ants_worker_pool

import "time"

const (
	RunTimes          = 1e6
	BenchParam        = 10
	BenchAntsSize     = 1e5 * 2
	DefaultExpireTime = 10 * time.Second
)

func demoFunc() {
	time.Sleep(time.Duration(BenchParam) * time.Millisecond)
}

func demoPoolFunc(args interface{}) {
	n := args.(int)
	time.Sleep(time.Duration(n) * time.Millisecond)
}
