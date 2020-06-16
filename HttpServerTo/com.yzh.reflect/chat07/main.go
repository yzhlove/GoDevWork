package main

import (
	"fmt"
	"reflect"
)

func main() {

	type T struct {
		A int
		B string
	}

	t := T{23, "skiddo"}
	s := reflect.ValueOf(&t).Elem()
	typeOf := s.Type()
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		fmt.Printf("%d:%s %s = %v \n", i, typeOf.Field(i).Name, f.Type(), f.Interface())
	}
	s.Field(0).SetInt(100)
	s.Field(1).SetString("hello world")
	fmt.Println("now t => ", t)
}
