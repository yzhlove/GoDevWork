package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"reflect"
)

func main() {

	var r io.Reader
	fmt.Println(reflect.TypeOf(r))
	r = os.Stdin
	fmt.Println(reflect.TypeOf(r))
	r = bufio.NewReader(r)
	fmt.Println(reflect.TypeOf(r))
	r = new(bytes.Buffer)
	fmt.Println(reflect.TypeOf(r))

}
