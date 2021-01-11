package middlewares

import "github.com/gin-gonic/gin"

type Admin struct {
}

// 定义后台中间件
func (*Admin) Admin(c *gin.Context) {
	c.JSON(200, gin.H{"message": "进入中间件"})

	// 终止
	c.Abort()

	// 继续往下面执行
	c.Next()
}
