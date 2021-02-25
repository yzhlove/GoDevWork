package main

import (
	"fmt"
	"regexp"
)

func main() {

	//str := "./123/456/789/tables/b/c/tables/d/e.xlsx"
	str := `.\tables\b\c\d\e.xlsx`

	r := regexp.MustCompile(`.*[\\/]tables([a-zA-Z0-9/\\]*)\.xlsx`)
	result := r.FindStringSubmatch(str)
	fmt.Println("result => ", result)

}
