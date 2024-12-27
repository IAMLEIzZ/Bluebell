package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/iamleizz/bluebell/logic"
	"github.com/iamleizz/bluebell/models"
	"go.uber.org/zap"
)

/*
	帖子投票功能梳理
	1. 接收参数，参数包含 投票人 ID，帖子 ID，以及投票方向
	2. 接受参数后，交由 redis 处理，分别存储三个 zset
		{
			1. PostID:Time	// 按照时间排序的 zset
			2. PostID:Score		// 按照分数排序的 zset
			3. PostId:		// 某篇帖子投票的用户 zset
		}
*/
func PostVoteHandler(c *gin.Context) {
	p := new(models.ParamVote)
	// 参数绑定成功
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Debug("c.ShouldBindJSON Failed", zap.Error(err))
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return 
		}
		ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
        return 
	}
	// 调用 logic 层
	userID, err := getCurrentUserID(c)
	if err != nil {
		ResponseError(c, CodeNeedLogin)
		return
	}
	if err = logic.PostVote(c.Request.Context(), userID, p); err != nil {
		zap.L().Error("logic.PostVote failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return 
	}
	ResponsSuccess(c, nil)
}