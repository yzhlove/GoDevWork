package main

import (
	"encoding/binary"
	"fmt"
)

//骚操作

func main() {

	var number uint32 = 123456

	bytesNumber := make([]byte, 4)
	binary.BigEndian.PutUint32(bytesNumber, number)
	fmt.Println("bytesNumber => ", bytesNumber)
	newBytesNumber := make([]byte, 4)
	for i, b := range bytesNumber {
		newBytesNumber[len(bytesNumber)-i-1] = b
	}
	fmt.Println("newBytesNumber => ", newBytesNumber)

	newNumber := binary.LittleEndian.Uint32(newBytesNumber)
	fmt.Println("newNumber => ", newNumber)

}
