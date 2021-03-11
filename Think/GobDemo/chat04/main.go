package main

import (
	"bufio"
	"bytes"
	"encoding/gob"
	"fmt"
)

func main() {

	data := bytes.NewBuffer([]byte{})
	buf := bufio.NewWriter(data)
	a := &A{Var: "A", Str: "AAA", Seq: 100}
	b := &B{Acc: "Actor", Gcc: "China", Ecc: 1000, Dcc: 12138, Fcc: 214748}

	encoder := gob.NewEncoder(buf)
	encoder.Encode(a)
	encoder.Encode(b)

	buf.Flush()
	fmt.Println("size => ", data.Len())

	decoder := gob.NewDecoder(data)

	var tA = &A{}
	var tB = &B{}
	decoder.Decode(tA)
	decoder.Decode(tB)

	fmt.Println("tA => ", tA, " tB => ", tB)

}

type A struct {
	Var string
	Str string
	Seq int64
}

type B struct {
	Acc string
	Gcc string
	Ecc int64
	Dcc int32
	Fcc int32
}
