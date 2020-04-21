package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	status := make(chan int)
	go func() {
		for {
			status <- rand.Int()
			time.Sleep(time.Second)
		}
	}()

	for {
		select {
		case i, ok := <-status:
			fmt.Printf("[%t] %d \n", ok, i)
			break
		}
	}

}
