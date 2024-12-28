package controller

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/iamleizz/bluebell/logic"
	"github.com/iamleizz/bluebell/models"
	"go.uber.org/zap"
)

// CreatePostHandler  创建帖子
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
	err = logic.CreatePost(c.Request.Context(), p)

	if err != nil {
		zap.L().Error("login.CreatePost failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return 
	}
	ResponsSuccess(c, nil)
}

// GetPostDetailHandler  获取帖子详细内容
func GetPostDetailHandler(c *gin.Context) {
	// 绑定参数
	pidStr := c.Param(CodeUrlQueryID)
	pid, err := strconv.ParseInt(pidStr, 10, 64)
	if err != nil {
		zap.L().Error("code invalid param", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return 
	}
	
	post, err := logic.GetPostDetail(pid)
	if err != nil {
		zap.L().Error("logic.GetPostDetail Failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return 
	}

	ResponsSuccess(c, post)
}

// GetPostListHandler  获取帖子列表
func GetPostListHandler(c *gin.Context) {
	page, size := getPageInfo(c)
	// 直接走logic
	postlist, err := logic.GetPostList(page, size)
	if err != nil {
		zap.L().Error("logic.GetPostList failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return 
	}
	ResponsSuccess(c, postlist)
}
// GetPostListHandler2  根据时间或热度,对帖子进行排序展示
// 1. 获取分页和排序参数
// 2. 去 redis 中获取对应的post_id查询结果
// 3. 在 mysql 中，根据post_id查询对应的帖子详情
func GetPostListOrderHandler(c *gin.Context) {
	p := &models.ParamPostList{
		Page: 1,
		Size: 10,
		Order: PostOrderByTime, 
	}

	err := c.ShouldBindQuery(p)
	if err != nil {
		zap.L().Error("c.ShouldBindQuery Failed:", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return 
	}

	// 参数绑定完毕，将参数传给 logic 层
	postlist, err := logic.GetPostListOrder(c.Request.Context(), p)

	if err != nil {
		zap.L().Error("logic.GetPostList Failed:", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return 
	}
	ResponsSuccess(c, postlist)
}