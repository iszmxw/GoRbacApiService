package utils

import (
	"github.com/gin-gonic/gin"
)

// Rjson 成功返回封装 参数 data interface{} 类型为可接受任意类型
func Rjson(c *gin.Context, data interface{}, msg string) {
	var rdata map[string]interface{}
	if len(msg) > 0 {
		rdata = gin.H{"code": 20000, "data": data, "msg": msg}
	} else {
		rdata = gin.H{"code": 20000, "data": data}
	}
	c.JSON(200, rdata)
}

// SuccessErr 错误返回封装
func SuccessErr(c *gin.Context, errCode int, msg interface{}) {
	c.JSON(200, gin.H{
		"code": errCode,
		"msg":  msg,
	})
}
