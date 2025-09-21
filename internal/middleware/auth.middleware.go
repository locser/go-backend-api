package middleware

import (
	"go-backend/responses"

	"github.com/gin-gonic/gin"
)

func AuthenMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {

		token := c.GetHeader("Authorization")

		if token != "valid token" {
			responses.ErrorResponse(c, responses.StatusUnauthorized, "Request Invalid")
			c.Abort()
			return

		}

		c.Next()
	}
}
