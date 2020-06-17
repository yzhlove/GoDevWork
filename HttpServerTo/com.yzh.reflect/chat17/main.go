package main

import (
	"fmt"
	"reflect"
)

func main() {

	structValue := reflect.Indirect(reflect.ValueOf(testMakeStruct(0, false, "")))
	for i := 0; i < structValue.NumField(); i++ {
		fmt.Println(structValue.Field(i).Type())
		fmt.Println(structValue.Field(i).Interface())
	}
	var values = []interface{}{1234, false, "what are you doing."}
	for i := 0; i < structValue.NumField(); i++ {
		structValue.Field(i).Set(reflect.ValueOf(values[i]))
	}

	fmt.Println(structValue)

}

func testMakeStruct(args ...interface{}) interface{} {

	var structList []reflect.StructField
	for index, value := range args {
		argType := reflect.TypeOf(value)
		item := reflect.StructField{
			Name: fmt.Sprintf("Item%d", index),
			Type: argType,
		}
		structList = append(structList, item)
	}
	structType := reflect.StructOf(structList)
	structValue := reflect.New(structType)
	return structValue.Interface()
}
