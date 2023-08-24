package controller

import (
	"github.com/gin-gonic/gin"
)

// Pong returns server status
func Pong(c *gin.Context) {
	Resp(c, "pong", nil)
}
