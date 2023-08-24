package controller

import (
	"go-web-example/conf"
	"go-web-example/errorcode"
	"net"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// Resp used for generally response to client
func Resp(c *gin.Context, payload interface{}, err error) {

	if err == nil {
		successResponse(c, http.StatusOK, payload)
	} else {
		errorResponse(c, http.StatusOK, err)
	}
}

// RespWithStatus used for generally response to client with http status
func RespWithStatus(c *gin.Context, statusCode int, payload interface{}, err error) {

	if err == nil {
		successResponse(c, statusCode, payload)
	} else {
		errorResponse(c, statusCode, err)
	}
}

// RespAbortWithStatus abort the handle chain and response with http status
func RespAbortWithStatus(c *gin.Context, statusCode int, err error) {

	c.Abort()

	c.JSON(statusCode, gin.H{
		"code":    guessErrorCode(err),
		"message": guessErrorMessage(err),
	})
}

// RespNeedAuthentication Response 401 Need Authenticate
func RespNeedAuthentication(c *gin.Context) {

	RespAbortWithStatus(c, http.StatusOK,
		errorcode.New(errorcode.ErrAuthenticationRequired, "登录状态无效，请重新登录", nil))
}

// RespRoleNeedPermission Response 401 role Need permission
func RespRoleNeedPermission(c *gin.Context) {

	RespAbortWithStatus(c, http.StatusOK,
		errorcode.New(errorcode.ErrRolePermissionRequired, errorcode.RolePermissionRequiredMsg, nil))
}

func successResponse(c *gin.Context, statusCode int, payload interface{}) {
	if payload == nil {
		payload = map[string]string{}
	}
	c.JSON(statusCode, gin.H{
		"req_ver": requestVersion(c),
		"code":    errorcode.Success,
		"message": errorcode.SuccessMsg,
		"data":    payload,
	})
}

func errorResponse(c *gin.Context, statusCode int, err error) {

	c.JSON(statusCode, gin.H{
		"req_ver": requestVersion(c),
		"code":    guessErrorCode(err),
		"message": guessErrorMessage(err),
		"data":    map[string]string{},
	})
}

func requestVersion(c *gin.Context) int {

	var requestVersion string
	if c.Request.Method == "POST" {
		requestVersion = c.PostForm("req_ver")
		if requestVersion == "" {
			requestVersion = "0"
		}

	} else {
		requestVersion = c.DefaultQuery("req_ver", "0")
	}

	version, err := strconv.Atoi(requestVersion)
	if err != nil {
		return 0
	}

	return version
}

func guessErrorCode(err error) int {

	if apiError, ok := err.(*errorcode.APIError); ok {
		return apiError.Code
	}

	if err, ok := err.(net.Error); ok && err.Timeout() {
		return errorcode.ErrServiceRespTimeout
	}

	return errorcode.ErrBizLogic
}

func guessErrorMessage(err error) string {
	if err != nil {
		conf.AppLog.WithFields(logrus.Fields{
			"error": err,
		}).Warn()
	}

	if apiError, ok := err.(*errorcode.APIError); ok {
		return apiError.Msg
	}

	if err, ok := err.(net.Error); ok && err.Timeout() {
		return "依赖服务请求超时"
	}

	return errorcode.DefaultErrorMsg
}
