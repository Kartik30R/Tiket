package utils

import (
	"github.com/golang-jwt/jwt/v5"
)

func GenerateJwt(claims *jwt.Claims, method jwt.SigningMethod, secret string)(string ,error){
	return jwt.NewWithClaims(method,*claims).SignedString([]byte(secret))
}