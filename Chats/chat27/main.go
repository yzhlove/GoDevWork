package main

import (
	"fmt"
	"time"
)

func main() {

	var a, b, c uint32
	a = 1
	b = 2
	c = 3

	fmt.Println("num => ", a&b&c)

	t1 := time.Now().Unix()
	time.Sleep(time.Second)
	t2 := time.Now().Unix()

	fmt.Println("points => ", (t2-t1)*100)

	fmt.Println("a => ",32 % 60 ," n =>" , 72 % 60)

}
