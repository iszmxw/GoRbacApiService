package helpers

// 全局函数助手

import (
	"github.com/google/uuid"
)

// GetUUID 获取UUID
func GetUUID() string {
	return uuid.New().String()
}
