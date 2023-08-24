package middleware

import (
	"go-web-example/conf"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// AccessLog access log
func AccessLog() gin.HandlerFunc {
	return func(c *gin.Context) {

		start := time.Now()

		c.Next()

		stop := time.Now()

		conf.AccessLog.WithFields(logrus.Fields{
			"url":     c.Request.RequestURI,
			"method":  c.Request.Method,
			"ip":      c.ClientIP(),
			"code":    c.Writer.Status(),
			"latency": stop.Sub(start).Milliseconds(),
		}).Info()
	}
}
