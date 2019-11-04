package main

import "fmt"

func main() {

	a := 1234365

	fmt.Println(a/1000, a%1000)

	var b int = 1e6

	fmt.Printf("number = %d \n", b)

	fmt.Println("number = ", 0%1000, " - ", 0/1000)

}
