package main

import (
	"crypto/rc4"
	"fmt"
)

//RC4 加解密

var key = []byte("abcdefghijklmnopqrstuvwxyz")

func main() {

	result := RC4Encrypt([]byte("yuzihan"))
	fmt.Printf("%x \n", result)

	fmt.Println(string(RC4Decrypt(result)))

}

func RC4Encrypt(value []byte) []byte {
	if cipher, err := rc4.NewCipher(key); err != nil {
		panic(err)
	} else {
		cipher.XORKeyStream(value, value)
	}
	return value
}

func RC4Decrypt(source []byte) []byte {
	if cipher, err := rc4.NewCipher(key); err != nil {
		panic(err)
	} else {
		cipher.XORKeyStream(source, source)
	}
	return source
}
