package admin

type BaseController struct {
}

// AdminGroup Admin 模块控制器
type AdminGroup struct {
	// 注册控制器
	BaseController
	LoginController
	MenusController
	RolesController
	SystemsController
	ToolsController
	UserController
}
