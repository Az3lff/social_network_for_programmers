package auth

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"social_network_for_programmers/pkg/auth/utils/tokenutil"
	"social_network_for_programmers/pkg/responses"
	"strings"
)

func JwtAuthMiddleware(secretKey string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		t := strings.Split(authHeader, " ")

		if len(t) == 2 {
			authToken := t[1]
			authorized, err := tokenutil.IsAuthorized(authToken, secretKey)
			if err != nil {
				c.JSON(http.StatusUnauthorized, responses.ErrorResponse{Message: err.Error()})
				c.Abort()
				return
			}

			if authorized {
				c.Next()
				return
			}
		}

		c.JSON(http.StatusUnauthorized, responses.ErrorResponse{Message: "Not authorized"})
		c.Abort()
	}
}
