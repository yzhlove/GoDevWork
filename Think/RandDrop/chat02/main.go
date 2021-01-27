package main

import (
	"fmt"
	"math/rand"
	"time"
)

//洗牌算法

func init() {
	rand.Seed(time.Now().UnixNano())
}

var rands = []int{1, 2, 3, 4, 5}

func main() {

	Shuffle1(rands)
	fmt.Println("rands ==> ", rands)

	Shuffle2(rands)
	fmt.Println("rands ==> ", rands)

}

func Shuffle1(trands []int) {
	for i := len(trands) - 1; i >= 0; i-- {
		loc := rand.Int() % (i + 1)
		trands[i], trands[loc] = trands[loc], trands[i]
	}
}

func Shuffle2(trands []int) {
	for i := 0; i < len(trands); i++ {
		loc := rand.Intn(len(trands) - i)
		trands[i], trands[loc] = trands[loc], trands[i]
	}
}
