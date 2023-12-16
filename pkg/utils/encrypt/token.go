package encrypt

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type Claim struct {
	Username string
	UserId   uint32
	jwt.StandardClaims
}

var jwtKey = []byte("golang-im-server")

func CreateToken(id uint32, username string) (string, error) {
	claims := &Claim{
		Username: username,
		UserId:   id,
		StandardClaims: jwt.StandardClaims{
			IssuedAt: time.Now().Unix(),
			Issuer:   "mgh",
			Subject:  "User_Token",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return tokenStr, nil
}

func ParseToken(token string) (string, uint32, error) {
	if token == "" {
		return "", 0, errors.New("empty String")
	}
	data, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		return "", 0, err
	}

	claim, ok := data.Claims.(jwt.MapClaims)
	if !ok {
		return "", 0, errors.New("parse Error")
	}

	return claim["Username"].(string), uint32(claim["UserId"].(float64)), nil
}
