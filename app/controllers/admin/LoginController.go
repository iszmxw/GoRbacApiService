package admin

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorbac/app/models/account"
	"gorbac/app/models/login_log"
	"gorbac/app/requests"
	"gorbac/app/validate/admin"
	"gorbac/pkg/echo"
	"gorbac/pkg/logger"
	"gorbac/pkg/redis"
)

type LoginController struct {
	BaseController
}

// LoginHandler 登录
func (h *LoginController) LoginHandler(c *gin.Context) {
	// 初始化数据模型结构体
	params := new(requests.Login)
	// 绑定接收的 json 数据到结构体中
	_ = c.Bind(params)
	// 验证器验证
	msg := admin.ValidateAccount(params)
	if len(msg) > 0 {
		echo.Error(c, "FAIL", msg)
		return
	}
	// 查询数据库
	where := make(map[string]interface{})
	where["username"] = params.Username
	res, _ := account.GetOne(where)
	logger.Info(res)
	if params.Password != res.Password {
		echo.Error(c, "FAIL", "密码输入错误，请您确认后再试！")
		return
	}
	// 登录成功，返回token
	token := uuid.New().String()
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
	userInfo, err := json.Marshal(res)
	if err != nil {
		echo.Error(c, "FAIL", err.Error())
		return
	}
	_, err = redis.Client.Add(token, userInfo, 7200) // 两个小时后失效
	if err != nil {
		echo.Error(c, "FAIL", err.Error())
		return
	}
	if loginLog.Id > 0 {
		//返回数据
		echo.Success(c, gin.H{
			"token": token,
		}, "登录成功！")
		return
	}
	// 返回数据
	echo.Error(c, "FAIL", "登录失败，请联系管理员！")
}

// LogoutHandler 退出
func (h *LoginController) LogoutHandler(c *gin.Context) {
	echo.Success(c, nil, "退出成功！")
}
