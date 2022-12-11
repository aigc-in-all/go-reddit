package main

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	"goreddit/dao/mysql"
	"goreddit/dao/redis"
	"goreddit/logger"
	"goreddit/pkg/snowflake"
	"goreddit/route"
	"goreddit/setting"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// @title
// @version 1.0
// @description
// @termsOfService
// @host 127.0.0.1:8080
// @BasePath /api/v1
func main() {
	// 1. 加载配置
	if err := setting.Init(); err != nil {
		fmt.Printf("init config failed. err: %v\n", err)
		return
	}
	// 2. 初始化日志
	if err := logger.Init(setting.Conf.LogConfig, setting.Conf.Mode); err != nil {
		fmt.Printf("init logger failed. err: %v\n", err)
		return
	}
	defer zap.L().Sync()
	zap.L().Debug("logger init success")

	// 3. 初始化Mysql连接
	if err := mysql.Init(setting.Conf.MySQLConfig); err != nil {
		fmt.Printf("init mysql failed. err: %v\n", err)
		return
	}
	defer mysql.Close()

	// 4. 初始化Redis连接
	if err := redis.Init(setting.Conf.RedisConfig); err != nil {
		fmt.Printf("init redis failed. err: %v\n", err)
		return
	}
	defer redis.Close()

	// 初始化ID生成模块
	if err := snowflake.Init(setting.Conf.StartTime, setting.Conf.MachineID); err != nil {
		fmt.Printf("init snowflake failed. err: %v\n", err)
		return
	}

	// 5. 注册路由
	r := route.Setup(setting.Conf.Mode)
	// 6. 启动服务（优雅关机）
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", setting.Conf.Port),
		Handler: r,
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			zap.L().Fatal("listen", zap.Error(err))
		}
	}()

	// 等待中断信号来优雅的关闭服务器，为关闭服务器设置一个5秒的超时
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	zap.L().Info("Shutdown server")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		zap.L().Fatal("Server shutdown", zap.Error(err))
	}
	zap.L().Info("Server exit")
}
