package main

import (
	"fmt"
	"strconv"
)

func main() {
	var inter Inter
	inter = &C{B: &B{A: &A{N: 100}}}
	fmt.Println(inter.ToStr())
	/*
		output:
		A:trans 为什么不是 C:trans???
		B:100
	*/
}

type Inter interface {
	ToStr() string
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

func (a *A) trans() {
	fmt.Println("A:trans")
}

func (b *B) ToStr() string {
	t := "B:" + strconv.Itoa(b.N)
	b.trans()
	return t
}

func (c *C) trans() {
	fmt.Println("C:trans")
}
