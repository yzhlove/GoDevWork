package main

import "fmt"

func main() {
	c := &C{B: &B{A: &A{N: 100}}}
	c.Show()
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

func (a *A) To() {
	fmt.Println("A:to")
}

func (a *A) Show() {
	a.To()
	fmt.Printf("show....")
}

func (a *B) To() {
	fmt.Println("A:to")
}

func (a *C) To() {
	fmt.Println("A:to")
}
