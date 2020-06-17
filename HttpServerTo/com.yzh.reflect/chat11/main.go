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

	user1 := User{"jack", 20}
	user1Value := reflect.ValueOf(&user1).Elem()
	user1Value.Field(0).SetString("jet Li")
	fmt.Println(user1)

	user2 := User{"jack", 20}
	user2Value := reflect.Indirect(reflect.ValueOf(&user2))
	user2Value.Field(0).Set(reflect.ValueOf("hello world"))
	fmt.Println(user2)

	user3 := User{"jack", 20}
	user3Value := reflect.Indirect(reflect.ValueOf(&user3))
	user3Value.FieldByName("Name").SetString("what are you doing")
	fmt.Println(user3)

	user4 := User{"jack", 20}
	user4Value := reflect.ValueOf(&user4).Elem()
	user4Value.FieldByName("Name").SetString("how are you ")
	fmt.Println(user4)

}
