package main

import (
	"fmt"
	"golang.org/x/time/rate"
	"time"
)

func main() {

	limiter := rate.NewLimiter(5, 2)
	for i := 0; i < 10; i++ {
		if limiter.Allow() {
			fmt.Println(i+1, " ok", time.Now().Unix())
		} else {
			fmt.Println("allow")
		}
		time.Sleep(time.Millisecond * 100)
	}

}
