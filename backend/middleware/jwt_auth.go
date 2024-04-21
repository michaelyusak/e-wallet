package middleware

import (
	"strings"

	"e-wallet/apperror"
	"e-wallet/constants"
	"e-wallet/helper"

	"github.com/gin-gonic/gin"
)

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		t := strings.Split(authHeader, " ")

		if len(t) == 2 {
			token := t[1]
			claims, err := helper.ParseAndVerify(token)
			if err != nil {
				c.Error(apperror.StatusUnauthorized())
				c.Abort()
				return
			}
			c.Set(string(constants.UserId), int(claims["id"].(float64)))
		}

		c.Next()
	}
}
