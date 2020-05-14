package main

import "fmt"

func main() {

	fmt.Println("A -> ", Get("A"))
	fmt.Println("P -> ", Get("P"))

	fmt.Println("A -> ", GetTeacher("A"))
	fmt.Println("P -> ", GetTeacher("P"))
}

var dict = map[string]string{
	"A": "apple",
	"B": "Bus",
	"C": "City",
}

type Teacher struct {
	Name    string
	Project string
}

var School = map[string]*Teacher{
	"A": &Teacher{Name: "apple", Project: "English"},
	"B": &Teacher{Name: "bob", Project: "Math"},
}

func Get(key string) string {
	return dict[key]
}

func GetTeacher(key string) *Teacher {
	return School[key]
}
