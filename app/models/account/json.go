package account

import (
	"gorm.io/gorm"
	"time"
)

// json 响应结构体定义，供查询数据引用

// JsonAccount 格式化返回登录日志
type JsonAccount struct {
	Id        uint64         `json:"id"`
	AccountId uint64         `json:"account_id"`
	Type      int            `json:"type"`
	Username  string         `json:"username"`
	RoleId    int            `json:"role_id"`
	RoleName  string         `json:"role_name"`
	Ip        string         `json:"ip"`
	Address   string         `json:"address"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}
