package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {

	var sb strings.Builder
	sb.WriteString("hello world")
	sb.WriteString(",")
	sb.WriteString("abc,")

	n := strings.TrimRight(sb.String(), ",")
	fmt.Println(sb.String(), " - ", n)

	test2()

}

func test2() {

	var sb bytes.Buffer
	sb.WriteString("hello world")
	sb.WriteString(",")
	sb.WriteString("abc,")
	sb.UnreadByte()
	fmt.Println(sb.String())
}
