package main

import "fmt"

func main() {

	var up float64 = 0.3

	var b uint32 = 100

	fmt.Println(float64(b) * (1 + up))
	fmt.Println(b + uint32(float64(b)*up))

	aa := map[int]int{1: 10, 2: 20, 3: 30}

	t := aa[4]

	fmt.Println("t = ", t)

}
