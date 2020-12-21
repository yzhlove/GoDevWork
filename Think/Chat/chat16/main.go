package main

import (
	"encoding/json"
	"fmt"
	"reflect"
	"time"
)

func main() {
	var ts = time.Now().Unix()
	var tmp = struct {
		Ts int64 `json:"timestamp"`
	}{
		Ts: ts,
	}
	data, err := json.Marshal(tmp)
	if err != nil {
		panic(err)
	}
	fmt.Println("data => ", string(data))
	var decoder interface{}
	if err := json.Unmarshal(data, &decoder); err != nil {
		panic(err)
	}
	fmt.Println(decoder)
	fmt.Println(reflect.TypeOf(decoder))
	result := decoder.(map[string]interface{})["timestamp"].(float64)
	fmt.Println("float64 => ", result)
	fmt.Println("int64 => ", int64(result))

	var decoder2 = map[string]int64{}
	if err := json.Unmarshal(data, &decoder2); err != nil {
		panic(err)
	}
	fmt.Println("decoder2 ==> ", decoder2)
	/*
		output:
			data =>  {"timestamp":1608272238}
			map[timestamp:1.608272238e+09]
			map[string]interface {}
			float64 =>  1.608272238e+09
			int64 =>  1608272238
			decoder2 ==>  map[timestamp:1608272407]
	*/
}
