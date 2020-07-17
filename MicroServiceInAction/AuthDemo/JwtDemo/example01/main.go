package main

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

////////////////////////////////////////////////
// jwt
////////////////////////////////////////////////

var tokenSecret = []byte("jwt_secret_v3")

type Token struct {
	Name  string `json:"name"`
	DcId  int    `json:"id"`
	ExpTs int64  `json:"ts"`
	jwt.StandardClaims
}

func CreateJwtToken(name string, dcId int) (string, error) {
	var token Token
	token.StandardClaims = jwt.StandardClaims{
		Audience:  "",                                     //受众群体
		ExpiresAt: time.Now().Add(3 * time.Second).Unix(), //到期时间
		Id:        "",                                     //编号
		IssuedAt:  time.Now().Unix(),                      //签发时间
		Issuer:    "kit_v3",                               //签发人
		NotBefore: time.Now().Unix(),                      //生效时间
		Subject:   "login",                                //主题
	}
	token.Name = name
	token.DcId = dcId
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, token)
	return tokenClaims.SignedString(tokenSecret)
}

func ParserToken(tokenString string) (jwt.MapClaims, error) {
	jwtToken, err := jwt.ParseWithClaims(tokenString, jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		if token.Method != jwt.SigningMethodHS256 {
			return nil, errors.New("signing method err")
		}
		return tokenSecret, nil
	})
	if err != nil || jwtToken == nil {
		return nil, err
	}
	claim, ok := jwtToken.Claims.(jwt.MapClaims)
	if ok && jwtToken.Valid {
		return claim, nil
	}
	return nil, errors.New("invalid token")
}

func main() {

	token, err := CreateJwtToken("love", 1234)
	if err != nil {
		panic(err)
	}
	fmt.Println("token => ", token)
	result, err := ParserToken(token)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)

}
