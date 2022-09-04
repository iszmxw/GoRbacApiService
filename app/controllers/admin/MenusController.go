package admin

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorbac/app/models/routes"
	"gorbac/app/response"
	"gorbac/app/validate/admin"
	"gorbac/pkg/echo"
	"gorbac/pkg/logger"
)

type MenusController struct {
	BaseController
}

// CreatedHandler 创建路由
func (h *MenusController) CreatedHandler(c *gin.Context) {
	logger.Info("MenusController.CreatedHandler")
	// 登录信息
	auth, _ := c.Get("auth")
	// 初始化数据模型结构体
	params := routes.Routes{}
	// 绑定接收的 json 数据到结构体中
	_ = c.ShouldBindJSON(&params)
	// 验证器验证
	msg := admin.ValidateMenuCreate(params)
	if len(msg) > 0 {
		echo.Error(c, "FAIL", msg)
		return
	}
	params.CreateBy = int(auth.(response.Account).Id)
	err := new(routes.Routes).Create(&params)
	if err != nil {
		echo.Error(c, "FAIL", err.Error())
		return
	}
	echo.Success(c, nil, "创建成功！")

}

// DeletedHandler 删除路由
func (h *MenusController) DeletedHandler(c *gin.Context) {
	logger.Info("MenusController.DeletedHandler")
	// 登录信息
	_, _ = c.Get("auth")
	var params routes.Routes
	_ = c.BindJSON(&params)
	err := new(routes.Routes).Delete(params)
	if err != nil {
		echo.Error(c, "FAIL", err.Error())
		return
	}
	echo.Success(c, nil, "删除成功！")
}

// DetailHandler 路由详情
func (h *MenusController) DetailHandler(c *gin.Context) {
	logger.Info("MenusController.DetailHandler")
	var (
		Route routes.Routes
	)
	// 绑定接收的 json 数据到结构体中
	_ = c.ShouldBindJSON(&Route)
	// 模型获取数据
	_ = Route.GetDetail(map[string]interface{}{"id": Route.Id}, &Route)
	echo.Success(c, Route, "查询成功！")
}

// EditHandler 编辑菜单路由
func (h *MenusController) EditHandler(c *gin.Context) {
	logger.Info("MenusController.EditHandler")
	//auth, _ := c.Get("auth")
	var (
		Route routes.Routes
	)
	data := make(map[string]interface{}) //注意该结构接受的内容
	// 绑定接收的 json 数据到结构体中
	_ = c.BindJSON(&data)
	// TODO 添加数据检验，后面在添加
	err := Route.Updates(map[string]interface{}{"id": fmt.Sprintf("%+v", data["id"])}, data)
	if err != nil {
		echo.Error(c, "FAIL", err.Error())
		return
	}
	echo.Success(c, nil, "操作成功！")
}

// ListHandler 菜单列表
func (h *MenusController) ListHandler(c *gin.Context) {
	logger.Info("MenusController.ListHandler")
	// 模型获取分页数据
	result, _ := new(routes.Routes).GetMenuList(map[string]interface{}{}, "sort asc")
	listTree := routes.GetMenuTree(result, 0)
	echo.Success(c, listTree, "查询成功！")
}

// AsyncRoutesHandler 登录时获取菜单路由
func (h *MenusController) AsyncRoutesHandler(c *gin.Context) {
	logger.Info("MenusController.AsyncRoutesHandler")
	// 模型获取分页数据
	result, _ := new(routes.Routes).GetMenuList(map[string]interface{}{"type": "page"}, "sort asc")
	listTree := routes.GetMenuTree(result, 0)
	echo.Success(c, listTree, "查询成功！")
}
