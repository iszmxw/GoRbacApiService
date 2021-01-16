package main

import (
	"gorbac/app/models"
	"gorbac/routes"
)

func main() {
	// 初始化 SQL
	models.SetupDB()
	// 注册路由服务
	routes.RegisterWebRoutes()
}
