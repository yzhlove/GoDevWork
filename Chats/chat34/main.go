package main

import (
	"encoding/json"
	"errors"
	"fmt"
)

func main() {
	To()
}

func To() (err error) {

	var User struct {
		Name string
	}
	User.Name = "yzh"

	err = errors.New("what are you doing")
	fmt.Printf("err -> %s \n", err)
	data, err := json.Marshal(User)
	if err != nil {
		fmt.Printf("json err -> %s \n", err)
	} else {
		fmt.Printf("json err -> %s \n", err)
	}
	data = data
	fmt.Printf("err -> %s \n", err)
	return
}
