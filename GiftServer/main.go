package main

import (
	"WorkSpace/GoDevWork/GiftServer/app"
	"WorkSpace/GoDevWork/GiftServer/service"
	log "github.com/sirupsen/logrus"
)

func main() {
	if err := service.Run(app.New()); err != nil {
		log.Fatal(err)
	}
	log.Exit(0)
}
