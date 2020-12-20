package main

type Cond interface {
	Id() uint32
	Update()
}
