package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	a, b := genRand()
	SrandA(a)
	SrandB(b)
}

func SrandA(rd chan int) {
	trands := []int{1, 2, 3, 4, 5}
	for i := len(trands) - 1; i >= 0; i-- {
		r := <-rd
		loc := r % (i + 1)
		fmt.Printf("a loc=%d r=%d i=%d swpA=%d swpB=%d detail=%v \n", loc, i, i+1, trands[i], trands[loc], trands)
		trands[i], trands[loc] = trands[loc], trands[i]
	}
	fmt.Println("A trands =>", trands)
}

func SrandB(rd chan int) {
	trands := []int{1, 2, 3, 4, 5}
	for i := 0; i < len(trands); i++ {
		r := <-rd
		loc := r % (len(trands) - i)
		fmt.Printf("b loc=%d r=%d i=%d swpA=%d swpB=%d detail=%v \n", loc, i, len(trands)-1, trands[i], trands[loc], trands)
		trands[i], trands[loc] = trands[loc], trands[i]
	}
	fmt.Println("B trands =>", trands)
}

func genRand() (chan int, chan int) {
	a := make(chan int)
	b := make(chan int)
	go func() {
		for {
			n := rand.Int()
			go func(n int) { a <- n }(n)
			go func(n int) { b <- n }(n)
			time.Sleep(time.Second)
		}
	}()
	return a, b
}
