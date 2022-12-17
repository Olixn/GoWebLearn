package controller

import (
	"errors"

	"github.com/Olixn/GoWebLearn/dao/mysql"

	"github.com/go-playground/validator/v10"

	"go.uber.org/zap"

	"github.com/Olixn/GoWebLearn/models"

	"github.com/Olixn/GoWebLearn/logic"
	"github.com/gin-gonic/gin"
)

// SignUpHandler 处理注册请求的函数
func SignUpHandler(c *gin.Context) {
	// 1. 获取参数并校验
	p := new(models.ParamSignUp)
	if err := c.ShouldBindJSON(p); err != nil {
		// 请求参数有误，直接返回响应
		zap.L().Error("SignUp with invalid param.", zap.Error(err))
		// 判断err是不是validator.ValidationErrors类型
		if errs, ok := err.(validator.ValidationErrors); ok {
			// 是此类型，对错误进行翻译并返回
			ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
			return
		} else {
			// 不是此类型
			ResponseError(c, CodeInvalidParam)
			return
		}
	}

	// 2. 业务处理
	if err := logic.SignUp(p); err != nil {
		zap.L().Error("logic.SignUp failed.", zap.Error(err))
		if errors.Is(err, mysql.ErrorUserExist) {
			ResponseError(c, CodeUserExist)
			return
		}
		ResponseError(c, CodeServerBusy)
		return
	}

	// 3. 返回响应
	ResponseSuccess(c, nil)
}

// LoginHandler 处理登录请求的函数
func LoginHandler(c *gin.Context) {
	p := new(models.ParamLogin)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("Login with invalid param.", zap.Error(err))
		if errs, ok := err.(validator.ValidationErrors); ok {
			ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
			return
		} else {
			ResponseError(c, CodeInvalidParam)
			return
		}
	}

	if token, err := logic.Login(p); err != nil {
		zap.L().Error("logic.login failed.", zap.String("username", p.Username), zap.Error(err))
		if errors.Is(err, mysql.ErrorUserNotExist) {
			ResponseError(c, CodeUserNotExist)
			return
		} else if errors.Is(err, mysql.ErrorInvalidPassword) {
			ResponseError(c, CodeInvalidPassword)
			return
		}
		ResponseError(c, CodeServerBusy)
		return
	} else {
		ResponseSuccess(c, map[string]interface{}{"token": token})
	}
}
