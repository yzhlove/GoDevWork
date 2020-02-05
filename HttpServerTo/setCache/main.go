package main

import (
	"WorkSpace/GoDevWork/HttpServerTo/setCache/base"
	"fmt"
	"time"
)

func main() {

	start := time.Now()
	size := 1000000

	set := base.NewSetData(128, base.NewFileWrite())
	data := base.NewDataSource(set.GetChanString())
	data.GenerateSource(size)
	set.Write()
	//set.BufferWrite()
	d := time.Now().Sub(start)
	fmt.Printf("count:%dms avg:%dms \n", d.Milliseconds(), d.Milliseconds()/int64(size))
}
