package admin

import (
	"github.com/gin-gonic/gin"
	"gorbac/app/models"
	"gorbac/app/models/account"
	"gorbac/app/models/login_log"
	"gorbac/app/models/operation_log"
	"gorbac/pkg/utils"
	"gorbac/pkg/utils/logger"
)

type SystemsController struct {
	BaseController
}

// LoginLogHandler 获取登录日志
func (h *SystemsController) LoginLogHandler(c *gin.Context) {
	logger.LogInfo("SystemsController.LoginLogHandler")
	auth, _ := c.Get("auth")
	// 接收数据使用的结构体
	type PostParams struct {
		Page    int    `json:"page"`
		Limit   int    `json:"limit"`
		OrderBy string `json:"orderBy"`
	}
	var (
		params        PostParams
		loginLogModel login_log.LoginLog
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
		params.OrderBy = "id desc"
	}

	// 申明要取出的数据结构体
	pageList := models.PageList{
		CurrentPage: int64(params.Page),
		PageSize:    int64(params.Limit),
	}

	// 模型获取分页数据
	loginLogModel.GetPaginate(auth.(account.Account).Id, params.OrderBy, &pageList)
	utils.Rjson(c, pageList, "查询成功！")
}

// StatisticsHandler 首页统计
func (h *SystemsController) StatisticsHandler(c *gin.Context) {
	logger.LogInfo("SystemsController.StatisticsHandler")
	//auth, _ := c.Get("auth")
}

// OperationLogHandler 操作日志
func (h *SystemsController) OperationLogHandler(c *gin.Context) {
	logger.LogInfo("SystemsController.OperationLogHandler")
	auth, _ := c.Get("auth")
	// 接收数据使用的结构体
	type PostParams struct {
		Page    int    `json:"page"`
		Limit   int    `json:"limit"`
		OrderBy string `json:"orderBy"`
	}
	var (
		params            PostParams
		operationLogModel operation_log.OperationLog
	)
	// 绑定接收的 json 数据到结构体中
	_ = c.ShouldBindJSON(&params)
	if len(params.OrderBy) <= 0 {
		params.OrderBy = "id desc"
	}
	// 申明要取出的数据结构体
	pageList := models.PageList{
		CurrentPage: int64(params.Page),
		PageSize:    int64(params.Limit),
	}
	// 模型获取分页数据
	operationLogModel.GetPaginate(auth.(account.Account).Id, params.OrderBy, &pageList)
	utils.Rjson(c, pageList, "查询成功！")
}

// ResetPasswordHandler 修改密码
func (SystemsController) ResetPasswordHandler(c *gin.Context) {
	logger.LogInfo("SystemsController.ResetPasswordHandler")
}
