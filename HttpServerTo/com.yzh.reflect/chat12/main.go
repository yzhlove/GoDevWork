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
	newUser := reflect.New(userType)
	newUser.Elem().FieldByName("Name").SetString("what are you doing")
	reflect.Indirect(newUser).FieldByName("Age").SetInt(128)
	fmt.Println(user, newUser)

	student := User{"student", 20}
	newStudent := reflect.Indirect(reflect.New(reflect.ValueOf(student).Type()))
	newStudent.FieldByName("Name").SetString("sssssssss")
	newStudent.FieldByName("Age").SetInt(123456)
	fmt.Println(student, newStudent)

	//panic 创建的并非是User类型而是指针类型
	//teacher := &User{"teacher", 20}
	//newTeacher := reflect.Indirect(reflect.New(reflect.ValueOf(&teacher).Type()))
	//newTeacher.Elem().Field(0).SetString("ttttttttttt")
	//fmt.Println(teacher, newTeacher)
}
