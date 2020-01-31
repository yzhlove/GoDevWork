package main

import "fmt"

func main() {

	key := []byte{0xa1, 0xb1}
	value := string(key)
	fmt.Println(key, " ", value, " ", []byte(value))
}
