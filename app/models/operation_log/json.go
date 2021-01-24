package operation_log

import (
	"gorm.io/gorm"
	"time"
)

// json 响应结构体定义，供查询数据引用

// 格式化返回操作日志
type JsonOperationLog struct {
	Id        uint64         `json:"id"`
	Type      int            `json:"type"`
	AccountId uint64         `json:"account_id"`
	Username  string         `json:"username"`
	Content   string         `json:"content"`
	Ip        string         `json:"ip"`
	Address   string         `json:"address"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}
