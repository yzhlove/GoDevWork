package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	type User struct {
		Name string
		Age  int64
	}

	u := &User{
		Name: "yzh",
		Age:  16,
	}

	jsonStr, _ := json.Marshal(u)

	var data map[string]interface{}

	_ = json.Unmarshal(jsonStr, &data)

	fmt.Println(data["Age"].(float64))

	fmt.Println(data)
	fmt.Printf("%T\n", data)
	fmt.Printf("%T\n", data["Name"])
	fmt.Printf("%T\n", data["Age"])

}
