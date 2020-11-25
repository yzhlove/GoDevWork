package main

import "fmt"

func main() {
	//bytes.Contains(h, []byte("\xF0\x9F")) || bytes.Contains(h, []byte("\xC2\xA0"))
	b := []byte("\xF0\x9F\xC2\xA0")
	fmt.Println("s ==> ", string(b))

}
