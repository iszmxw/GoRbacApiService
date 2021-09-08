package routes

import (
	"github.com/gin-gonic/gin"
	middlewares "gorbac/app/middlewares/admin"
)

// RegisterWebRoutes 注册路由
func RegisterWebRoutes(router *gin.RouterGroup) {
	// 路由分组
	// 后台 模块
	Admin := router.Group("/admin")
	{
		Admin.Use(middlewares.Admin())

		// 登录系统
		Admin.POST("/user/login", AdminGroup.LoginController.LoginHandler)
		// 退出系统
		Admin.POST("/user/logout", AdminGroup.LoginController.LogoutHandler)

		// 获取用户信息
		Admin.POST("/user/info", AdminGroup.UserController.UserInfoHandler)
		// 获取用户列表
		Admin.POST("/user/get_list", AdminGroup.UserController.UserListHandler)
		// 创建用户
		Admin.POST("/user/create", AdminGroup.UserController.UserAddHandler)
		// 删除用户
		Admin.POST("/user/deleted", AdminGroup.UserController.UserDelHandler)

		// 系统控制器
		// 首页统计
		Admin.POST("/dashboard/statistics", AdminGroup.SystemsController.StatisticsHandler)
		// 获取登录日志
		Admin.POST("/dashboard/login_log", AdminGroup.SystemsController.LoginLogHandler)
		// 获取操作日志
		Admin.POST("/dashboard/operation_log", AdminGroup.SystemsController.OperationLogHandler)
		// 修改登录密码
		Admin.POST("/dashboard/reset_password", AdminGroup.SystemsController.ResetPasswordHandler)

		// 角色列表
		Admin.POST("/roles/list", AdminGroup.RolesController.ListHandler)
		// 角色详情
		Admin.POST("/roles/routes", AdminGroup.RolesController.DetailHandler)

		// 菜单路由控制器
		// 创建菜单路由
		Admin.POST("/menus/add", AdminGroup.MenusController.CreatedHandler)
		// 删除菜单路由
		Admin.POST("/menus/delete", AdminGroup.MenusController.DeletedHandler)
		// 获取路由菜单详细信息
		Admin.POST("/menus/detail", AdminGroup.MenusController.DetailHandler)
		// 编辑菜单路由
		Admin.POST("/menus/edit", AdminGroup.MenusController.EditHandler)
		// 获取菜单路由
		Admin.POST("/menus/list", AdminGroup.MenusController.ListHandler)
		// 登录时获取菜单路由
		Admin.POST("/menus/async_routes", AdminGroup.MenusController.AsyncRoutesHandler)

		// 工具控制器
		Admin.POST("/uploads/images", AdminGroup.ToolsController.UploadsHandler)
	}
}
