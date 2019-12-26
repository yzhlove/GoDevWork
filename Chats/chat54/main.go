package main

import "fmt"

func main() {

	var i int
Fuck:
	for i = 0; i < 10; i++ {
		switch {
		case i > 5:
			break Fuck
		}
	}
	fmt.Println(" i +> ", i)
}
