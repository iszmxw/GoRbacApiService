package admin

import (
	"github.com/gin-gonic/gin"
	"gorbac/app/models/account"
	"gorbac/app/models/menus"
	"gorbac/app/models/role"
	"gorbac/app/models/route"
	"gorbac/app/validate/admin"
	"gorbac/pkg/utils"
)

type MenusController struct {
}

// 创建路由
func (Controller MenusController) CreatedHandler(c *gin.Context) {
	// 登录信息
	auth, _ := c.Get("auth")
	// 初始化数据模型结构体
	params := menus.Menus{}
	// 绑定接收的 json 数据到结构体中
	_ = c.ShouldBindJSON(&params)
	// 验证器验证
	msg := admin.ValidateMenuCreate(params)
	if len(msg) > 0 {
		utils.SuccessErr(c, 50000, msg)
		return
	}
	params.CreateBy = int(auth.(account.Account).Id)
	err := menus.Menus{}.Create(params)
	if err != nil {
		utils.SuccessErr(c, 50000, err.Error())
	}
	utils.Rjson(c, "", "创建成功！")

}

// 删除路由
func (Controller MenusController) DeletedHandler(c *gin.Context) {
	// 登录信息
	_, _ = c.Get("auth")
	var params menus.Menus
	_ = c.BindJSON(&params)
	err := menus.Menus{}.Delete(params)
	if err != nil {
		utils.SuccessErr(c, 50000, err.Error())
	}
	utils.Rjson(c, "", "删除成功！")
}

// 菜单列表
func (MenusController) ListHandler(c *gin.Context) {
	// 模型获取分页数据
	result, _ := menus.Menus{}.GetList("sort asc")
	listTree := menus.GetTree(result, 0)
	utils.Rjson(c, listTree, "查询成功！")
}

// 路由详情
func (MenusController) DetailHandler(c *gin.Context) {
	auth, _ := c.Get("auth")
	// 接收数据使用的结构体
	type PostParams struct {
		Id int `json:"id"`
	}
	var (
		params       PostParams
		disabled     bool
		Role         role.Role
		Route        route.Route
		result       role.JsonRoleDetail
		AllRouteList []role.AllRouteList
	)
	// 绑定接收的 json 数据到结构体中
	_ = c.ShouldBindJSON(&params)
	if auth.(account.Account).Id == 1 {
		disabled = false
	} else {
		disabled = true
	}
	// 模型获取数据
	AllRouteList, _ = Route.GetTreeData()
	result.DefaultChecked, _ = Role.GetValue(uint64(params.Id))
	result.AllRouteList = route.GetTree(AllRouteList, 0, disabled)
	utils.Rjson(c, result, "查询成功！")

}

// 编辑路由
func (Controller MenusController) EditHandler(c *gin.Context) {
	//auth, _ := c.Get("auth")
}
