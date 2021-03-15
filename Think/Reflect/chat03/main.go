package main

import (
	"fmt"
	"reflect"
)

func Add(a, b int) int { return a + b }

func main() {

	v := reflect.ValueOf(Add)
	if v.Kind() != reflect.Func {
		return
	}
	t := v.Type()
	args := make([]reflect.Value, t.NumIn())
	for i := range args {
		if t.In(i).Kind() != reflect.Int {
			return
		}
		args[i] = reflect.ValueOf(i + 10)
	}
	for _, r := range v.Call(args) {
		fmt.Println("r ==> ", r)
	}
}
