package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/iamleizz/bluebell/logic"
	"go.uber.org/zap"
)

func CommunityListHandler(c *gin.Context) {
	// 接收请求
	// 返回查询数据
	data, err := logic.GetCommunityList()
	if err != nil {
		zap.L().Error("logic.GetCommunityList failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return 
	}

	ResponsSuccess(c, data)
}