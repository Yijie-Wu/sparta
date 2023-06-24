package utils

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"github.com/spf13/viper"
	"time"
)

var (
	SigningKey = []byte(viper.GetString("jwt.SigningKey"))
)

type JwtCustomClaims struct {
	ID uint
	NT string
	jwt.RegisteredClaims
}

func GenerateToken(id uint, nt string) (string, error) {
	jwtCustomClaims := JwtCustomClaims{
		ID: id,
		NT: nt,
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   "Token",
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(viper.GetDuration("jwt.TokenExpire") * time.Minute)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtCustomClaims)
	return token.SignedString(SigningKey)

}

func ParseToken(tokenStr string) (JwtCustomClaims, error) {
	jwtCustomClaims := JwtCustomClaims{}
	token, err := jwt.ParseWithClaims(tokenStr, &jwtCustomClaims, func(token *jwt.Token) (interface{}, error) {
		return SigningKey, nil
	})
	if err == nil && !token.Valid {
		err = fmt.Errorf("Invalid Token")
		return JwtCustomClaims{}, err
	}
	return jwtCustomClaims, err
}

func IsTokenValid(tokenStr string) bool {
	if _, err := ParseToken(tokenStr); err != nil {
		return false
	}
	return true
}
