package account

import (
	"gorbac/pkg/utils/times"
	"gorm.io/gorm"
)

// json 响应结构体定义，供查询数据引用

// JsonAccount 格式化返回登录日志
type JsonAccount struct {
	Id        uint64         `json:"id"`
	Username  string         `json:"username"`
	Level     int            `json:"level"`
	RoleId    int            `json:"role_id"`
	RoleName  string         `json:"role_name"`
	Mobile    string         `json:"mobile"`
	Status    int            `json:"status"`
	CreatedAt times.MyTime   `json:"created_at"`
	UpdatedAt times.MyTime   `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}
