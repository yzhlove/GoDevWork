package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

func main() {

	var x int32 = 123
	var y int32
	buf := new(bytes.Buffer)
	_ = binary.Write(buf, binary.BigEndian, x)
	_ = binary.Read(buf, binary.LittleEndian, &y)
	fmt.Println("y = ", y)

}
