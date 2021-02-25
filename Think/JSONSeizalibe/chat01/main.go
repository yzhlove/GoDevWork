package main

import (
	"encoding/json"
	"fmt"
)

var str = `{"name":"hello world","birthday":"1988-04-21","age":10}`

func main() {

	if !json.Valid([]byte(str)) {
		panic("json data type error")
	}

	var data json.RawMessage
	if err := json.Unmarshal([]byte(str), &data); err != nil {
		panic(err)
	}

	fmt.Println("rawMessage => ", string(data))

	var maps = map[string]interface{}{"text": data}

	if result, err := json.Marshal(maps); err != nil {
		panic(err)
	} else {
		fmt.Println("result => ", string(result))
	}

}
