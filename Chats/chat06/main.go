package main

import (
	"fmt"
	"time"
)

func main() {
	var ts int64 = 1563850041
	obj := time.Unix(ts, 0)
	fmt.Println(obj.Format("2006-01-02 15:04:05"))
}
