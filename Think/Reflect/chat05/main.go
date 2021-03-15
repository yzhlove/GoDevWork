package main

import (
	"fmt"
	"reflect"
)

type Stu struct {
	Name string
}

func (this *Stu) Fn(p1, p2 int) int {
	return p1 + p2
}

func (this Stu) Ft(p1, p2 int) int {
	return p1 - p2
}

func main() {
	s := &Stu{"Hank"}
	valueS := reflect.ValueOf(s)

	method := valueS.MethodByName("Fn")
	paramList := []reflect.Value{
		reflect.ValueOf(22),
		reflect.ValueOf(20),
	}
	resultList := method.Call(paramList)
	fmt.Println(resultList[0].Int()) // 42

	ms := reflect.TypeOf(s).Method(1)
	fmt.Println("ms.name => ", ms.Name)
	resultList = ms.Func.Call([]reflect.Value{
		reflect.ValueOf(&Stu{}),
		reflect.ValueOf(22),
		reflect.ValueOf(20),
	})
	fmt.Println(resultList[0].Int())

}
