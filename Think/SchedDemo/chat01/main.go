package main

import (
	"runtime"
	"time"
)

func main() {

	runtime.GOMAXPROCS(1)
	go ok()
	time.Sleep(time.Second)

}

func ok() {
	i := 1
	for {
		i++
	}
}
