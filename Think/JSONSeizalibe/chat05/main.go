package main

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func main() {

	buffer := bytes.NewBuffer([]byte{})

	a := &A{Name: "aaa", Birthday: "1996-8080", Age: 12345}
	b := &B{Name: "bbb", Type: "string", Class: "B_CLASS", No: 123}

	encoder := json.NewEncoder(buffer)
	if err := encoder.Encode(a); err != nil {
		panic(err)
	}
	if err := encoder.Encode(b); err != nil {
		panic(err)
	}

	fmt.Println("string encoder ==> ", buffer.String())

	decoder := json.NewDecoder(buffer)
	newA := &A{}
	newB := &B{}
	if err := decoder.Decode(newA); err != nil {
		panic(err)
	}
	if err := decoder.Decode(newB); err != nil {
		panic(err)
	}

	fmt.Println(newA, newB)

}

type A struct {
	Name     string
	Birthday string
	Age      int
}

type B struct {
	Name  string
	Type  string
	Class string
	No    int
}
