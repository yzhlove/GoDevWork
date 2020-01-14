package main

import (
	"encoding/base32"
	"fmt"
)

func main() {

	encode := base32.StdEncoding.EncodeToString([]byte("/usr/local/go/bin/go #gosetup"))
	fmt.Println(encode)

	decode, err := base32.StdEncoding.DecodeString(encode)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(decode))

}
