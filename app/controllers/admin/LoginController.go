package admin

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorbac/app/models"
	"gorbac/app/models/account"
	"gorbac/app/models/login_log"
	"gorbac/app/validate/admin"
	"gorbac/pkg/jwt"
	"gorbac/pkg/utils"
	"gorbac/pkg/utils/logger"
)

type LoginController struct {
}

// LoginHandler 登录
func (LoginController) LoginHandler(c *gin.Context) {
	logger.LogInfo("LoginController.LoginHandler")
	// 初始化数据模型结构体
	params := account.Account{}
	// 绑定接收的 json 数据到结构体中
	_ = c.ShouldBindJSON(&params)
	// 验证器验证
	msg := admin.ValidateAccount(params)
	if len(msg) > 0 {
		utils.SuccessErr(c, 50000, msg)
		return
	}

	// 查询数据库
	where := make(map[string]interface{})
	where["username"] = params.Username
	res, _ := account.GetOne(where)
	if params.Password != res.Password {
		utils.SuccessErr(c, 50000, "密码输入错误，请您确认后再试！")
		return
	}
	// 登录成功，返回token
	t, err := jwt.GetToken(res)
	if err != nil {
		utils.SuccessErr(c, 50000, err)
		return
	}

	// 添加登录日志
	loginLog := login_log.LoginLog{
		AccountId: res.Id,
		Type:      res.Level,
		Username:  res.Username,
		RoleId:    res.RoleId,
		Ip:        "127.0.0.1",
		Address:   "本地开发",
	}
	_ = loginLog.Create()
	if loginLog.Id > 0 {
		//返回数据
		utils.Rjson(c, gin.H{
			"token": t["token"],
		}, "登录成功！")
		return
	}
	// 返回数据
	utils.SuccessErr(c, 50000, "登录失败，请联系管理员！")
}

// LogoutHandler 退出
func (LoginController) LogoutHandler(c *gin.Context) {
	logger.LogInfo("LoginController.LogoutHandler")
	c.JSON(200, gin.H{"code": 20000, "msg": "退出成功"})
}

// UserInfoHandler 获取用户信息
func (LoginController) UserInfoHandler(c *gin.Context) {
	auth, _ := c.Get("auth")
	c.JSON(200, gin.H{
		"message": "登录成功",
		"code":    20000,
		"data": gin.H{
			"roles":        fmt.Sprintf("[%s]", "admin"),
			"name":         auth.(account.Account).Username,
			"introduction": auth.(account.Account).Username,
			"avatar":       "http://blog.54zm.com/style/web/iszmxw_simple_pro/static/images/head.jpg",
		},
	})
}

// UserListHandler 获取用户列表
func (LoginController) UserListHandler(c *gin.Context) {
	auth, _ := c.Get("auth")
	auth_id, _ := c.Get("auth_id")
	auth_id = auth_id.(uint64)

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

	fmt.Println(fmt.Sprintf("%v",params))

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
