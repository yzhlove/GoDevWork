package main

import (
	"fmt"
	"io/ioutil"
	"yuzihan/Chats/chat05/obj"
)

func main() {
	read()
}

func write() {

	user := obj.UserInfo{Name: "yzh", Age: 18}

	data, err := user.MarshalMsg(nil)
	if err != nil {
		panic(err)
	}

	if err := ioutil.WriteFile("temp.dat", data, 0666); err != nil {
		panic(err)
	}
}

func read() {

	var user obj.UserInfo

	data, err := ioutil.ReadFile("temp.dat")
	if err != nil {
		panic(err)
	}

	if _, err = user.UnmarshalMsg(data); err != nil {
		panic(err)
	}

	fmt.Printf("user => %+v \n", user)

}
