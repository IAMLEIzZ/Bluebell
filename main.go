package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/iamleizz/bulebell/dao/mysql"
	"github.com/iamleizz/bulebell/dao/redis"
	"github.com/iamleizz/bulebell/logger"
	"github.com/iamleizz/bulebell/routes"
	"github.com/iamleizz/bulebell/setting"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func main() {
	//1. 加载配置
	if err := setting.Init(); err != nil {
		fmt.Printf("init setting failed, err:%v\n", err)
		return 
	}
	//2. 初始化日志
	if err := logger.Init(); err != nil {
		fmt.Printf("init logger failed, err:%v\n", err)
		return 
	}
	defer zap.L().Sync()
	zap.L().Debug("logger init success...")
	//3. 初始化数据库
	if err := mysql.Init(); err != nil {
		fmt.Printf("init mysql failed, err:%v\n", err)
		return 
	}
	defer mysql.Close()
	//4. 初始化redis
	if err := redis.Init(); err != nil {
		fmt.Printf("init redis failed, err:%v\n", err)
		return 
	}
	defer redis.Close()
	//5. 注册路由
	r := routes.SetUp()
	//6. 启动服务（优雅启动和重启）
	srv := &http.Server{
		Addr: fmt.Sprintf(":%d", viper.GetInt("app.port")),
		Handler: r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			zap.L().Fatal("listen: ", zap.Error(err))
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	zap.L().Info("Shutdown Server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	
	if err := srv.Shutdown(ctx); err != nil {
		zap.L().Fatal("Server Shutdown", zap.Error(err))
	}

	zap.L().Info("Server exiting")
}