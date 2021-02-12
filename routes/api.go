package routes

import (
	"github.com/gin-gonic/gin"
	"gorbac/app/controllers/admin"
	middlewares "gorbac/app/middlewares/admin"
)

// 注册路由
func RegisterWebRoutes(router *gin.RouterGroup) {
	// 路由分组
	// 后台 模块
	Admin := router.Group("/admin")
	{
		Admin.Use(middlewares.Admin())

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

		// 菜单路由控制器
		MenusController := admin.MenusController{}
		// 创建菜单路由
		Admin.POST("/menus/add", MenusController.CreatedHandler)
		// 删除菜单路由
		Admin.POST("/menus/delete", MenusController.DeletedHandler)
		// 获取路由菜单详细信息
		Admin.POST("/menus/detail", MenusController.DetailHandler)
		// 编辑菜单路由
		Admin.POST("/menus/edit", MenusController.EditHandler)
		// 获取菜单路由
		Admin.POST("/menus/list", MenusController.ListHandler)
		// 登录时获取菜单路由
		Admin.POST("/menus/async_routes", MenusController.AsyncRoutesHandler)
	}
}
