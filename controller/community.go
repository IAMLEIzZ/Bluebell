package controller

import (
	"strconv"

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

func CommunityDetailHandler(c *gin.Context) {
	// 接收路径参数 id
	idStr := c.Param(CodeUrlQueryID)
	communityID, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		zap.L().Error("Code Invalid Param", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return 
	}

	// 查询社区详细内容
	data, err := logic.GetCommunityDetail(communityID)
	if err != nil {
		zap.L().Error("logic.GetCommunityDetail failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return 
	}

	ResponsSuccess(c, data)
}