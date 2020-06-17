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
	userValue := reflect.ValueOf(user)
	userInterface := userValue.Interface()

	if u, ok := userInterface.(User); ok {
		fmt.Printf("User: name: %s age:%d \n", u.Name, u.Age)
	} else {
		panic("type err")
	}
}
