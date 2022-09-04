package response

import (
	"gorbac/pkg/utils/helpers"
	"gorm.io/gorm"
)

type Account struct { // 格式化返回账户信息
	Id        uint64             `json:"id"`
	Username  string             `json:"username"`
	Password  string             `json:"password"`
	Level     int                `json:"level"`
	RoleId    int                `json:"role_id"`
	RoleName  string             `json:"role_name"`
	Mobile    string             `json:"mobile"`
	Status    int                `json:"status"`
	CreatedAt helpers.TimeNormal `json:"created_at"`
	UpdatedAt helpers.TimeNormal `json:"updated_at"`
	DeletedAt gorm.DeletedAt     `json:"deleted_at"`
}
