package main

import "fmt"

func main() {

	get()
	//set()

}

func get() {

	//1101100100
	var a uint64 = 868
	a = a
	for i := 0; i < 64; i++ {
		//fmt.Printf("%d -> %b \n", i, uint64(1)<<uint(i))
		if a&(uint64(1)<<uint(i)) != 0 {
			fmt.Printf("%d ", i)
		}
	}
}

func set() {

	var a uint64 = 868
	a = a

	b := a | uint64(1)<<3
	fmt.Printf("%b \n", b)

}
