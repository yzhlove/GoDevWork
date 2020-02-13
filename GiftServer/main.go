package main

import (
	"WorkSpace/GoDevWork/GiftServer/service"
	log "github.com/sirupsen/logrus"
)

func main() {
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
	log.Exit(0)
}
