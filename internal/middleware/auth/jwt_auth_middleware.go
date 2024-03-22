package auth

import (
	"github.com/gin-gonic/gin"
	"strings"
)

func AuthMiddleware(secretKey string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		token := strings.Split(authHeader, " ")

	}
}
