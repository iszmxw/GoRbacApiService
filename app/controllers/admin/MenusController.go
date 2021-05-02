package admin

import (
	"github.com/gin-gonic/gin"
	"gorbac/app/models/account"
	"gorbac/app/models/routes"
	"gorbac/app/validate/admin"
	"gorbac/pkg/utils"
	"gorbac/pkg/utils/logger"
)

type MenusController struct {
}

// CreatedHandler 创建路由
func (MenusController) CreatedHandler(c *gin.Context) {
	logger.LogInfo("MenusController.CreatedHandler")
	// 登录信息
	auth, _ := c.Get("auth")
	// 初始化数据模型结构体
	params := routes.Routes{}
	// 绑定接收的 json 数据到结构体中
	_ = c.ShouldBindJSON(&params)
	// 验证器验证
	msg := admin.ValidateMenuCreate(params)
	if len(msg) > 0 {
		utils.SuccessErr(c, 50000, msg)
		return
	}
	params.CreateBy = int(auth.(account.Account).Id)
	err := routes.Routes{}.Create(params)
	if err != nil {
		utils.SuccessErr(c, 50000, err.Error())
	}
	utils.Rjson(c, "", "创建成功！")

}

// DeletedHandler 删除路由
func (MenusController) DeletedHandler(c *gin.Context) {
	logger.LogInfo("MenusController.DeletedHandler")
	// 登录信息
	_, _ = c.Get("auth")
	var params routes.Routes
	_ = c.BindJSON(&params)
	err := routes.Routes{}.Delete(params)
	if err != nil {
		utils.SuccessErr(c, 50000, err.Error())
	}
	utils.Rjson(c, "", "删除成功！")
}

// DetailHandler 路由详情
func (MenusController) DetailHandler(c *gin.Context) {
	logger.LogInfo("MenusController.DetailHandler")
	var (
		Route routes.Routes
	)
	// 绑定接收的 json 数据到结构体中
	_ = c.ShouldBindJSON(&Route)
	// 模型获取数据
	_ = Route.GetDetail(map[string]interface{}{"id": Route.Id}, &Route)
	utils.Rjson(c, Route, "查询成功！")
}

// EditHandler 编辑菜单路由
func (MenusController) EditHandler(c *gin.Context) {
	logger.LogInfo("MenusController.EditHandler")
	//auth, _ := c.Get("auth")
	var (
		Route routes.Routes
	)
	// 绑定接收的 json 数据到结构体中
	_ = c.ShouldBindJSON(&Route)

	// TODO 添加数据检验，后面在添加
	// 更新数据
	err := Route.Updates(map[string]interface{}{"id": Route.Id}, &Route)
	if err != nil {
		utils.SuccessErr(c, 5000, "操作失败！"+err.Error())
	}
	utils.Rjson(c, "", "操作成功！")
}

// ListHandler 菜单列表
func (MenusController) ListHandler(c *gin.Context) {
	logger.LogInfo("MenusController.ListHandler")
	// 模型获取分页数据
	result, _ := routes.Routes{}.GetMenuList(map[string]interface{}{}, "sort asc")
	listTree := routes.GetMenuTree(result, 0)
	utils.Rjson(c, listTree, "查询成功！")
}

// AsyncRoutesHandler 登录时获取菜单路由
func (MenusController) AsyncRoutesHandler(c *gin.Context) {
	logger.LogInfo("MenusController.AsyncRoutesHandler")
	// 模型获取分页数据
	result, _ := routes.Routes{}.GetMenuList(map[string]interface{}{"type": "page"}, "sort asc")
	listTree := routes.GetMenuTree(result, 0)
	utils.Rjson(c, listTree, "查询成功！")
}
