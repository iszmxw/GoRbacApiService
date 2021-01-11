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
		PrefixAdmin := v1.Group("/admin")
		{
			LoginController := new(admin.LoginController)
			// 登录系统
			PrefixAdmin.POST("/login", LoginController.Login)
			// 退出系统
			PrefixAdmin.POST("/logout", LoginController.Logout)
			AdminMiddleware := new(middlewares.Admin)
			// 获取用户信息
			PrefixAdmin.Use(AdminMiddleware.Admin).POST("/userinfo", LoginController.UserInfo)
		}
	}

	// 启动路由
	router.Run(":80")

}
