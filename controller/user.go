package controller

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/iamleizz/bluebell/dao/mysql"
	"github.com/iamleizz/bluebell/logic"
	"github.com/iamleizz/bluebell/models"
	"go.uber.org/zap"
)

// SignUpHandler  处理注册请求
func SignUpHandler(c *gin.Context) {
    // 拦截请求
    // 获取参数结构体
    p := new(models.ParamSignUP)
    // 业务处理
    if err := c.ShouldBindJSON(p); err != nil {
        zap.L().Error("SignUp with invalid param", zap.Error(err))
        // 判断 err 是否为 validator 类型
        errs, ok := err.(validator.ValidationErrors)
        if !ok {
            ResponseError(c, CodeInvalidParam)
            return 
        }
        ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
        return 
    }


    if err := logic.SignUp(p); err != nil {
        if errors.Is(err, mysql.ErrUserExist) {
            ResponseError(c, CodeUserExist)
        }
        ResponseError(c, CodeServerBusy)
        return 
    }

    ResponsSuccess(c, nil)
     
}

// LoginHandler  处理登录请求
func LoginHandler(c *gin.Context) {
    // 绑定登录用户参数
    p := new(models.ParamLogin)
    // 绑定参数
    if err := c.ShouldBindJSON(p); err != nil {
        zap.L().Error("Login with invalid param", zap.Error(err))
        errs, ok := err.(validator.ValidationErrors)
        if !ok {
            ResponseError(c, CodeInvalidParam)
            return 
        }
        ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
        return 
    }

    token, err := logic.Login(p)

    if err != nil {
        if errors.Is(err, mysql.ErrUserNotExist) {
            ResponseError(c, CodeUserNotExist)
            return
        }
        ResponseError(c, CodeInvalidPassword)
        return 
    }

    ResponsSuccess(c, token)
}