package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"reflect"
)

func main() {

	buf := bytes.NewBuffer([]byte{})
	e := &A{T: 12138, S: "super", N: "to"}
	if err := JsonEncoder(buf, e); err != nil {
		panic(err)
	}

	var d1 = &A{}
	if err := JsonDecoder(buf, d1); err != nil {
		panic(err)
	}
	fmt.Println("buf -> ", buf.String())
	fmt.Println("A =", d1)

	if err := JsonEncoder(buf, e); err != nil {
		panic(err)
	}

	var d2 = A{}
	argv := reflect.New(reflect.Indirect(reflect.ValueOf(d2)).Type())
	fmt.Println("argv interface => ", reflect.TypeOf(argv), reflect.TypeOf(argv.Interface()))
	if err := JsonDecoder(buf, argv.Interface()); err != nil {
		panic(err)
	}
	fmt.Println("A2 =", argv, argv.Interface())

}

func JsonEncoder(writer io.Writer, data interface{}) error {
	encoder := json.NewEncoder(writer)
	return encoder.Encode(data)
}

func JsonDecoder(reader io.Reader, body interface{}) error {
	decoder := json.NewDecoder(reader)
	return decoder.Decode(body)
}

type A struct {
	N string
	S string
	T int32
}
