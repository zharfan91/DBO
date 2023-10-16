package middleware

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func GetTokenRemainingValidity(timestamp interface{}) bool {
	if validity, ok := timestamp.(float64); ok {
		tm := time.Unix(int64(validity), 0)
		remainer := tm.Sub(time.Now())
		if remainer > 0 {
			return true
		}
	}
	return false
}
func ValidateToken(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			//nil secret key
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
}

// AuthUserMiddlerWare ...
func AuthorizeUserMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		const BearerSchema string = "Bearer "
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "No Authorization header found"})
			return

		}

		data := strings.Split(authHeader, " ")
		if len(data) != 2 {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "No Authorization bearer header found"})
			return
		}
		log.Println(len(data))
		tokenString := authHeader[len(BearerSchema):]

		if token, err := ValidateToken(tokenString); err != nil {

			fmt.Println("token", tokenString, err.Error())
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Not Valid Token or Token Expired"})

		} else {

			if claims, ok := token.Claims.(jwt.MapClaims); !ok {
				ctx.AbortWithStatus(http.StatusUnauthorized)

			} else {
				if token.Valid {
					ctx.Set("name", claims["name"])
					fmt.Println("during authorization", claims["name"])
					isExp := GetTokenRemainingValidity(claims["exp"])
					log.Println(isExp)
					if !isExp {
						ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
							"error": "Token Expired",
						})
						return
					}
				} else {
					ctx.AbortWithStatus(http.StatusUnauthorized)
				}

			}
		}
	}
}
