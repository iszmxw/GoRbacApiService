package admin

import (
	"github.com/gin-gonic/gin"
	"gorbac/app/models"
	"gorbac/app/validate/admin"
	"gorbac/pkg/utils"
)

// 登录
func LoginHandler(c *gin.Context) {
	// 初始化数据模型结构体
	params := models.Account{}
	// 绑定接收的 json 数据到结构体中
	_ = c.ShouldBindJSON(&params)
	// 验证器验证
	msg := admin.ValidateAccount(params)
	if len(msg) > 0 {
		utils.SuccessErr(c, 500, msg)
		return
	}

	//查询数据库
	where := make(map[string]interface{})
	where["username"] = params.Username
	res := models.GetOne(where, models.Account{})
	if params.Password != res.Password {
		utils.SuccessErr(c, 500, "密码输入错误，请您确认后再试！")
		return
	}
	//登录成功，返回token

	//返回数据
	utils.SuccessData(c, "用户名为："+params.Username+"，输入的密码为："+params.Password)
}

// 退出
func LogoutHandler(c *gin.Context) {
	c.JSON(200, gin.H{"message": "退出系统"})
}

// 获取用户信息
func UserInfoHandler(c *gin.Context) {
	c.JSON(200, gin.H{"message": "获取用户信息"})
}
