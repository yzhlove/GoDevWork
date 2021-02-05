package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

func main() {

	var data bytes.Buffer
	enc := gob.NewEncoder(&data)
	dec := gob.NewDecoder(&data)

	if err := enc.Encode(&User{ID: 123, Name: "yzh", Birthday: "1996", Expire: 1234567}); err != nil {
		panic(err)
	}

	fmt.Println("encoder => ", data.String())

	var user User
	if err := dec.Decode(&user); err != nil {
		panic(err)
	}

	fmt.Println("User => ", user)

}

type User struct {
	ID       uint64
	Name     string
	Birthday string
	Expire   int64
}
