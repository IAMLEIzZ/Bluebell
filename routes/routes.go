package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/iamleizz/bluebell/logger"
)

func SetUp() *gin.Engine{
	// 注册路由
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	r.GET("/", func (c *gin.Context){
		c.String(http.StatusOK, "hello world")
	})

	return r
}