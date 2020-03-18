package main

import "fmt"

func main() {

	fmt.Println("number", uint64(1)<<63)

	fmt.Printf("%b \n", uint64(1)<<63)
	fmt.Printf("%9b \n", uint64(1<<64-1))

	fmt.Printf("%b \n", uint64(1)<<11-1)

}
