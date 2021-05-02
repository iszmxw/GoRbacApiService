package admin

import (
	"github.com/gin-gonic/gin"
	"gorbac/app/models/account"
	"gorbac/pkg/jwt"
	"gorbac/pkg/utils"
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
	// 获取 "Admin-Token"
	tokenString := utils.GetParam(c, "Admin-Token")
	if len(tokenString) <= 0 || tokenString == "<nil>" {
		utils.SuccessErr(c, 50000, "未检测到token")
		c.Abort()
		return
	}
	id, err := jwt.AuthToken(tokenString)
	if err != nil {
		utils.SuccessErr(c, 50000, err)
		c.Abort()
		return
	}
	//查询数据库
	where := make(map[string]interface{})
	where["id"] = id
	auth, _ := account.GetOne(where)
	// 保存用户到 上下文
	c.Set("auth", auth)
	c.Set("auth_id", id)
	// 继续往下面执行
	c.Next()
}
