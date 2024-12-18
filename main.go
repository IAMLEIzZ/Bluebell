package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/iamleizz/bluebell/controller"
	"github.com/iamleizz/bluebell/dao/mysql"
	"github.com/iamleizz/bluebell/dao/redis"
	"github.com/iamleizz/bluebell/logger"
	"github.com/iamleizz/bluebell/pkg/snowflake"
	"github.com/iamleizz/bluebell/routes"
	"github.com/iamleizz/bluebell/setting"
	"go.uber.org/zap"
)

func main() {
	//1. 加载配置
	if err := setting.Init(); err != nil {
		fmt.Printf("init setting failed, err:%v\n", err)
		return 
	}
	//2. 初始化日志
	if err := logger.Init(setting.Conf.LogConfig, setting.Conf.Mode); err != nil {
		fmt.Printf("init logger failed, err:%v\n", err)
		return 
	}
	defer zap.L().Sync()
	zap.L().Debug("logger init success...")
	//3. 初始化数据库
	if err := mysql.Init(setting.Conf.MySQLConfig); err != nil {
		fmt.Printf("init mysql failed, err:%v\n", err)
		return 
	}
	defer mysql.Close()
	//4. 初始化redis
	if err := redis.Init(setting.Conf.RedisConfig); err != nil {
		fmt.Printf("init redis failed, err:%v\n", err)
		return 
	}
	defer redis.Close()
	// 初始化 gin 框架内置的校验器的翻译器
	if err := controller.InitTrans("zh"); err != nil {
		fmt.Printf("init validator trans faild, err:%v\n", err)
		return 
	}
	//5. 注册路由
	r := routes.SetUp(setting.Conf.Mode)
	//7. 初始化分布式 id 生成器
	if err := snowflake.Init("2024-12-12", 1); err != nil {
		fmt.Printf("init snowflake faild, err:%v\n", err)
		return 
	}
	//8. 启动服务（优雅启动和重启）
	port := fmt.Sprintf("%s%d%s", "++++++++-> PORT:", setting.Conf.Port, " <-++++++++")
	fmt.Println(port)
	srv := &http.Server{
		Addr: fmt.Sprintf(":%d", setting.Conf.Port),
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