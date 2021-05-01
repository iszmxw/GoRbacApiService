package config

import (
	"gorbac/pkg/config"
	"gorbac/pkg/utils/helpers"
)

// 全局常量设计
func init() {
	config.Add("Global", config.StrMap{
		// 全局唯一ID
		"RequestId": helpers.GetUUID(),
	})
}
