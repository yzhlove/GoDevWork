package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	type User struct {
		Name string `json:"name"`
		Age  int64  `json:"age"`
	}

	u := User{}
	_ = json.Unmarshal([]byte(`{"name":"yzh","age":"123456789"}`), &u)
	fmt.Println(u)

	_ = json.Unmarshal([]byte(`{"name":"yzh","age":123456789}`), &u)
	fmt.Println(u)

}
