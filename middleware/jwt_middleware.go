package middleware

import (
	"net/http"
	"strings"

	"be-logbook-ppds/app/auth"
	"be-logbook-ppds/pkg/response"

	"github.com/gin-gonic/gin"
)

func JWTMiddleware(jwtSecret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenHeader := c.GetHeader("Authorization")

		if tokenHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, response.Response{
				Code:    http.StatusUnauthorized,
				Message: "Authorization header is missing",
			})
			return
		}

		tokenStr := strings.TrimPrefix(tokenHeader, "Bearer ")
		if tokenStr == tokenHeader {
			c.AbortWithStatusJSON(http.StatusUnauthorized, response.Response{
				Code:    http.StatusUnauthorized,
				Message: "Bearer token format required",
			})
			return
		}

		claims, err := auth.ValidateToken(tokenStr, jwtSecret)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, response.Response{
				Code:    http.StatusUnauthorized,
				Message: "Invalid or expired token: " + err.Error(),
			})
			return
		}

		c.Set("user_id", claims.UserID)
		c.Set("username", claims.Username)
		c.Set("role", claims.Role)

		c.Next()
	}
}