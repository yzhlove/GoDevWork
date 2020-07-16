package utils

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"io"
)

func GetMD5String(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

func GetUUID() string {
	buf := make([]byte, 48)
	io.ReadFull(rand.Reader, buf)
	return GetMD5String(base64.URLEncoding.EncodeToString(buf))
}
