package main

import (
	"fmt"
	"path/filepath"
)

func main() {

	a := "/User/local/a.b.c.d.txt"

	res := filepath.Base(a)
	fmt.Println("res => ", res)

}
