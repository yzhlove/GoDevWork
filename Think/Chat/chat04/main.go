package main

import "fmt"

func main() {

	a := ArrayList{}
	a.Push(1)
	a.Push(2)
	a.Push(3)
	a.Push(4)
	fmt.Println(a)

}

type ArrayList []int

func (a ArrayList) Pop() int {
	x := a[len(a)-1]
	a = a[:len(a)-1]
	return x
}

func (a ArrayList) Push(x int) {
	a = append(a, x)
}
