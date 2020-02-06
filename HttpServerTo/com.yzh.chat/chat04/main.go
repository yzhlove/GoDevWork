package main

import "fmt"

func main() {
	l := make([]string, 0)
	fmt.Printf("len %d cap %d ", len(l), cap(l))
	l = append(l, "aaaa")
	fmt.Printf("len %d cap %d ", len(l), cap(l))
	l = append(l, "aaaa")
	fmt.Printf("len %d cap %d ", len(l), cap(l))
	l = append(l, "aaaa")
	fmt.Printf("len %d cap %d ", len(l), cap(l))
}
