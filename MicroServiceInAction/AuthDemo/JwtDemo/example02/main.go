package main

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

var tokenSecret = []byte("jwt_secret_vae")

type Token struct {
	jwt.StandardClaims
	Name string `json:"name"`
	Id   int    `json:"id"`
}

func createToken(name string, id int) (string, error) {
	var token Token
	t := time.Now()
	token.StandardClaims = jwt.StandardClaims{
		ExpiresAt: t.Add(20 * time.Second).Unix(),
		IssuedAt:  t.Unix(),
		NotBefore: t.Add(5 * time.Second).Unix(),
		Issuer:    "google.inc",
	}
	token.Name = name
	token.Id = id
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, token)
	return tokenClaims.SignedString(tokenSecret)
}

func parseToken(tokenString string) (*Token, error) {
	jwtToken, err := jwt.ParseWithClaims(tokenString, &Token{}, func(token *jwt.Token) (interface{}, error) {
		fmt.Println("parser token :", token)
		if token.Method != jwt.SigningMethodHS256 {
			return nil, errors.New("sign algorithm err")
		}
		return tokenSecret, nil
	})
	if err != nil {
		panic("parse err:" + err.Error())
	}
	if claims, ok := jwtToken.Claims.(*Token); ok && jwtToken.Valid {
		return claims, nil
	}
	return nil, errors.New("parse token err")
}

func main() {

	tokenString, err := createToken("love", 1314)
	if err != nil {
		panic(err)
	}

	fmt.Println("token string =>", tokenString)
	time.Sleep(7 * time.Second)
	tk, err := parseToken(tokenString)
	if err != nil {
		panic(err)
	}
	fmt.Println("token tk --> ", tk)

}
