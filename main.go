package main

import (
	"fmt"
	"gorbac/bootstrap"
	"gorbac/config"
	conf "gorbac/pkg/config"
	"gorbac/pkg/utils/logger"
)

func init() {
	// 初始化配置信息
	config.Initialize()
	// 定义日志目录
	logger.Init("GoRbacApiService")
}

func main() {
	// 初始化 SQL
	bootstrap.SetupDB()
	// 初始化路由绑定
	router := bootstrap.SetupRoute()
	// 启动路由
	_ = router.Run(fmt.Sprintf(":%s", conf.GetString("app.port")))
}
