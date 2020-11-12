package main

import (
	"fmt"
	"reflect"
)

func main() {

	r := &record{data: make(map[uint32]*result)}
	r.data[1] = &result{Id: 1234}

	fmt.Println(r.Get(1), " ", r.Get(2))

	fmt.Println(reflect.TypeOf(r.Get(2)), " == ", reflect.ValueOf(r.Get(2)))

	if rs1 := r.Get(2); rs1 == nil {
		fmt.Println("Get 2 is nil")
	} else {
		fmt.Println("Get 2 is not nil")
	}

}

type result struct {
	Id uint32
}

type record struct {
	data map[uint32]*result
}

func (r *record) Get(id uint32) *result {
	return r.data[id]
}
