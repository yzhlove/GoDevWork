package main

import (
	"WorkSpace/GoDevWork/Chats/auth/test11/config"
	"WorkSpace/GoDevWork/Chats/auth/test11/storage"
	"fmt"
)

func main() {

	config.ACLFilePath = "Chats/auth/test11/storage/test.ini"

	s, err := storage.NewStorage()
	if err != nil {
		panic(err)
	}

	fmt.Println(s.LoadAuth())
}
