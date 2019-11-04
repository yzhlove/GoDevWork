package main

import "fmt"

func main() {

	a := make([]int, 0, 8)
	fmt.Printf("a = %p \n", a)
	fmt.Printf("a %v %v \n", len(a), cap(a))
	test(a)

	fmt.Printf("a %v %v \n", len(a), cap(a))

	a = append(a, 5, 6)
	
	fmt.Printf("a = %v \n", a)
	fmt.Printf("a %v %v \n", len(a), cap(a))
}

func test(b []int) {
	fmt.Printf("b = %p \n", b)
	b = append(b, 0, 1, 2, 3, 4)
	fmt.Printf("b %v %v \n", len(b), cap(b))
	fmt.Printf("b = %v \n", b)
}
