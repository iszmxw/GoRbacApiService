package admin

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorbac/app/models"
	"gorbac/app/models/account"
	"gorbac/pkg/utils"
)

type UserController struct {
}

// UserInfoHandler 获取用户信息
func (*UserController) UserInfoHandler(c *gin.Context) {
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
func (UserController) UserListHandler(c *gin.Context) {
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