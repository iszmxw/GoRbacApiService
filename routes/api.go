package routes

import (
	"github.com/gin-gonic/gin"
	"gorbac/app/http/controllers/admin"
	"gorbac/app/http/middlewares"
)

// 注册路由
func RegisterWebRoutes(router *gin.RouterGroup) {
	// 路由分组
	// 后台 模块
	Admin := router.Group("/admin")
	{
		Admin.Use(middlewares.Admin())
		// 登录系统
		Admin.POST("/user/login", admin.LoginHandler)
		// 退出系统
		Admin.POST("/user/logout", admin.LogoutHandler)
		// 获取用户信息
		Admin.POST("/user/info", admin.UserInfoHandler)
	}
}
