package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	tsf := "2006-01-02 15:04:05"
	obj, err := time.ParseInLocation(tsf, "2019-09-27 09:00:00", time.Local)
	if err != nil {
		log.Fatal(err)
	}

	now := time.Now()

	fmt.Printf("%v %v %v \n", obj.Unix(), now.Unix(), now.Unix()-obj.Unix())

	fmt.Printf("sub hour => %v ", time.Now().Sub(obj).Hours())

}
