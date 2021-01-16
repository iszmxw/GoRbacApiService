package middlewares

import (
	"github.com/gin-gonic/gin"
	"gorbac/app/models"
	"gorbac/pkg/jwt"
	"gorbac/pkg/utils"
)

// 定义后台中间件
func Admin(c *gin.Context) {
	tokenString := c.Request.Header.Get("Authorization")
	id, err := jwt.AuthToken(tokenString)
	if err != nil {
		utils.SuccessErr(c, 401, err)
		c.Abort()
		return
	}
	//查询数据库
	where := make(map[string]interface{})
	where["id"] = id
	auth := models.GetOne(where, models.Account{})
	// 保存用户到 上下文
	c.Set("auth", auth)
	c.Set("auth_id", id)
	// 继续往下面执行
	c.Next()
}
