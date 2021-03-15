package main

import (
	"fmt"
	"reflect"
)

func main() {

	var err error

	etyp := reflect.TypeOf(err)

	fmt.Println(etyp)

	var x int
	fmt.Println(reflect.TypeOf(x))

	ntyp := reflect.TypeOf((*error)(nil))

	fmt.Println(ntyp, ntyp.Name())

	ntype := ntyp.Elem()
	fmt.Println(ntype, ntype.Name())

	eetyp := reflect.TypeOf((*error)(nil))
	fmt.Println(eetyp, eetyp.Name())
}
