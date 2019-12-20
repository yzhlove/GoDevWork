package main

import (
	"WorkSpace/GoDevWork/Chats/chat52/obj"
	"fmt"
	"io/ioutil"
)

func main() {
	data, err := ioutil.ReadFile("Chats/chat52/kitchen")
	if err != nil {
		panic(err)
	}

	var k obj.KitchenController
	if _, err := k.UnmarshalMsg(data); err != nil {
		panic(err)
	}
	fmt.Println(k)
}
