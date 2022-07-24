package response

import (
	"gorm.io/gorm"
	"time"
)

type JsonOperationLog struct { // 格式化返回操作日志
	Id        uint64         `json:"id"`
	Type      int            `json:"type"`
	AccountId uint64         `json:"account_id"`
	Username  string         `json:"username"`
	RoleName  string         `json:"role_name"`
	Content   string         `json:"content"`
	Ip        string         `json:"ip"`
	Address   string         `json:"address"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}
