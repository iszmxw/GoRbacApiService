package admin

import (
	"github.com/gin-gonic/gin"
	"gorbac/app/models/account"
	"gorbac/app/validate/admin"
	"gorbac/utils"
)

// 登录
func LoginHandler(c *gin.Context) {
	// 初始化数据模型结构体
	dataAccount := account.Model{}
	// 绑定接收的 json 数据到结构体重
	_ = c.ShouldBindJSON(&dataAccount)
	// 验证器验证
	errs := admin.ValidateAccount(dataAccount)
	// 遍历返回错误
	if len(errs) > 0 {
		for index, item := range errs {
			utils.SuccessErr(c, 500, index+" "+item[0])
			return
		}
	}
	// 返回数据
	utils.SuccessData(c, "用户名为："+dataAccount.Username+"，输入的密码为："+dataAccount.Password)
}

// 退出
func LogoutHandler(c *gin.Context) {
	c.JSON(200, gin.H{"message": "退出系统"})
}

// 获取用户信息
func UserInfoHandler(c *gin.Context) {
	c.JSON(200, gin.H{"message": "获取用户信息"})
}
