package main

import "fmt"

/*
	成就系统设计
*/

type Event interface {
	Notify()
}

type Watcher interface {
	Update()
}

type Events struct {
	ws []Watcher
}

func (es *Events) Add(w Watcher) {
	es.ws = append(es.ws, w)
}

func (es *Events) Notify() {
	for _, w := range es.ws {
		w.Update()
	}
}

type A struct{}

func (a A) Update() {
	fmt.Println("a.Update")
}

type B struct{}

func (b B) Update() {
	fmt.Print("b.Update")
}

func main() {

	es := &Events{}
	es.Add(A{})
	es.Add(B{})
	es.Notify()
	
}
