package main

import (
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
