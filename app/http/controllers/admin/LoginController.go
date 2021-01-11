package admin

import "github.com/gin-gonic/gin"

// 定义登录控制器结构体
type LoginController struct {
}

// 登录
func (* LoginController) Login(c *gin.Context){
	c.JSON(200, gin.H{"message":"登录系统"})
}

// 退出
func (* LoginController) Logout(c *gin.Context){
	c.JSON(200, gin.H{"message":"退出系统"})
}

// 获取用户信息
func (* LoginController) UserInfo(c *gin.Context){
	c.JSON(200, gin.H{"message":"获取用户信息"})
}
