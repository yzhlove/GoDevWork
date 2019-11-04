package main

import (
	"fmt"
	"io/ioutil"
	"yuzihan/Chats/chat25/obj"
)

func main() {

	write()
	read()

}

func read() {
	var c obj.Controller
	data, err := ioutil.ReadFile("temp.data")
	if err != nil {
		fmt.Println("read file err")
		return
	}
	c.DecodeUnmarshalMsg(data)
}

func write() {

	c := obj.NewController()
	s := obj.Student{
		Name: "yzh",
		Age:  22,
	}
	t := obj.Teacher{
		Id:   1,
		Name: "lcm",
		Age:  22,
	}

	c.List = append(c.List, &s, &t, &s, &t)

	data, err := c.EncodeMarshalMsg()
	if err != nil {
		return
	}

	_ = ioutil.WriteFile("temp.data", data, 0666)
}

//out:
//======student: {yzh 22}
//======teacher: {1 lcm 22}
