package auth

import (
	"go-web-example/biz/systembiz"
	"go-web-example/controller"
	"go-web-example/errorcode"
	"go-web-example/model"
	"go-web-example/request/authrequest"
	"go-web-example/utils"

	"github.com/gin-gonic/gin"
)

// Register
func Register(c *gin.Context) {

	inputParams := authrequest.RegisterInputParams{}

	if err := inputParams.CheckInputParams(c); err != nil {
		controller.Resp(c, nil, err)
		return
	}

	err := systembiz.UserAdd(inputParams)

	controller.Resp(c, nil, err)
}

// Login login user and set session token
func Login(c *gin.Context) {
	inputParams := authrequest.LoginInputParams{}

	if err := inputParams.CheckInputParams(c); err != nil {
		controller.Resp(c, nil, err)
		return
	}

	//根据phone获取用户信息
	userInfo, err := model.GetInfoByPhone(inputParams.Phone)
	if err != nil {
		controller.Resp(c, nil, err)
		return
	}

	if userInfo.UserPwd != utils.Sha1([]byte(inputParams.Password)) {
		controller.Resp(c, nil, errorcode.New(errorcode.ErrParams, "密码错误", nil))
		return
	}

	if userInfo.ISUse == model.ModelStop {
		controller.Resp(c, nil, errorcode.New(errorcode.ErrParams, "账号停用，请联系管理员", nil))
		return
	}

	//生成 token
	token, err := utils.GenerateJWT(userInfo.UserPhone, userInfo.ID)
	if err != nil {
		controller.Resp(c, nil, err)
		return
	}

	//返回
	controller.Resp(c, map[string]interface{}{
		"user_name":  userInfo.UserName,
		"user_phone": userInfo.UserPhone,
		"token":      token,
	}, nil)
}

// Logout return login out
func Logout(c *gin.Context) {
	//清除token  jwt无法清除  讲 jwt产生的token 放入redis

	controller.Resp(c, nil, nil)
}

// DrawCaptcha 生成验证码
func DrawCaptcha(c *gin.Context) {
	res, err := utils.GenerateCaptcha()

	controller.Resp(c, res, err)
}

// VerifyCaptcha 验证码校验
func VerifyCaptcha(c *gin.Context) {
	inputParams := authrequest.VerifyCaptchaInputParams{}
	if err := inputParams.CheckInputParams(c); err != nil {
		controller.Resp(c, nil, err)
		return
	}

	if ok := utils.VerifyCaptcha(inputParams.CaptchaID, inputParams.Code); !ok {
		controller.Resp(c, nil, errorcode.New(errorcode.ErrParams, "验证码错误请重新输入", nil))
		return
	}

	controller.Resp(c, nil, nil)
}

// UserInfo return user info
func UserInfo(c *gin.Context) {
	tokenUserInfo, _ := c.Get("token_info")

	res, err := systembiz.UserInfo(tokenUserInfo.(utils.TokenUserInfo).UserID)

	controller.Resp(c, res, err)
}
