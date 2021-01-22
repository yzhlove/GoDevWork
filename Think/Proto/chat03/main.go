package main

import (
	"fmt"
	"strconv"
)

func main() {
	var toer Toer
	toer = &C{B: &B{A: &A{N: 100}}}
	fmt.Println(toer.Too())

}

type Toer interface {
	Too() string
}

type A struct {
	N int
}

type B struct {
	*A
}

type C struct {
	*B
}

func (a *A) Too() string {
	return "A:" + strconv.Itoa(a.N)
}

func (c *C) Too() string {
	return "C:" + strconv.Itoa(c.N)
}
