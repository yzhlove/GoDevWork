package main

import (
	"fmt"
	"time"
)

func main() {
	transform := "2006-01-02 15:04:05"
	start, err := time.Parse(transform, "")
	if err != nil {
		panic(err)
	}

	end, err := time.Parse(transform, "2019-07-23 11:00:00")
	if err != nil {
		panic(err)
	}

	obj := time.Now()
	fmt.Println("ct ", obj.Format(transform))

	expired := obj.Unix() + (end.Unix() - start.Unix())

	tobj := time.Unix(expired, 0)
	fmt.Println("nt ", tobj.Format(transform))

}
