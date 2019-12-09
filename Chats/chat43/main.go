package main

import "fmt"

type IntA interface {
	GetNumber() int
}

type IntB interface {
	GetNumber() int
}

type c struct {
	n int
}

func (tc c) GetNumber() int {
	return tc.n
}

func main() {

	t := c{n: 1000}
	fmt.Printf("%T %v \n", t, t.GetNumber())
}
