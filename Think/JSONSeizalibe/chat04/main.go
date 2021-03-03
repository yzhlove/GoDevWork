package main

import (
	"encoding/json"
	"fmt"
)

var str = `[{"A":null,"B":"null","C":"123"},{"A":"a","B":"null","C":"123"}]`

func main() {

	var aa = make([]Dec, 0, 4)
	if err := json.Unmarshal([]byte(str), &aa); err != nil {
		panic(err)
	}
	fmt.Println(aa)
	
}

type Dec struct {
	A string
	B string
}
