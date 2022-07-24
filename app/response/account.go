package response

import (
	"gorbac/pkg/utils/times"
	"gorm.io/gorm"
)


type JsonAccount struct { // 格式化返回账户信息
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
