package main

import (
	"log"
	"time"
)

func main() {

	a := time.Now()
	time.Sleep(time.Second)
	b := time.Now()

	if b.Before(a) {
		log.Println("true")
	} else {
		log.Println("false")
	}

}
