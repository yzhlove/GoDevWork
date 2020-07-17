package main

import (
	"context"
	"fmt"
	"golang.org/x/time/rate"
	"time"
)

func main() {

	//r:每秒钟10个 b:桶容量为1
	limiter := rate.NewLimiter(2, 2)
	for i := 0; i < 10; i++ {
		if err := limiter.Wait(context.TODO()); err != nil {
			fmt.Println("Err:", err)
		} else {
			fmt.Println(i+1, " ok", time.Now().Unix())
		}
	}

}
