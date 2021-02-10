package admin

import (
	"github.com/thedevsaddam/govalidator"
	"gorbac/app/models/menus"
)

// ValidateMenuCreate 验证表单，开始验证，并返回一条错误消息
func ValidateMenuCreate(menu menus.Menus) string {
	var msg string
	// 1. 定制认证规则
	rules := govalidator.MapData{
		"name":      []string{"required"},
		"route":     []string{"required"},
	}

	// 2. 定制错误消息
	messages := govalidator.MapData{
		"name": []string{
			"required:菜单名称为必填项",
		},
		"route": []string{
			"required:请填写路由地址",
		},
	}

	// 3. 配置初始化
	opts := govalidator.Options{
		Data:          &menu,
		Rules:         rules,
		Messages:      messages,
		TagIdentifier: "validate", // 模型中的 Struct 标签标识符
	}

	// 4. 开始验证
	errors := govalidator.New(opts).ValidateStruct()

	for key, val := range errors {
		if len(key) > 0 {
			msg += key
			for _, item := range val {
				if len(item) > 0 {
					msg += " " + item
					return msg
				}
			}
		}
	}
	return msg
}
