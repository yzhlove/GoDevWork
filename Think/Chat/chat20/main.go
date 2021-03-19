package main

import "fmt"

func main() {
	fmt.Println(rt())
	fmt.Println(rs())
}

func rt() (a int) {

	a = 10
	defer func() {
		a = 20
		fmt.Println("a = ", a)
	}()
	a = 15
	return
}

func rs() int {
	var a = 10
	defer func() {
		a = 100
		fmt.Println("a =>", a)
	}()
	a = 9
	return a
}
