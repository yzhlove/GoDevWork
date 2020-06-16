package main

import (
	"fmt"
	"reflect"
)

func main() {

	var x float64 = 3.4
	fmt.Println("type:", reflect.TypeOf(x))
	fmt.Println("value:", reflect.ValueOf(x))

	var xx float64 = 3.14
	v := reflect.ValueOf(xx)
	fmt.Println("type:", v.Type())
	fmt.Println("kind:", v.Kind(), " is float64:", v.Kind() == reflect.Float64)
	fmt.Println("value:", v.Float(), " string:", v.String())

	type MyInt int
	var y MyInt = 16
	tp := reflect.ValueOf(y)
	fmt.Println(reflect.TypeOf(y))
	fmt.Println("type:", tp.Type())
	fmt.Println("kind:", tp.Kind(), " is int:", tp.Kind() == reflect.Int)
	fmt.Println("value:", tp.Int(), " string:", tp.String())
}
