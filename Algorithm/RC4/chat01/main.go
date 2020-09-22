package main

import (
	"crypto/rc4"
	"fmt"
)

//RC4加密

func main() {

	var secret = []byte("*#06#")
	cipher, err := rc4.NewCipher(secret)
	if err != nil {
		panic(err)
	}

	str := []byte("hello world")
	dst := make([]byte, len(str))
	cipher.XORKeyStream(dst, str)

	fmt.Printf("%x\n", dst)
	source := make([]byte, len(dst))
	cipher2, err := rc4.NewCipher(secret)
	if err != nil {
		panic(err)
	}
	cipher2.XORKeyStream(source, dst)
	fmt.Printf("%s\n", string(source))

}
