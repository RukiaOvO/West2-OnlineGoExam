package utils

import (
	"github.com/golang-jwt/jwt"
	"time"
	"todolist/consts"
)

type MyClaims struct {
	Id       int64  `json:"id"`
	UserName string `json:"user_name"`
	jwt.StandardClaims
}

func TokenGen(id int64, name string) (string, error) {
	temp := MyClaims{
		id,
		name,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(consts.AccessTokenExpireTime).Unix(),
			Issuer:    consts.AccessIssuer,
		},
	}

	return jwt.NewWithClaims(jwt.SigningMethodHS256, temp).SignedString([]byte(consts.JwtSecret))
}

func TokenParse(tokenString string) (*MyClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (i interface{}, err error) {
		return []byte(consts.JwtSecret), nil
	})
	if token != nil {
		if claims, ok := token.Claims.(*MyClaims); ok && token.Valid {
			return claims, nil
		}
	}

	return nil, err
}
