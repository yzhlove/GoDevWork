package main

import (
	"fmt"
	"reflect"
)

func main() {

	var a int

	r := reflect.New(reflect.TypeOf(a))
	fmt.Println(r.Type(), r.Elem())

	rp := reflect.New(reflect.TypeOf(&a))
	fmt.Println(rp.Type(), rp.Elem())

	fmt.Println("=====================================")
	fmt.Println(reflect.Indirect(reflect.ValueOf(&a)).Type())
	fmt.Println(reflect.Indirect(reflect.ValueOf(a)).Type())

	rts := reflect.New(reflect.Indirect(reflect.ValueOf(a)).Type())
	fmt.Println(rts.Type(), rts.Elem())

	rtss := reflect.New(reflect.TypeOf(&a).Elem())
	fmt.Println(rtss.Type(), rtss.Elem())

}
