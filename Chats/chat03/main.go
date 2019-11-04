package main

import (
	"fmt"
	"time"
)

func main() {

	now := time.Now()

	year, week := now.ISOWeek()

	fmt.Println(year, " - ", week)

}
