package main

import (
	"fmt"
	"runtime"
	"time"
)

var quit = make(chan struct{}, 1)

func loop() {
	for i := 0; i < 100; i++ {
		fmt.Println("i = ", i)
		if i%10 == 0 {
			time.Sleep(time.Second)
		}
	}
	quit <- struct{}{}
}

func main() {

	runtime.GOMAXPROCS(1)

	go loop()
	go loop()

	for i := 0; i < 2; i++ {
		<-quit
	}

}
