package main

import (
	"fmt"
	"math"
)

func main() {

	var (
		a float64 = 1
		b float64 = 10
	)

	fmt.Printf("%v \n", a/b)
	fmt.Printf("%v \n", math.Ceil(a/b))

}
