package main

import (
	"fmt"
	"strings"
)

type Decoder interface {
	Parse(str string) string
}

type A struct{}

func (a *A) Parse(str string) string {
	return strings.ToLower(str)
}

type B struct{}

func (b *B) Parse(str string) string {
	return strings.ToUpper(str)
}

func Comp(x, y Decoder) {
	if x == y {
		fmt.Println("comp")
	} else {
		fmt.Println("not comp")
	}
}

func main() {

	a := new(A)
	b := new(B)
	Comp(a, b)

}
