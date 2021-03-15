package main

import (
	"fmt"
	"reflect"
)

//判断是否实现接口类型

type CustomError struct{}

func (c *CustomError) Error() string {
	return reflect.TypeOf(c).Name()
}

func main() {

	typeOfErr := reflect.TypeOf((*error)(nil)).Elem()
	customErrPtr := reflect.TypeOf(&CustomError{})
	cutomErr := reflect.TypeOf(CustomError{})

	fmt.Println(customErrPtr.Implements(typeOfErr))
	fmt.Println(cutomErr.Implements(typeOfErr))

}
