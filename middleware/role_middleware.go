package middleware

import (
	"net/http"

	"be-logbook-ppds/pkg/response"

	"github.com/gin-gonic/gin"
)

func RoleMiddleware(allowedRoles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		userRole := c.GetString("role")

		allowed := false
		for _, r := range allowedRoles {
			if userRole == r {
				allowed = true
				break
			}
		}

		if !allowed {
			c.AbortWithStatusJSON(http.StatusForbidden, response.Response{
				Code:    http.StatusForbidden,
				Message: "Access forbidden: insufficient role permissions",
			})
			return
		}

		c.Next()
	}
}