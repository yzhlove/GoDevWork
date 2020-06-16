package main

import (
	"fmt"
	"io"
	"os"
	"reflect"
)

func main() {

	var r io.Reader
	tty, err := os.OpenFile("/dev/tty", os.O_RDWR, 0)
	if err != nil {
		panic(err)
	}
	fmt.Println(reflect.TypeOf(tty))
	fmt.Println(reflect.TypeOf(r))
	r = tty
	fmt.Println(reflect.TypeOf(r))

	var w io.Writer
	fmt.Println(reflect.TypeOf(w))
	w = r.(io.Writer)
	fmt.Println(reflect.TypeOf(w))

	var empty interface{}
	fmt.Println(reflect.TypeOf(empty))
	empty = w
	fmt.Println(reflect.TypeOf(empty))
}
