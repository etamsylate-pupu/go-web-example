package router

import (
	"go-web-example/controller"
	"go-web-example/middleware"

	"github.com/gin-gonic/gin"
)

// InitRouter init all the routes and start the engine
func InitRouter() *gin.Engine {

	e := gin.New()

	e.Use(middleware.Cross(), middleware.AccessLog(), gin.Recovery())

	e.GET("/ping", controller.Pong)

	//initAuthRouter 用户登录的相关接口
	initAuthRouter(e)

	return e
}
