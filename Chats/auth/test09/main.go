package main

import (
	"WorkSpace/GoDevWork/Chats/auth/test09/apt"
	"WorkSpace/GoDevWork/Chats/auth/test09/storage"
)

func main() {

	s, err := storage.NewStorage()
	if err != nil {
		panic(err)
	}

	adapter, err := apt.NewEnforcerContext(s)
	if err != nil {
		panic(err)
	}

	adapter = adapter
}
