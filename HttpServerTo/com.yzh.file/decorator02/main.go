package main

import (
	"fmt"
	"reflect"
	"runtime"
	"time"
)

type SumFunc func(start int64, end int64) int64

func getFuncName(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}

func timeSumFunc(f SumFunc) SumFunc {
	return func(start int64, end int64) int64 {
		defer func(t time.Time) {
			fmt.Println("---- Time End ----", getFuncName(f), time.Since(t))
		}(time.Now())
		return f(start, end)
	}
}

func main() {
	a := timeSumFunc(Sum1)(-1000, 1e6)
	b := timeSumFunc(Sum2)(-1e6, 1e6)
	fmt.Println(a, b)
}

func Sum1(start, end int64) int64 {
	var sum int64
	if start > end {
		start, end = end, start
	}
	for i := start; i < end; i++ {
		sum += i
	}
	return sum
}

func Sum2(start, end int64) int64 {
	if start > end {
		start, end = end, start
	}
	return (end - start + 1) * (end + start) / 2
}
