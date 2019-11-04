package main

import (
	"fmt"
	"time"
)

func main() {

	/*
		NewTrainBook ==> &{trainBookRecord:0xc0003b4b40 StartTime:1563854400 EndTime:1563861600}
		 timestamp: 1563867006
	*/

	transform := "2006-01-02 15:04:05"

	fmt.Println(time.Unix(1565404111, 0).Format(transform))
	fmt.Println(time.Unix(1569476383, 0).Format(transform))
	fmt.Println(time.Unix(1578984064, 0).Format(transform))

	fmt.Println()
	//
	//fmt.Println(3 ^ 3)
	//
	//fmt.Println(2401 / 100 % 10)
	//fmt.Println(2102 % 100)

	fmt.Println("value = ", (30865-29865)/5)
	fmt.Println("time = ", 1569476383-1569476183)

	//fmt.Println("hour => ",time.Now().Hour())



}
