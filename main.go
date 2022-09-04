package main

import (
	"flag"
	"fmt"
	"gorbac/bootstrap"
	"gorbac/config"
	conf "gorbac/pkg/config"
	"gorbac/pkg/logger"

	"github.com/gin-contrib/pprof"
)

func init() {
	// 初始化配置信息
	config.Initialize()
	// 定义日志目录
	logger.Init("web")
}

func main() {
	AppPort := flag.Int64("APP_PORT", conf.GetInt64("app.port"), "服务端口")
	flag.Parse()
	logger.Info("初始化 SQL")
	bootstrap.SetupDB()
	logger.Info("初始化 Redis")
	bootstrap.SetupRedis(conf.GetInt("redis.db"))
	defer bootstrap.RedisClose()
	logger.Info("加载 web 路由")
	router := bootstrap.SetupRoute()
	logger.Info("开启 pprof")
	pprof.Register(router)
	logger.Info("启动路由")
	logger.Info(fmt.Sprintf("当前端口:%v", *AppPort))
	_ = router.Run(fmt.Sprintf(":%v", *AppPort))
}
