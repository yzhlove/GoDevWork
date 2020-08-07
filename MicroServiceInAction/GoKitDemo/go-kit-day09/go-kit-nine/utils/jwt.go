package utils

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"reflect"
	"time"
)

var secret = []byte("jet_secret_verify")

type Token struct {
	Name string
	ID   int
	jwt.StandardClaims
}

func GenerToken(name string, id int) (string, error) {
	var token Token
	current := time.Now()
	token.StandardClaims = jwt.StandardClaims{
		ExpiresAt: current.Add(30 * time.Second).Unix(),
		IssuedAt:  current.Unix(),
		NotBefore: current.Unix(),
		Subject:   "login",
		Issuer:    name,
	}
	token.Name = name
	token.ID = id
	return jwt.NewWithClaims(jwt.SigningMethodHS256, token).SignedString(secret)
}

func ParseToken(tokenStr string) (*Token, error) {
	result, err := jwt.ParseWithClaims(tokenStr, &Token{}, func(token *jwt.Token) (interface{}, error) {
		if token.Valid || token.Method != jwt.SigningMethodHS256 {
			return nil, errors.New("sign error")
		}
		return secret, nil
	})
	if err != nil {
		return nil, err
	}
	if token, ok := result.Claims.(*Token); result.Valid && ok {
		return token, nil
	}
	return nil, errors.New("type err:" + reflect.TypeOf(result.Claims).String())
}
