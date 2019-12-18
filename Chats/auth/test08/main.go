package main

import (
	"flag"
	"fmt"
)

func main() {

	var a string
	a = *flag.String("abc", "12345", "")
	flag.Parse()

	if a != "null" && a != "" {
		fmt.Println(a)
	}

}
