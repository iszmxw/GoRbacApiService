package main

import (
	"fmt"
	"gorbac/bootstrap"
	"gorbac/config"
	conf "gorbac/pkg/config"
)

func init() {
	// 初始化配置信息
	config.Initialize()
}

func main() {
	// 初始化 SQL
	bootstrap.SetupDB()
	// 初始化路由绑定
	router := bootstrap.SetupRoute()
	fmt.Println(111)
	fmt.Println(conf.GetString("app.port"))
	fmt.Println(222)
	// 启动路由
	router.Run(fmt.Sprintf(":%s", conf.GetString("app.port")))
}
