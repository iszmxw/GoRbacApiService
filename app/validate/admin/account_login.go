package admin

import (
	"github.com/thedevsaddam/govalidator"
	"gorbac/app/models/account"
)

// ValidateAccount 验证表单，返回 errs 长度等于零即通过
func ValidateAccount(data account.Account) map[string][]string {
	// 1. 定制认证规则
	rules := govalidator.MapData{
		"username": []string{"required", "between:5,20"},
		"password": []string{"required", "min:6"},
	}

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
		Data:          &data,
		Rules:         rules,
		TagIdentifier: "valid", // 模型中的 Struct 标签标识符
		Messages:      messages,
	}

	// 4. 开始验证
	errs := govalidator.New(opts).ValidateStruct()
	return errs
}
