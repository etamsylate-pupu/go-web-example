package middleware

import (
	"strings"

	"go-web-example/controller"
	"go-web-example/utils"

	"github.com/gin-gonic/gin"
)

// Authentication authenticate token
func Authentication() gin.HandlerFunc {

	return func(c *gin.Context) {

		authenticationHeader := c.Request.Header.Get("Authorization")
		if authenticationHeader == "" {
			controller.RespNeedAuthentication(c)
			return
		}

		token := strings.TrimPrefix(authenticationHeader, "Bearer ")

		tokenUserInfo, err := utils.ValidateJWT(token)
		if err != nil {
			controller.RespNeedAuthentication(c)
			return
		}

		c.Set("token_info", tokenUserInfo)

		c.Next()
	}
}
