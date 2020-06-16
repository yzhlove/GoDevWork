package main

import (
	"fmt"
	"reflect"
)

func main() {

	var x float64 = 7.4
	v := reflect.ValueOf(x)
	fmt.Println("is set:", v.CanSet())
	if v.CanSet() {
		v.SetFloat(7.1)
	} else {
		fmt.Println("not set")
	}

	var y float64 = 7.4
	vv := reflect.ValueOf(&y)
	fmt.Println("type:", vv.Type())
	fmt.Println("kind:", vv.Kind())
	//fmt.Println("value 1:", vv.Int())
	fmt.Println("value 2:", vv.String())
	fmt.Println("value 3:", vv.Pointer())
	fmt.Println("is set:", vv.CanSet())
	fmt.Println()
	p := vv.Elem()
	fmt.Println("type:", reflect.ValueOf(p).Type())
	fmt.Println("kind:", reflect.ValueOf(p).Kind())
	fmt.Println("is can:", reflect.ValueOf(p).CanSet())
	fmt.Println("p is can:", p.CanSet())
	p.SetFloat(7.8)
	fmt.Println("y => ", reflect.ValueOf(y), " --> ", y, " vv value:", vv.Pointer())
	fmt.Println("yi => ", p.Interface())

}
