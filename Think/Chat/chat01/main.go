package main

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io"
)

//golang rand.Reader

func main() {

	data := make([]byte, 32)
	if _, err := io.ReadFull(rand.Reader, data); err != nil {
		panic(err)
	}

	str := base64.URLEncoding.EncodeToString(data)

	md := md5.New()
	md.Write([]byte(str))
	fmt.Print(hex.EncodeToString(md.Sum(nil)))

}
