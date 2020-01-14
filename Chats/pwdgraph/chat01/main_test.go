package main

import (
	"crypto/rc4"
	"fmt"
	"strconv"
	"testing"
)

func Test_RC4(t *testing.T) {

	status := make(map[int]struct{}, 8)

	for i := 100; i < 10000000; i++ {
		str := fmt.Sprintf("%x", RC4Encrypt([]byte(strconv.Itoa(i))))
		if _, ok := status[len(str)]; !ok {
			status[len(str)] = struct{}{}
		}
	}

	for key := range status {
		fmt.Println("length => ", key)
	}
}

func Test_RC4Length(t *testing.T) {

	cipher, err := rc4.NewCipher(key)
	if err != nil {
		t.Error(err)
		return
	}

	src := []byte("yuzihan")
	dst := make([]byte, len(src))
	cipher.XORKeyStream(dst, src)

	fmt.Println("src => ", len(src), " dst => ", len(dst))
	fmt.Printf("%v %v \n", src, dst)
	for _, b := range dst {
		fmt.Print(string(b))
	}

}
