package main

import (
	"fmt"
	"reflect"
)

func main() {

	intSlice := make([]int, 0)
	intMap := make(map[string]int)

	sliceType := reflect.TypeOf(intSlice)
	mapType := reflect.TypeOf(intMap)

	newSlice := reflect.MakeSlice(sliceType, 0, 0)
	newMap := reflect.MakeMapWithSize(mapType, 10)

	v := 10
	rv := reflect.ValueOf(v)
	newSlice = reflect.Append(newSlice, rv)
	realSlice := newSlice.Interface().([]int)
	fmt.Println(realSlice)

	k := "hello world"
	rk := reflect.ValueOf(k)
	newMap.SetMapIndex(rk, rv)
	realMap := newMap.Interface().(map[string]int)
	fmt.Println(realMap)
}
