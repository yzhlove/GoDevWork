package main

import "fmt"

func main() {

	var a uint16 = 8

	b := a | (1 << 15)
	fmt.Printf("%16b\n", b)

	c := (1 << 15) &^ b
	fmt.Println("c => ", c, "a => ", a)
	fmt.Printf("%0.16b\n", c)

	fmt.Println(a &^ (1 << 15))

	var f uint16 = 128
	var s uint16 = 1 << 15
	var ff uint16 = f | s

	fmt.Println(f&s == s)
	fmt.Println(ff&s == s)

}
