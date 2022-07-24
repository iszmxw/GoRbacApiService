package response

import "gorbac/pkg/utils/times"

type JsonLoginLog struct { // 格式化返回登录日志
	Id        uint64           `json:"id"`
	AccountId uint64           `json:"account_id"`
	Type      int              `json:"type"`
	Username  string           `json:"username"`
	RoleId    int              `json:"role_id"`
	RoleName  string           `json:"role_name"`
	Ip        string           `json:"ip"`
	Address   string           `json:"address"`
	CreatedAt times.TimeNormal `json:"created_at"`
	UpdatedAt times.TimeNormal `json:"updated_at"`
	DeletedAt times.TimeNormal `json:"deleted_at"`
}
