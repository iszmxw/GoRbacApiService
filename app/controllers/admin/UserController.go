package admin

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorbac/app/models"
	"gorbac/app/models/account"
	"gorbac/pkg/utils"
	"gorbac/pkg/utils/logger"
)

type UserController struct {
	BaseController
}

// UserInfoHandler 获取用户信息
func (h *UserController) UserInfoHandler(c *gin.Context) {
	auth, _ := c.Get("auth")
	c.JSON(200, gin.H{
		"msg":  "登录成功",
		"code": 20000,
		"data": gin.H{
			"roles":        fmt.Sprintf("[%s]", "admin"),
			"name":         auth.(account.Account).Username,
			"introduction": auth.(account.Account).Username,
			"avatar":       "https://blog.54zm.com/style/web/iszmxw_simple_pro/static/images/head.jpg",
		},
	})
}

// UserListHandler 获取用户列表
func (h *UserController) UserListHandler(c *gin.Context) {
	auth, _ := c.Get("auth")
	authId, _ := c.Get("auth_id")
	authId = authId.(uint64)

	// 接收数据使用的结构体
	type PostParams struct {
		Page    int    `json:"page"`
		Limit   int    `json:"limit"`
		OrderBy string `json:"orderBy"`
	}
	var (
		params       PostParams
		accountModel account.Account
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

	fmt.Println(fmt.Sprintf("%v", params))

	// 申明要取出的数据结构体
	pageList := models.PageList{
		CurrentPage: int64(params.Page),
		PageSize:    int64(params.Limit),
	}
	if auth.(account.Account).Id == 1 {
		// 模型获取分页数据
		accountModel.GetPaginate(auth.(account.Account).Id, params.OrderBy, &pageList)
		utils.Rjson(c, pageList, "查询成功！")
	}
}

// UserAddHandler 添加用户
func (h *UserController) UserAddHandler(c *gin.Context) {
	reqData := new(account.Account)
	// todo 数据验证

	if err := c.Bind(reqData); err != nil {
		logger.LogInfo(reqData)
		utils.SuccessErr(c, 500, "参数格式有误")
		return
	}
	if err := account.Create(reqData); err != nil {
		logger.LogInfo(reqData)
		utils.SuccessErr(c, 500, "创建失败")
		return
	}
	utils.Rjson(c, reqData, "ok")
}

// UserDelHandler 删除用户
func (h *UserController) UserDelHandler(c *gin.Context) {
	reqData := make(map[string]interface{})
	// todo 数据验证
	if err := c.Bind(&reqData); err != nil {
		logger.LogInfo(reqData)
		utils.SuccessErr(c, 500, "参数格式有误")
		return
	}
	if len(reqData["id"].(string)) < 0 {
		utils.SuccessErr(c, 500, "请确认要删除的账号id")
	}
	model := account.Account{}
	err := model.Delete(reqData)
	if err != nil {
		utils.Rjson(c, reqData["id"], "删除失败")
		return
	}
	utils.Rjson(c, reqData["id"], "ok")
}
