package requests

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

// GetParam 多方面获取一个参数
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
		param = fmt.Sprintf("%v", json[name])
		//fmt.Println(fmt.Sprintf("名称是：%v，值是%v", name, param))
		if "<nil>" == param {
			param = ""
		}
	}
	return param
}
