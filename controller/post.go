package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/iamleizz/bluebell/logic"
	"github.com/iamleizz/bluebell/models"
	"go.uber.org/zap"
)

func CreatePostHandler(c *gin.Context) {
	p := new(models.Post)
	// 绑定参数
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("code invalid params", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return 
	}

	authorID, err := getCurrentUserID(c)
	if err != nil {
		zap.L().Error("用户未登录", zap.Error(err))
		ResponseError(c, CodeNeedLogin)
		return 
	}

	p.AuthorID = authorID
	err = logic.CreatePost(p)

	if err != nil {
		zap.L().Error("login.CreatePost failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return 
	}
	ResponsSuccess(c, nil)
}