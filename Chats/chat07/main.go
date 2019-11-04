package main

import (
	"fmt"
	"time"
)

func main() {

	obj, err := time.Parse("2006-01-02 15:04:05", "2019-07-23 13:00:00")
	if err != nil {
		panic(err)
	}
	now := time.Now().Unix()
	fmt.Println(now, " - ", obj.Unix())

}
