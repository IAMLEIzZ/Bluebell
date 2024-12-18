package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/iamleizz/bluebell/controller"
	"github.com/iamleizz/bluebell/logger"
)

func SetUp(mode string) *gin.Engine{
	// 注册路由
	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))
	// 注册业务请求
	r.POST("/signup", controller.SignUpHandler)
	
	r.GET("/", func (c *gin.Context){
		c.String(http.StatusOK, "hello world")
	})

	return r
}