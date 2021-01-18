package utils

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

// 多方面获取一个参数
func GetParam(c *gin.Context, name string) string {
	// 从头部获取
	param := c.Request.Header.Get(name)
	if len(param) <= 0 {
		param, _ = c.GetQuery(name)
	}
	// 表单获取
	if len(param) <= 0 {
		param = c.PostForm(name)
	}
	// 提交的 json 中获取
	if len(param) <= 0 {
		var json map[string]interface{}
		_ = c.Bind(&json)
		param = fmt.Sprintf("%s", json[name])
		if "%!s(<nil>)" == param {
			return ""
		}
	}
	return param
}
