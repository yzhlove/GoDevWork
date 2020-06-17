package main

import (
	"fmt"
	"reflect"
	"time"
)

func main() {

	funcType := reflect.TypeOf(testMakeFunc)
	fmt.Println("type => ", funcType)
	funcValue := reflect.ValueOf(testMakeFunc)

	newFunc := reflect.MakeFunc(funcType, func(args []reflect.Value) (results []reflect.Value) {
		start := time.Now()
		out := funcValue.Call(args)
		end := time.Now()
		fmt.Println(end.Sub(start))
		return out
	})

	var count int = 4
	result := newFunc.Call([]reflect.Value{reflect.ValueOf(count)})
	for _, value := range result {
		fmt.Println(value)
	}
}

func testMakeFunc(count int) int {

	sum := 0
	for i := 0; i < count; i++ {
		sum += i
	}
	return sum
}
