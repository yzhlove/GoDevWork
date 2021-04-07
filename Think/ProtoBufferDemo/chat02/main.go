package main

import "fmt"

func main() {
	var a int32 = 2
	fmt.Printf("%b\n", a)

	var b = (a << 1) ^ (a >> 31)
	fmt.Printf("%b\n", b)

}
