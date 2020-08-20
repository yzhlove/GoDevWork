package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {

	runtime.GOMAXPROCS(1)

	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println(" i = ", i)
		}
	}()
	for i := 0; i < 3; i++ {
		runtime.Gosched()
		fmt.Println("main i = ", i)
	}
	time.Sleep(time.Second)
}
