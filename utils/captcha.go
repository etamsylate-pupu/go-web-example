package utils

import (
	"github.com/mojocn/base64Captcha"
)

//CaptchaResult 生成结果
type CaptchaResult struct {
	CaptchaID  string `json:"captcha_id"`
	Base64Blog string `json:"base_64_blog"`
}

// 默认存储10240个验证码，每个验证码10分钟过期
var store = base64Captcha.DefaultMemStore

/*
// 多机使用redis 重定义
//Captcha 重定义captcha
type Captcha struct {
	Driver base64Captcha.Driver
	Store  base64Captcha.Store
}

//NewCaptcha creates a captcha instance from driver and store
func NewCaptcha(driver base64Captcha.Driver, store base64Captcha.Store) *Captcha {
	return &Captcha{Driver: driver, Store: store}
}

//Generate generates a random id, base64 image string or an error if any
func (c *Captcha) Generate() (id, b64s, answer string, err error) {
	id, content, answer := c.Driver.GenerateIdQuestionAnswer()
	item, err := c.Driver.DrawCaptcha(content)
	if err != nil {
		return "", "", "", err
	}
	err = c.Store.Set(id, answer)
	if err != nil {
		return "", "", "", err
	}
	b64s = item.EncodeB64string()
	return
}*/

//GenerateCaptcha 生成图片验证码
func GenerateCaptcha() (interface{}, error) {
	// 生成默认数字
	//driver := base64Captcha.DefaultDriverDigit
	// 此尺寸的调整需要根据网站进行调试，链接：
	// https://captcha.mojotv.cn/
	driver := base64Captcha.NewDriverDigit(70, 130, 4, 0.8, 100)

	/*
		//多机使用redis
			captcha := NewCaptcha(driver, store)
			id, b64s, code, err := captcha.Generate()*/

	// 生成base64图片
	captcha := base64Captcha.NewCaptcha(driver, store)
	// 获取
	id, b64s, err := captcha.Generate()

	if err != nil {
		return nil, err
	}

	//多机 把id、code 放入redis
	captchaResult := CaptchaResult{CaptchaID: id, Base64Blog: b64s}

	return captchaResult, nil
}

//VerifyCaptcha 校验图片验证码,并清除内存空间
func VerifyCaptcha(id string, value string) bool {
	// 只要id存在，就会校验并清除，无论校验的值是否成功, 所以同一id只能校验一次
	// 注意：id,b64s是空 也会返回true
	if id == "" || value == "" {
		return false
	}
	verifyResult := store.Verify(id, value, true) // 不管验证结果成功或失败都会清除
	return verifyResult
}
