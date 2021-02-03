package main

import (
	"fmt"
	"strconv"
)

type toer interface {
	to() string
}

type A struct {
	N int
	S string
}

type B struct {
	*A
}

type C struct {
	*B
}

func (a *A) to() string {
	t := "A:" + strconv.Itoa(a.N)
	fmt.Println(t)
	a.too()
	return t
}

func (c *C) to() string {
	t := "C:" + strconv.Itoa(c.N)
	fmt.Println(t)
	c.too()
	return t
}

func (a *A) too() {
	fmt.Println("B:", a.S)
}

func (b *B) too() {
	fmt.Println("B:", b.S)
}

//func (c *C) too() {
//	fmt.Println("C:", c.S)
//}

func main() {
	var tt toer
	tt = &C{B: &B{A: &A{S: "hello", N: 123}}}
	fmt.Println("tto => ", tt.to())
}
