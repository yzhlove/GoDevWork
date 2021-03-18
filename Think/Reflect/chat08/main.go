package main

import (
	"fmt"
	"go/ast"
	"go/token"
	"reflect"
)

// reflect elem or interface

func main() {

	var str string = "123456"

	fmt.Println("ast ==> ", ast.IsExported(reflect.TypeOf(str).Name()))
	fmt.Println("token ==> ", token.IsExported(reflect.TypeOf(str).Name()))
	fmt.Println("is ==> ", isExportedOrBuiltinType(reflect.TypeOf(str)))

	fmt.Println("==========================================")

	vt := reflect.ValueOf(&str)

	fmt.Println(vt.Elem())
	fmt.Println(reflect.Indirect(vt).Interface())

	test()
	testMap()
	tttt()
	testSet()

}

func isExportedOrBuiltinType(t reflect.Type) bool {
	fmt.Println("t.Kind => ", t.Kind())
	for t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	fmt.Println(token.IsExported(t.Name()))
	fmt.Println(t.PkgPath() == "")
	// PkgPath will be non-empty even for an exported type,
	// so we need to check the type name as well.
	return token.IsExported(t.Name()) || t.PkgPath() == ""
}

func test() {

	var ss = "1314"

	vt := reflect.ValueOf(&ss)

	fmt.Println(vt.Elem().Kind())

	rep := reflect.New(reflect.TypeOf(&ss).Elem())
	fmt.Println("rep =>", rep.Kind())
	fmt.Println(reflect.Indirect(rep).Kind())

}

func testMap() {

	var array = make([]string, 0, 0)

	view := reflect.New(reflect.TypeOf(&array).Elem())
	fmt.Println("view kind => ", reflect.Indirect(view).Kind())
	if view.Elem().CanSet() {
		intS := make([]int, 0)
		sas := reflect.MakeSlice(reflect.TypeOf(intS), 0, 0)
		sas = reflect.Append(sas, reflect.ValueOf(128))
		fmt.Println(sas)

		intsS := make([]int, 0)
		sass := reflect.MakeSlice(reflect.TypeOf(intsS), 0, 0)
		sass = reflect.Append(sass, reflect.ValueOf(10000))
		//intSlice2 := intSliceReflect.Interface().([]int)
		fmt.Println(sass)

	}

}

func testSet() {

	var array = []string{"a", "b", "c"}
	fmt.Println("array kind -> ", reflect.TypeOf(array).Kind())
	view := reflect.New(reflect.TypeOf(&array).Elem())
	fmt.Println(" view kind -> ", view.Kind())
	if view.Elem().CanSet() {
		view.Elem().Set(reflect.MakeSlice(reflect.TypeOf(array), 0, 0))
	}

	reflect.Append(view.Elem(), reflect.ValueOf("hello world"))
	fmt.Println(view.Elem())
}

func tttt() {
	intSlice := make([]int, 0)
	mapStringInt := make(map[string]int)

	sliceType := reflect.TypeOf(intSlice)
	mapType := reflect.TypeOf(mapStringInt)

	intSliceReflect := reflect.MakeSlice(sliceType, 0, 0)
	mapReflect := reflect.MakeMap(mapType)

	v := 10
	rv := reflect.ValueOf(v)
	intSliceReflect = reflect.Append(intSliceReflect, rv)
	intSlice2 := intSliceReflect.Interface().([]int)
	fmt.Println(intSlice2)

	k := "hello"
	rk := reflect.ValueOf(k)
	mapReflect.SetMapIndex(rk, rv)
	mapStringInt2 := mapReflect.Interface().(map[string]int)
	fmt.Println(mapStringInt2)
}
