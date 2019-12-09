package main

import (
	"fmt"
	"os"
)

func main() {

	path := "yzh/love/"

	if err := os.MkdirAll(path, 0755); err != nil {
		panic(err)
	}

	fmt.Println("create ok.")

}
