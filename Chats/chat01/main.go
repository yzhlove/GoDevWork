package main

import "fmt"

func main() {

	var status uint8 = 0

	status |= 0x1

	fmt.Printf("%d %x \n",status,status)

	status |= 0x2

	fmt.Printf("%d %x \n",status,status)

	if status == 0x3 {
		fmt.Println("hhh")
	}
}
