package main

import (
	"fmt"
	"golang.org/x/time/rate"
	"time"
)

func main() {

	limiter := rate.NewLimiter(5, 1)
	for i := 0; i < 20; i++ {
		r := limiter.Reserve()
		if r.OK() {
			fmt.Println("limiter ")
			fmt.Println(r.Delay().Milliseconds())
			time.Sleep(r.Delay())
		}
	}

}
