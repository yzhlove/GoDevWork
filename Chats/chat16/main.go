package main

import (
	"fmt"
	"time"
)

func main() {

	str := "2006-01-02 15:04:05"

	t := time.Now()
	ntm := t.AddDate(0, 0, -1)

	fmt.Println(ntm.Format(str))

	t1 := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
	fmt.Println(t1.Format(str))

	fmt.Println(ntm.Unix(), " - ", t1.Unix())

	t2 := t1.AddDate(0, 0, -1)
	fmt.Println(t2.Format(str))

}
