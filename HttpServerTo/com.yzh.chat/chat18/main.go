package main

import "fmt"

func main() {

	var str string = "我和我的祖国"

	for i, v := range str {
		fmt.Printf("i %d v %q \n", i, v)
	}

	fmt.Println(`\\\\\\\\\\\\\\\\\`)

	for i, v := range []rune(str) {
		fmt.Printf("i %d v %q \n", i, v)
	}

}
