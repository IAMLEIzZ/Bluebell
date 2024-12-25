package routes

import (

	"github.com/gin-gonic/gin"
	"github.com/iamleizz/bluebell/controller"
	"github.com/iamleizz/bluebell/logger"
	"github.com/iamleizz/bluebell/middlewares"
)

func SetUp(mode string) *gin.Engine{
	// 注册路由
	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	v1 := r.Group("/api/v1")

	// 注册业务请求
	{
		v1.POST("/signup", controller.SignUpHandler)
		v1.POST("/login", controller.LoginHandler)
	}
	v1.Use(middlewares.JWTAuthMiddleware())
	// 社区业务
	{
		v1.GET("/community", controller.CommunityListHandler)
		v1.GET("/community/:id", controller.CommunityDetailHandler)
	}
	// 帖子请求
	{
		v1.POST("/post", controller.CreatePostHandler)
		v1.GET("/post/:id", controller.GetPostDetailHandler)
	}
	

	return r
}