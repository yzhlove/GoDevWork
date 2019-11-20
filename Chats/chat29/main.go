package main

import (
	"fmt"
	"math/rand"
)

func main() {

	var a []int

	for len(a) <= 5 {
		a = append(a, rand.Int())
	}

	fmt.Println("ok", " a =>", a)

}
