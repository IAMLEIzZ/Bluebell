package controller

import (
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
)


var ErrorUserNotLogin = errors.New("用户未登录")

func getCurrentUserID(c *gin.Context) (userID int64, err error) {
	uid, ok := c.Get(ContextUserIDKey)
	if !ok {
		err = ErrorUserNotLogin
		return
	}

	userID, ok = uid.(int64)
	if !ok {
		err = ErrorUserNotLogin
		return
	}
	return
}

// getPageInfo  获取分页参数
func getPageInfo(c *gin.Context) (int64, int64) {
	pageStr := c.Query(CodeUrlQueryPage)
	sizeStr := c.Query(CodeUrlQuerySize)

	var (
		page int64
		size int64
		err  error
	)
	page, err = strconv.ParseInt(pageStr, 10, 64)
	if err != nil {
		page = 1
	}
	size, err = strconv.ParseInt(sizeStr, 10, 64)
	if err != nil {
		size = 10
	}

	return page, size
}
