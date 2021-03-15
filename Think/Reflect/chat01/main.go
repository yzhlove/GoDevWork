package main

import (
	"fmt"
	"reflect"
	"strings"
)

func main() {
	test()
}

func test() {
	var wg strings.Builder
	typ := reflect.TypeOf(&wg)
	for i := 0; i < typ.NumMethod(); i++ {
		method := typ.Method(i)
		args := make([]string, 0, method.Type.NumIn())
		returns := make([]string, 0, method.Type.NumOut())

		for j := 0; j < method.Type.NumIn(); j++ {
			args = append(args, method.Type.In(j).Name())
		}

		for k := 0; k < method.Type.NumOut(); k++ {
			returns = append(returns, method.Type.Out(k).Name())
		}

		fmt.Printf("func (%s) %s(%q) (%q) \n",
			typ.Elem().Name(), method.Name,
			strings.Join(args, ","), strings.Join(returns, ","))
	}
}
