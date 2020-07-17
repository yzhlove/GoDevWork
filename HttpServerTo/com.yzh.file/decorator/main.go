package main

import "fmt"

type decoratorFunc func(s string)

func decorator(f decoratorFunc) decoratorFunc {
	d := func(s string) {
		fmt.Println("start")
		f(s)
		fmt.Println("end")
	}
	return d
}

func Hello(s string) {
	fmt.Println("hello string -> ", s)
}

func main() {
	decorator(Hello)("hi")
}
