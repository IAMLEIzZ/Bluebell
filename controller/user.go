package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/iamleizz/bluebell/logic"
	"github.com/iamleizz/bluebell/models"
	"go.uber.org/zap"
)

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
            c.JSON(http.StatusOK, gin.H{
            "msg": err.Error(),
            })
            return 
        }
        c.JSON(http.StatusOK, gin.H{
            "msg": removeTopStruct(errs.Translate(trans)),
        })
        return 
    }


    if err := logic.SignUp(p); err != nil {
        c.JSON(http.StatusOK, gin.H{
            "msg": "注册失败",
        })
        return 
    }

    c.JSON(http.StatusOK, gin.H{
        "mag": "success",
    })
     
}