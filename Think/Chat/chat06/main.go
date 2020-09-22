package main

import (
	"fmt"
	"strings"
)

func main() {

	str := "heart_beat_req"

	check := "_req"

	res := strings.LastIndex(str, check)
	fmt.Println(res)
	fmt.Println(str[:res])

	hashMap := make(map[string][]string)
	hashMap[str[:res]] = make([]string, 1)
	hashMap[str[:res]][0] = "hello world"

	fmt.Println(hashMap)

}
