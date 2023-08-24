package router

import (
	"go-web-example/controller/auth"
	"go-web-example/middleware"

	"github.com/gin-gonic/gin"
)

// initAuthRouter .
func initAuthRouter(e *gin.Engine) {
	//注册
	e.POST("/register", auth.Register)
	//登录
	e.POST("/login", auth.Login)
	//登出
	e.POST("/logout", middleware.Authentication(), auth.Logout)
	//生成验证码
	e.GET("/draw/captcha", auth.DrawCaptcha)
	//校验验证码
	e.GET("/verify/captcha", auth.VerifyCaptcha)
	//查看用户信息
	e.GET("user/info", middleware.Authentication(), auth.UserInfo)
}
