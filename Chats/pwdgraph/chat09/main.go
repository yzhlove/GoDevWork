package main

import (
	"crypto/rc4"
	"encoding/binary"
	"fmt"
)

//uint64

// 10 42 12
// id 随机数 校验码
// - RC4 crc32(随机数)
// 13为编码

func main() {
	encode()
}

func encode() {

	key := []byte("*#06#")
	solt := []byte("*#123456#*")
	_ = solt

	var id uint64 = 1000
	var number uint64 = 444333222

	bytesNumber := make([]byte, 42)
	binary.BigEndian.PutUint64(bytesNumber, number)
	cipher, err := rc4.NewCipher(key)
	if err != nil {
		panic(err)
	}

	newBytesNumber := make([]byte, len(bytesNumber))
	cipher.XORKeyStream(newBytesNumber, bytesNumber)

	id = id
	newNumber := binary.LittleEndian.Uint64(newBytesNumber)
	fmt.Println(newNumber)
	fmt.Printf("%0.4b \n", newNumber)

	number32 := binary.LittleEndian.Uint32(newBytesNumber)
	fmt.Println(number32)

}
