package utils

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"io"
)

func encode(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

func GetUID() string {
	sb := make([]byte, 32)
	if _, err := io.ReadFull(rand.Reader, sb); err != nil {
		panic(err)
	}
	return encode(base64.URLEncoding.EncodeToString(sb))
}
