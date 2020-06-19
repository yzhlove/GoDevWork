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

	userptrs := []*User{&User{Name: "helllo"}, &User{Name: "World"}}
	users := []User{{Name: "hello"}, {Name: "heiheihei"}}
	userptrSlice := reflect.ValueOf(userptrs)
	userSlice := reflect.ValueOf(users)
	fmt.Println(userptrSlice.Type(), " - ", userSlice.Type())

	realptrSlice := reflect.Indirect(userptrSlice)
	realSlice := reflect.Indirect(userSlice)
	fmt.Println(realptrSlice.Type(), " - ", realSlice.Type())

	elementptrSlice := realptrSlice.Type().Elem()
	elementSlice := realSlice.Type().Elem()
	fmt.Println(reflect.TypeOf(elementptrSlice), " - ", reflect.TypeOf(elementSlice))

	newptrSlice := reflect.New(elementptrSlice)
	newSlice := reflect.New(elementSlice)
	fmt.Println(reflect.TypeOf(newptrSlice), " - ", reflect.TypeOf(newSlice))

	resultptrSlice := newptrSlice.Elem()
	resultSlice := newSlice.Elem()
	fmt.Println(resultptrSlice.Type(), " - ", resultSlice.Type())

	interfaceptrSlice := resultptrSlice.Interface()
	interfaceSlic := resultSlice.Interface()
	fmt.Println(interfaceptrSlice, " - ", interfaceSlic)

}
