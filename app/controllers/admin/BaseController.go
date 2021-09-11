package admin

type BaseController struct {
}

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
