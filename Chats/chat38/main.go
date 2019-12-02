package main

import (
	"encoding/json"
	"fmt"
)

//test json

func main() {

	type User struct {
		Name     string `json:"name"`
		Age      int    `json:"age"`
		Birthday string `json:"birthday"`
	}

	data, err := json.Marshal(User{"yzhlove", 22, "12-24"})
	if err != nil {
		panic(err)
	}

	fmt.Println("data => ", string(data))

	var Temp struct {
		Value string `json:"name"`
	}

	if err := json.Unmarshal(data, &Temp); err != nil {
		panic(err)
	}

	fmt.Println("Temp => ", Temp)

}
