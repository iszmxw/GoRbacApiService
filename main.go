package main

import (
	"fmt"
	"gorbac/bootstrap"
	"gorbac/config"
	conf "gorbac/pkg/config"
	"gorbac/pkg/utils/helpers"
)

func init() {
	// 初始化配置信息
	config.Initialize()
}

func main() {
	logId := helpers.GetUUID()
	// 初始化 SQL
	bootstrap.SetupDB()
	// 初始化路由绑定
	router := bootstrap.SetupRoute(logId)
	// 启动路由
	_ = router.Run(fmt.Sprintf(":%s", conf.GetString("app.port")))
}
