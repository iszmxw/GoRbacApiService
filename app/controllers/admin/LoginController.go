package admin

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorbac/app/models/account"
	"gorbac/app/models/login_log"
	"gorbac/app/validate/admin"
	"gorbac/pkg/jwt"
	"gorbac/pkg/utils"
)

type LoginController struct {
}

// 登录
func (Controller LoginController) LoginHandler(c *gin.Context) {
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

	//查询数据库
	where := make(map[string]interface{})
	where["username"] = params.Username
	res, _ := account.GetOne(where)
	if params.Password != res.Password {
		utils.SuccessErr(c, 50000, "密码输入错误，请您确认后再试！")
		return
	}
	//登录成功，返回token
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
	//返回数据
	utils.SuccessErr(c, 50000, "登录失败，请联系管理员！")
}

// 退出
func (Controller LoginController) LogoutHandler(c *gin.Context) {
	c.JSON(20000, gin.H{"message": "退出系统"})
}

// 获取用户信息
func (Controller LoginController) UserInfoHandler(c *gin.Context) {
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
