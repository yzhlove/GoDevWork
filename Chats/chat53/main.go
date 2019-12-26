package main

import (
	"fmt"
	"math/rand"
)

var count int

func main() {

	for i, ok := getRand(); ok; {
		fmt.Println("i ===> ", i)
	}

}

func getRand() (i int, ok bool) {
	i = rand.Intn(10)
	fmt.Println("g i==> ", i)
	if count > 10 {
		ok = false
		return
	}
	count++
	ok = true
	return
}
