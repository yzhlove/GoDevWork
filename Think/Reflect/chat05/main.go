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

	//TestCall()
	//
	//TestCall2()

	TestCall3()
}

func TestCall() {

	s := Stu{}
	rtyp := reflect.New(reflect.Indirect(reflect.ValueOf(&s)).Type())

	res := rtyp.Method(0).Call([]reflect.Value{reflect.ValueOf(100), reflect.ValueOf(200)})
	fmt.Println("res[0] = ", res[0])

	res = rtyp.Method(1).Call([]reflect.Value{reflect.ValueOf(100), reflect.ValueOf(200)})
	fmt.Println("res[0] = ", res[0])

	fmt.Println(reflect.TypeOf(s).Name())
	fmt.Println(reflect.TypeOf(s).Method(0).Name)

}

func TestCall2() {

	s := Stu{}

	rtyp := reflect.TypeOf(&s)

	fmt.Println("=================================")

	for i := 0; i < rtyp.NumMethod(); i++ {
		mtd := rtyp.Method(i)
		fmt.Println(mtd.Name)
		for m := 0; m < mtd.Type.NumIn(); m++ {
			fmt.Print("\t", mtd.Type.In(m))
		}
		fmt.Println()
		for n := 0; n < mtd.Type.NumOut(); n++ {
			fmt.Print("\t", mtd.Type.Out(n))
		}
		fmt.Println()
		fmt.Println("--------------------")
	}

	res := rtyp.Method(0).Func.Call(
		[]reflect.Value{
			reflect.New(reflect.Indirect(reflect.ValueOf(s)).Type()),
			reflect.ValueOf(1000),
			reflect.ValueOf(2000),
		})
	fmt.Println("res => ", res[0])

}

func TestCall3() {

	s := Stu{}
	rtyp := reflect.ValueOf(&s)
	fmt.Println("==================================")

	for i := 0; i < rtyp.NumMethod(); i++ {
		mtd := rtyp.Method(i)
		mtp := rtyp.Type().Method(i)
		fmt.Println("mtp -> ", mtp.Name)
		fmt.Println("mtd -> ", mtd.Type().Name())

		for m := 0; m < mtd.Type().NumIn(); m++ {
			fmt.Print("\t", mtd.Type().In(m))
		}
		fmt.Println()
		for n := 0; n < mtd.Type().NumOut(); n++ {
			fmt.Print("\t", mtd.Type().Out(n))
		}
		fmt.Println()
		fmt.Print("--------------------------------")

		res := mtd.Call([]reflect.Value{
			reflect.ValueOf(2000),
			reflect.ValueOf(1000),
		})
		fmt.Println("res[0] = ", res[0])
	}

}
