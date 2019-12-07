package main

import "fmt"

type inter interface {
	Get() int
}

type A struct {
	N int
}

func (a A) Get() int {
	return a.N
}

func main() {

	var i inter = &A{N: 100}
	fmt.Printf("%T %v \n", i, i.Get())

	a := A{N: 200}
	var it inter
	it = a
	fmt.Printf("%T %v \n", it, it)
}
