package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	var maps = map[string]interface{}{}
	maps["A"] = nil
	maps["B"] = "null"
	maps["C"] = "123"

	data, _ := json.Marshal(maps)
	fmt.Println("data => ", string(data))

}
