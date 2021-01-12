package admin

import (
	"github.com/gin-gonic/gin"
)

// 登录
func LoginHandler(c *gin.Context) {
	c.JSON(200, gin.H{"message": "登录系统"})
}

// 退出
func LogoutHandler(c *gin.Context) {
	c.JSON(200, gin.H{"message": "退出系统"})
}

// 获取用户信息
func UserInfoHandler(c *gin.Context) {
	c.JSON(200, gin.H{"message": "获取用户信息"})
}
