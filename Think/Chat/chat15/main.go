package main

import (
	"fmt"
	"time"
)

func main() {

	t := time.Now().Add(time.Millisecond * 86400)
	fmt.Println(t.Sub(time.Now()))
}
