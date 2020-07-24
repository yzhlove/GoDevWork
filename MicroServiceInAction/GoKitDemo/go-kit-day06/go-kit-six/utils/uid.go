package utils

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"io"
)

func md5Str(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

func GetUID() string {
	buf := make([]byte, 42)
	if _, err := io.ReadFull(rand.Reader, buf); err != nil {
		panic(err)
	}
	return md5Str(base64.URLEncoding.EncodeToString(buf))
}
