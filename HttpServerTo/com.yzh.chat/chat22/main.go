package main

import "fmt"

func main() {

	var a = map[string]string{
		"A": "1",
		"B": "2",
		"C": "3",
	}

	delete(a, "A")
	delete(a, "E")
	a = nil
	delete(a, "B")
	fmt.Println(a)

}
