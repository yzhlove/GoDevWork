package main

import (
	"fmt"
	"reflect"
)

func main() {

	type User struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	user := User{"jack", 20}
	userType := reflect.TypeOf(user)
	firstField := userType.Field(0)
	fmt.Println(firstField.Name, firstField.Type, firstField.Tag, firstField.Index, firstField.PkgPath)

}
