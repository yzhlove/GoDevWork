package main

import "fmt"

func main() {

	var s []int
	for i := 0; i < 10; i++ {
		s = append(s, i)
	}
	fmt.Printf("%T %v %T %v \n", s[:], s[:], s, s)
}

