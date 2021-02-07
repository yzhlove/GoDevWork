package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

func main() {

	a := &A{ID: 123, Name: "AA"}
	b := &B{ID: 1234, Title: "What", Remark: "aaaaaaa", City: "China Shanghai"}

	var buf bytes.Buffer
	if encoder := gob.NewEncoder(&buf); encoder != nil {
		if err := encoder.Encode(a); err != nil {
			panic(err)
		}
		if err := encoder.Encode(b); err != nil {
			panic(err)
		}
	}

	if decoder := gob.NewDecoder(&buf); decoder != nil {
		var da = &A{}
		if err := decoder.Decode(da); err != nil {
			panic(err)
		}
		fmt.Println("decoder A => ", da)

		var db = &B{}
		if err := decoder.Decode(db); err != nil {
			panic(err)
		}
		fmt.Println("decoder B => ", db)
	}

}

type A struct {
	ID   uint32
	Name string
}

type B struct {
	ID     uint32
	Title  string
	Remark string
	City   string
}
