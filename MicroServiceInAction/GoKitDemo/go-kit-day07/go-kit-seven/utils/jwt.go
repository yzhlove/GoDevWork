package utils

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

var jwtSecret = []byte("jwt_secret_private")

const JwtContextKey = "jwt_context_key"

type Token struct {
	Name string
	DcId int
	jwt.StandardClaims
}

func GenericToken(name string, id int) (string, error) {
	var token Token
	t := time.Now()
	token.StandardClaims = jwt.StandardClaims{
		ExpiresAt: t.Add(30 * time.Second).Unix(),
		IssuedAt:  t.Unix(),
		NotBefore: t.Unix(),
		Subject:   "login",
		Issuer:    "go-kit",
	}
	token.Name = name
	token.DcId = id
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, token)
	return tokenClaims.SignedString(jwtSecret)
}

func ParseToken(tokenString string) (*Token, error) {
	jwtToken, err := jwt.ParseWithClaims(tokenString, &Token{}, func(token *jwt.Token) (interface{}, error) {
		if token.Method != jwt.SigningMethodHS256 {
			return nil, errors.New("sign error")
		}
		return jwtSecret, nil
	})
	if err != nil {
		return nil, err
	}
	if token, ok := jwtToken.Claims.(*Token); ok && jwtToken.Valid {
		return token, nil
	}
	return nil, errors.New("parse token err")
}
