package admin

import (
	"github.com/gin-gonic/gin"
	"gorbac/app/models"
	"gorbac/pkg/utils"
)

// 获取登录日志
func LoginLogHandler(c *gin.Context) {
	auth, _ := c.Get("auth")
	// 接收数据使用的结构体
	type PostParams struct {
		Page    int    `json:"page"`
		Limit   int    `json:"limit"`
		OrderBy string `json:"orderBy"`
	}
	var (
		params         PostParams
		loginLogModel  models.LoginLog
		loginLogModels []models.LoginLog
	)
	// 绑定接收的 json 数据到结构体中
	_ = c.ShouldBindJSON(&params)

	if params.Page <= 0 {
		params.Page = 1
	}
	if params.Limit <= 0 {
		params.Page = 10
	}
	if len(params.OrderBy) <= 0 {
		params.OrderBy = "id asc"
	}

	// 申明要取出的数据结构体
	pageList := models.PageList{
		CurrentPage: int64(params.Page),
		PageSize:    int64(params.Limit),
		Data:        &loginLogModels,
	}
	// 查询条件
	where := map[string]interface{}{
		"account_id": auth.(models.Account).Id,
	}
	// 模型获取分页数据
	loginLogModel.GetPaginate(where, params.OrderBy, &pageList)
	utils.Rjson(c, pageList, "查询成功！")
}

// 首页统计
func StatisticsHandler(c *gin.Context) {
	//auth, _ := c.Get("auth")
}

// 操作日志
func OperationLogHandler(c *gin.Context) {
	auth, _ := c.Get("auth")
	// 接收数据使用的结构体
	type PostParams struct {
		Page    int    `json:"page"`
		Limit   int    `json:"limit"`
		OrderBy string `json:"orderBy"`
	}
	var (
		params             PostParams
		operationLogModel  models.OperationLog
		operationLogModels []models.OperationLog
	)
	// 绑定接收的 json 数据到结构体中
	_ = c.ShouldBindJSON(&params)

	if len(params.OrderBy) <= 0 {
		params.OrderBy = "id asc"
	}

	// 申明要取出的数据结构体
	pageList := models.PageList{
		CurrentPage: int64(params.Page),
		PageSize:    int64(params.Limit),
		Data:        &operationLogModels,
	}
	// 查询条件
	where := map[string]interface{}{
		"account_id": auth.(models.Account).Id,
	}
	// 模型获取分页数据
	operationLogModel.GetPaginate(where, params.OrderBy, &pageList)
	utils.Rjson(c, pageList, "查询成功！")
}

// 修改密码
func ResetPasswordHandler(c *gin.Context) {

}
