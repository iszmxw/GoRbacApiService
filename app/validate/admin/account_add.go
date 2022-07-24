package admin

import (
	"github.com/thedevsaddam/govalidator"
	"gorbac/app/models/account"
	"gorbac/pkg/utils/logger"
)

// ValidateAddAccount 验证表单，开始验证，并返回一条错误消息
func ValidateAddAccount(account *account.Account) string {
	var msg string
	// 1. 定制认证规则
	rules := govalidator.MapData{
		"username": []string{"required", "between:5,20"},
		"password": []string{"required", "min:6"},
	}
	logger.Info(account)

	// 2. 定制错误消息
	messages := govalidator.MapData{
		"username": []string{
			"required:用户名为必填项",
			"between:格式错误，长度必须为5到20为",
		},
		"password": []string{
			"required:密码为必填项",
			"min:长度需大于 6",
		},
	}

	// 3. 配置初始化
	opts := govalidator.Options{
		Data:          account,
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
