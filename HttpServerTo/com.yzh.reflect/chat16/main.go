package main

import (
	"fmt"
	"reflect"
)

func main() {

	funcType := reflect.TypeOf(testFunc)
	fmt.Println("funcType => ", funcType)
	funcValue := reflect.ValueOf(testFunc)
	result := funcValue.Call([]reflect.Value{reflect.ValueOf(10)})

	for _, v := range result {
		if res, ok := v.Interface().(int); ok {
			fmt.Println("res => ", res)
		}
	}

	fmt.Println()

	funcValue2 := reflect.ValueOf(testFunc)
	if f, ok := funcValue2.Interface().(func(int) int); ok {
		res := f(5)
		fmt.Println("res => ", res)
	} else {
		panic("type err")
	}

}

func testFunc(count int) int {
	sum := 0
	for i := 0; i < count; i++ {
		sum += i
	}
	return sum
}
