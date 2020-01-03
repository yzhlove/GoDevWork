package main

import (
	"fmt"
	"stathat.com/c/consistent"
)

func main() {

	//一致性哈希
	hash := consistent.New()
	hash.Add("A")
	hash.Add("B")

	for _, s := range []string{"1", "2", "3", "4", "5", "6", "7"} {
		fmt.Println(hash.Get(s))
	}

	fmt.Println()

	hash.Add("C")
	hash.Add("D")

	for _, s := range []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"} {
		fmt.Println(hash.Get(s))
	}

	fmt.Println()

	hash.Remove("A")
	hash.Remove("C")

	for _, s := range []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"} {
		fmt.Println(hash.Get(s))
	}

}
