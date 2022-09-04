package admin

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"gorbac/app/requests"
	"gorbac/app/response"
	"gorbac/pkg/echo"
	"gorbac/pkg/redis"
)

// Admin 定义后台中间件
func Admin() gin.HandlerFunc {
	return func(c *gin.Context) {
		path := c.FullPath()
		switch path {
		case "/v1/admin/user/login":
		case "/v1/admin/user/logout":
			// 继续往下面执行
			c.Next()
			break
		default:
			CheckLogin(c)
			break

		}
	}
}

// CheckLogin 检测登录
func CheckLogin(c *gin.Context) {
	var account response.Account
	// 获取 "Admin-Token"
	tokenString := requests.GetParam(c, "Admin-Token")
	if len(tokenString) <= 0 || tokenString == "<nil>" {
		echo.Error(c, "FAIL", "未检测到token")
		c.Abort()
		return
	}
	str, _ := redis.Client.Get(tokenString)
	if len(str) == 0 {
		echo.Error(c, "FAIL", "token无效，登录失败，请重新登录后再试")
		c.Abort()
		return
	}
	err := json.Unmarshal([]byte(str), &account)
	if err != nil {
		echo.Error(c, "FAIL", "token无效，登录失败，请重新登录后再试")
		c.Abort()
		return
	}
	// 保存用户到 上下文
	c.Set("auth", account)
	c.Set("auth_id", account.Id)
	// 继续往下面执行
	c.Next()
}
