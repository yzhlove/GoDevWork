package main

import (
	"fmt"
	"reflect"
)

func main() {

	var x float64 = 3.14
	v := reflect.ValueOf(x)

	fmt.Println("type:", v.Type())
	fmt.Println("kind:", v.Kind())
	fmt.Println("value:", v.Float())

	y := v.Interface().(float64)
	fmt.Println("y = ", y)
	fmt.Println()
	type MyInt int
	var zz MyInt = 7
	vv := reflect.ValueOf(zz)
	fmt.Println("type:", vv.Type())
	fmt.Println("kind:", vv.Kind())
	fmt.Println("value:", vv.Int())

	fmt.Println("tt=> ", reflect.TypeOf(vv.Interface()))
	res, ok := vv.Interface().(int)
	fmt.Println("res => ", res, " ok:", ok)
	ress := vv.Interface().(MyInt)
	fmt.Println("ress => ", ress)

}
