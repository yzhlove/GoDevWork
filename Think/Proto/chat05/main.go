package main

import "fmt"

func main() {

	aa := &Aa{baseA: &baseA{N: 100}}
	bb := &Ab{baseA: &baseA{N: 200}}

	cc := &CA{a: aa}
	cc.ToShow()

	cc.a = bb
	cc.ToShow()

}

type IA interface {
	ToStr() string
}

type IC interface {
	ToShow()
}

type baseA struct {
	N int
}

func (b *baseA) ToStr() string {
	return fmt.Sprintf("baseA --> %d", b.N)
}

type Aa struct {
	*baseA
}

type Ab struct {
	*baseA
}

func (b *Ab) ToStr() string {
	return fmt.Sprintf("AAbb ---> %d", b.N*100)
}

type CA struct {
	a IA
}

func (c *CA) ToShow() {
	fmt.Println("show ==> ", c.a.ToStr())
}
