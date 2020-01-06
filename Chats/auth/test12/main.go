package main

import (
	"fmt"
	"time"
)

func main() {

	t1 := time.Now()

	time.Sleep(time.Millisecond * 500)

	t2 := time.Now()

	fmt.Println(t2.Sub(t1).Nanoseconds()/1e6, ".ms")

}
