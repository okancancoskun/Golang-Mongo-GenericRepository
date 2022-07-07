package middlewares

import (
	"fmt"
	"myapp/services"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func AuthorizeJWT() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		const BearerSchema = "Bearer "
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "No Authorization header found"})

		}
		tokenString := authHeader[len(BearerSchema):]

		if token, err := services.ValidateToken(tokenString); err != nil {

			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Not Valid Token"})

		} else {
			if claims, ok := token.Claims.(jwt.MapClaims); !ok {
				ctx.AbortWithStatus(http.StatusUnauthorized)
			} else {
				fmt.Println("claimss", claims)
				if token.Valid {
					ctx.Set("user", map[string]interface{}{
						"_id":   claims["_id"],
						"email": claims["email"],
					})
				} else {
					ctx.AbortWithStatus((http.StatusUnauthorized))
				}
			}
		}
	}
}
