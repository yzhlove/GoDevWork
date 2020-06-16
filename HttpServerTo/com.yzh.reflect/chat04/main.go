package main

import (
	"fmt"
	"reflect"
)

func main() {

	var x uint8 = 'x'
	v := reflect.ValueOf(x)
	fmt.Println("type:", v.Type())
	fmt.Println("kind:", v.Kind(), " ", v.Kind() == reflect.Uint8)
	fmt.Println("value:", reflect.ValueOf(x))
	fmt.Println("return ", reflect.TypeOf(v.Uint()), " ", v.Uint())

	type MyInt int
	var xx MyInt = 7
	vv := reflect.ValueOf(xx)
	fmt.Println(vv, " type:", reflect.TypeOf(vv))

}
