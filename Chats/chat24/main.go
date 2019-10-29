package main

import (
	"fmt"
	"math"
)

func main() {

	fmt.Println("value = ", (0 % 100))

	var a []int
	a = append(a, 1, 2, 3, 4, 5)
	fmt.Println(a)

	b := 10
	if b > 3 {
		fmt.Println("a")
	} else if b > 4 {
		fmt.Println("b")
	} else {
		fmt.Println("c")
	}

	fmt.Println(math.MaxUint16 / 8 / 1000)

	fmt.Println("id => ", 2131%100/10)

	fmt.Println("point =>", 0/10)

}
