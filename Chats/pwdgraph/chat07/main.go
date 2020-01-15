package main

import (
	"fmt"
	"math/rand"
)

func main() {

	var reward uint64 = 55
	var code uint64
	for i := 0; i < 7; i++ {
		key := uint64(rand.Int63() % 32)
		code |= key << (5 * i)
	}

	fmt.Println("reward => ", reward<<35)

	code |= reward << (5 * 7)
	fmt.Println(code)

	fmt.Println("code => ", code>>35)

}
