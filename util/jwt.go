package jwt

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
)

type JwtPayLoad struct {
	Appid  string `json:"appid"`
	Appkey string `json:"appkey"`
}

type CustomClaims struct {
	JwtPayLoad
	jwt.StandardClaims
}

func GenToken(user JwtPayLoad) (string, error) {
	var secretKey = []byte("wenli")
	claims := CustomClaims{
		user,
		jwt.StandardClaims{
			ExpiresAt: 15000, // 过期时间
			Issuer:    "wenli",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secretKey)
}

func ParseToken(tokenStr string) (*CustomClaims, error) {
	var secretKey = []byte("wenli")
	token, err := jwt.ParseWithClaims(tokenStr, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		println("--------------------------")
		return claims, nil
	}
	return nil, errors.New("invalid token")
}
