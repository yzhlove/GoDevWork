package util

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io"
	"reflect"
)

func GetMD5String(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

func GetUID() string {
	b := make([]byte, 48)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}
	return GetMD5String(base64.URLEncoding.EncodeToString(b))
}

type RegisterStructMaps struct {
	maps map[string]reflect.Type
}

func NewRegisterStructMaps() *RegisterStructMaps {
	return &RegisterStructMaps{maps: make(map[string]reflect.Type)}
}

func (rsm *RegisterStructMaps) New(name string) (interface{}, error) {
	if v, ok := rsm.maps[name]; ok {
		return reflect.New(v).Interface(), nil
	}
	return nil, fmt.Errorf("not found %s struct", name)
}

func (rsm *RegisterStructMaps) CheckElem(name string) bool {
	if _, ok := rsm.maps[name]; ok {
		return true
	}
	return false
}

func (rsm *RegisterStructMaps) Register(name string, c interface{}) {
	rsm.maps[name] = reflect.TypeOf(c).Elem()
}
