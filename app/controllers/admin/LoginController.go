package admin

import (
	"github.com/gin-gonic/gin"
	"gorbac/app/models/account"
	"gorbac/app/models/login_log"
	"gorbac/app/validate/admin"
	"gorbac/pkg/jwt"
	"gorbac/pkg/utils"
	"gorbac/pkg/utils/logger"
)

type LoginController struct {
	BaseController
}

// LoginHandler 登录
func (h *LoginController) LoginHandler(c *gin.Context) {
	logger.LogInfo("LoginController.LoginHandler")
	// 初始化数据模型结构体
	params := new(account.Account)
	// 绑定接收的 json 数据到结构体中
	_ = c.Bind(params)
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
func (h *LoginController) LogoutHandler(c *gin.Context) {
	logger.LogInfo("LoginController.LogoutHandler")
	c.JSON(200, gin.H{"code": 20000, "msg": "退出成功"})
}
