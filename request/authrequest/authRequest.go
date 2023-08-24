package authrequest

import (
	"strings"

	"go-web-example/errorcode"
	"go-web-example/request"

	"github.com/gin-gonic/gin"
)

// RegisterInputParams return 输入参数
type RegisterInputParams struct {
	Phone    string `json:"phone" form:"phone" validate:"required,len=11,checkPhoneRex" comment:"手机号"`
	Password string `json:"password" form:"password" validate:"required" comment:"密码"`
	UserName string `json:"user_name" form:"user_name" validate:"required" comment:"姓名"`
	NickName string `json:"nick_name" form:"nick_name" validate:"required" comment:"昵称"`
}

// CheckInputParams 检查参数是否正确
func (u *RegisterInputParams) CheckInputParams(c *gin.Context) error {
	if err := c.ShouldBind(u); err != nil {
		return err
	}

	if errs, err := request.Validate(u); err != nil {
		return errorcode.New(errorcode.ErrParams, strings.Join(errs, ","), nil)
	}

	return nil
}


// LoginInputParams return 输入参数
type LoginInputParams struct {
	Phone    string `json:"phone" form:"phone" validate:"required,len=11,checkPhoneRex" comment:"手机号"`
	Password string `json:"password" form:"password" validate:"required" comment:"密码"`
}

// CheckInputParams 检查参数是否正确
func (l *LoginInputParams) CheckInputParams(c *gin.Context) error {
	if err := c.ShouldBind(l); err != nil {
		return err
	}

	if errs, err := request.Validate(l); err != nil {
		return errorcode.New(errorcode.ErrParams, strings.Join(errs, ","), nil)
	}

	return nil
}

// VerifyCaptchaInputParams return 输入参数
type VerifyCaptchaInputParams struct {
	Code      string `json:"code" form:"code" validate:"required,len=4" comment:"验证码"`
	CaptchaID string `json:"captcha_id" form:"captcha_id" validate:"required" comment:"验证码id"`
}

// CheckInputParams 检查参数是否正确
func (v *VerifyCaptchaInputParams) CheckInputParams(c *gin.Context) error {
	if err := c.ShouldBind(v); err != nil {
		return err
	}

	if errs, err := request.Validate(v); err != nil {
		return errorcode.New(errorcode.ErrParams, strings.Join(errs, ","), nil)
	}

	return nil
}
