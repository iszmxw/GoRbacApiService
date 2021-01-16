package config

import "gorbac/pkg/config"

// jwt参数
func init() {
	config.Add("jwt", config.StrMap{
		// 默认7200秒 => 两个小时的过期时间
		"export": config.Env("JWT_EXPORT", 7200),
		// 默认加密盐
		"secretary": config.Env("JWT_SECRETARY", "iszmxw"),
	})
}
