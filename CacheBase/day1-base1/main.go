package main

import (
	"container/list"
	"fmt"
)

////////////////////////////
//golang list
////////////////////////////

func main() {
	//
	//l := list.New()
	//for i := 0; i < 10; i++ {
	//	l.PushBack(i)
	//}
	//show(l)
	//l.Init()
	//
	//for i := 0; i < 10; i++ {
	//	l.PushFront(i)
	//}
	//
	//show(l)
	//
	//fmt.Println("font => ", l.Front().Value)
	//fmt.Println("back => ", l.Back().Value)
	//
	//l.InsertAfter(123, l.Front())
	//
	//show(l)
	//
	//l.InsertBefore(12345, l.Front())
	//
	//show(l)

	cc()
}

func cc() {

	l := list.New()
	e := l.PushFront(5)
	l.InsertAfter(4, e)
	fmt.Println(l.Len())
	e2 := &list.Element{}
	t := l.InsertAfter(3, e2)
	fmt.Println(l.Len(), " - ", t)

}

func show(l *list.List) {
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Printf("%v ", e.Value)
	}
	fmt.Println()
}
