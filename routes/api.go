package routes

import (
	"github.com/gin-gonic/gin"
	"gorbac/app/http/controllers/admin"
	"gorbac/app/http/middlewares"
)

// 注册路由
func RegisterWebRoutes() {
	// 初始化路由
	router := gin.Default()
	// 路由分组
	v1 := router.Group("/v1")
	{
		// 后台 模块
		Admin := v1.Group("/admin")
		{
			// 登录系统
			Admin.POST("/login", admin.LoginHandler)
			// 退出系统
			Admin.POST("/logout", admin.LogoutHandler)
			// 调用中间件
			Admin.Use(middlewares.Admin)
			// 获取用户信息
			Admin.POST("/userinfo", admin.UserInfoHandler)
		}
	}

	// 启动路由
	router.Run(":80")

}
