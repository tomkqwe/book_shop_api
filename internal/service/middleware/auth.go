package middleware

import (
	"book_shop_api/internal/transport/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")

		if authHeader == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": "Authorization header missing",
			})
			ctx.Abort()
			return
		}

		tokenString := strings.TrimSpace(strings.TrimPrefix(authHeader, "Bearer"))

		if tokenString == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Token missing"})
			ctx.Abort()
			return
		}

		claims, err := utils.VerifyToken(tokenString)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			ctx.Abort()
			return
		}

		ctx.Set("id", claims.Id)
		ctx.Set("email", claims.Email)
		ctx.Next()
	}
}
