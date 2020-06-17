package main

import (
	"fmt"
	"reflect"
)

func main() {

	type User struct {
		Name string
		Age  int
	}

	user := User{"jack", 20}
	userType := reflect.TypeOf(user)
	fmt.Println(userType.Name(), " ", userType.Kind())

}
