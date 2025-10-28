package middleware

import (
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

func AuthProtected(db *gorm.DB) gin.HandlerFunc{
	return func(ctx *gin.Context){
		autherizationHeader := ctx.GetHeader("authorization")
       if len(autherizationHeader)==0{
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "fail",
			"message": "Unauthorize",
		})
		ctx.Abort()
			return
	   }

	   tokenParts := strings.Split(autherizationHeader," ")

	   if len(tokenParts)!=2 || tokenParts[0] != "Bearer"{
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "fail",
			"message": "Unauthorize",
		})
		ctx.Abort()
			return
	   }

	   tokenstr:= tokenParts[1]
	   secret := []byte(os.Getenv("SECRET"))

	token, err := jwt.Parse(tokenstr, func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrSignatureInvalid
			}
			return secret, nil
		})

		if err != nil || !token.Valid {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"status":  "fail",
				"message": "Invalid or expired token",
			})
			ctx.Abort()
			return
		}

			claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"status":  "fail",
				"message": "Invalid token claims",
			})
			ctx.Abort()
			return
		}


        userId := claims["id"]
		ctx.Set("userId", userId)	
	

		ctx.Next()
	}
}