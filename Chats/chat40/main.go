package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	var c struct {
		E string `json:"event"`
	}
	_ = json.Unmarshal([]byte(``), &c)
	fmt.Println(c.E)

}
