package main

import "fmt"

func main() {

	test()
}

func test() {

	fmt.Println("one")
	defer fmt.Println("1")
	{
		defer fmt.Println("2")
		fmt.Println("two")
	}
	fmt.Println("three")
	defer fmt.Println("3")

	fmt.Println("message ==> ", 1<<12)
}
