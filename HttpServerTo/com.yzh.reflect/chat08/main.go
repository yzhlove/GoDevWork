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

	t := T{100, "hello"}
	s := reflect.Indirect(reflect.ValueOf(&t))
	typeOf := s.Type()
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		fmt.Printf("%d:%s %s = %v \n", i, typeOf.Field(i).Name, f.Type(), f.Interface())
	}
	s.Field(0).Set(reflect.ValueOf(1000))
	s.Field(1).SetString("yayayayayayya")
	fmt.Println("now => ", t)

}
