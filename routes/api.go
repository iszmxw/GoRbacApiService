package routes

import (
	"github.com/gin-gonic/gin"
	"gorbac/app/controllers/admin"
	admin2 "gorbac/app/middlewares/admin"
)

// 注册路由
func RegisterWebRoutes(router *gin.RouterGroup) {
	// 路由分组
	// 后台 模块
	Admin := router.Group("/admin")
	{
		Admin.Use(admin2.Admin())

		// 登录控制器
		LoginController := admin.LoginController{}
		// 登录系统
		Admin.POST("/user/login", LoginController.LoginHandler)
		// 退出系统
		Admin.POST("/user/logout", LoginController.LogoutHandler)
		// 获取用户信息
		Admin.POST("/user/info", LoginController.UserInfoHandler)

		// 系统控制器
		SystemsController := admin.SystemsController{}
		// 首页统计
		Admin.POST("/dashboard/statistics", SystemsController.StatisticsHandler)
		// 获取登录日志
		Admin.POST("/dashboard/login_log", SystemsController.LoginLogHandler)
		// 获取操作日志
		Admin.POST("/dashboard/operation_log", SystemsController.OperationLogHandler)
		// 修改登录密码
		Admin.POST("/dashboard/reset_password", SystemsController.ResetPasswordHandler)

		// 角色控制器
		RolesController := admin.RolesController{}
		// 角色列表
		Admin.POST("/roles/list", RolesController.ListHandler)
		// 角色详情
		Admin.POST("/roles/routes", RolesController.DetailHandler)
	}
}
