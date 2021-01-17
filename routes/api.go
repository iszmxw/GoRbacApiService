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
		// 登录系统
		Admin.POST("/user/login", admin.LoginHandler)
		// 退出系统
		Admin.POST("/logout", admin.LogoutHandler)
		// 调用中间件
		Admin.Use(middlewares.Admin)
		// 获取用户信息
		Admin.POST("/userinfo", admin.UserInfoHandler)
	}
}
