package admin

import (
	"github.com/gin-gonic/gin"
	"gorbac/app/models"
	"gorbac/app/models/account"
	"gorbac/app/models/role"
	"gorbac/app/models/routes"
	"gorbac/pkg/utils"
	"gorbac/pkg/utils/logger"
)

type RolesController struct {
}

// ListHandler 角色列表
func (h *RolesController) ListHandler(c *gin.Context) {
	logger.LogInfo("RolesController.ListHandler")
	auth, _ := c.Get("auth")
	// 接收数据使用的结构体
	type PostParams struct {
		Page    int    `json:"page"`
		Limit   int    `json:"limit"`
		OrderBy string `json:"orderBy"`
	}
	var (
		params PostParams
		Model  role.Role
	)
	// 绑定接收的 json 数据到结构体中
	_ = c.ShouldBindJSON(&params)
	if params.Page <= 0 {
		params.Page = 1
	}
	if params.Limit <= 0 {
		params.Limit = 10
	}
	if len(params.OrderBy) <= 0 {
		params.OrderBy = "id asc"
	}

	// 申明要取出的数据结构体
	pageList := models.PageList{
		CurrentPage: int64(params.Page),
		PageSize:    int64(params.Limit),
	}
	// 模型获取分页数据
	Model.GetPaginate(auth.(account.Account).Id, params.OrderBy, &pageList)
	utils.Rjson(c, pageList, "查询成功！")
}

// DetailHandler 角色详情
func (h *RolesController) DetailHandler(c *gin.Context) {
	logger.LogInfo("RolesController.DetailHandler")
	auth, _ := c.Get("auth")
	// 接收数据使用的结构体
	type PostParams struct {
		Id int `json:"id"`
	}
	var (
		params       PostParams
		disabled     bool
		Role         role.Role
		Routes       routes.Routes
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
	AllRouteList, _ = Routes.GetRouteList()
	result.DefaultChecked, _ = Role.GetValue(uint64(params.Id))
	result.AllRouteList = routes.GetRoleTree(AllRouteList, 0, disabled)
	utils.Rjson(c, result, "查询成功！")

}

// EditHandler 编辑角色
func (h *RolesController) EditHandler(c *gin.Context) {
	logger.LogInfo("RolesController.EditHandler")
	//auth, _ := c.Get("auth")
}

// CreatedHandler 创建角色
func (h *RolesController) CreatedHandler(c *gin.Context) {
	logger.LogInfo("RolesController.CreatedHandler")

}

// DeletedHandler 删除角色
func (h *RolesController) DeletedHandler(c *gin.Context) {
	logger.LogInfo("RolesController.DeletedHandler")

}
