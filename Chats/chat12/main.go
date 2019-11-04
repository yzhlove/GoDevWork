package main

import (
	"fmt"
	"math/rand"
)

func main() {

	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Println(rand.Perm(len(a)))

}
