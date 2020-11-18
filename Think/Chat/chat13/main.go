package main

import "fmt"

func main() {

	a := AList{A{Id: 10, Lv: 100}}
	fmt.Println(a)
	a.Update()
	fmt.Println(a)

}

type A struct {
	Id int
	Lv int
}

type AList []A

type Manager struct {
	As AList
}

func (a *AList) Update() {
	c := []A{{Id: 1, Lv: 1}, {Id: 2, Lv: 2}}
	*a = c
}
