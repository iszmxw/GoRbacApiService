package login_log

import (
	"gorbac/app/models"
	"gorbac/pkg/config"
)

func (m *LoginLog) TableName() string {
	prefix := config.GetString("database.mysql.prefix")
	table := "login_log"
	return prefix + table
}

type LoginLog struct { // 登录日志表
	models.BaseModel
	AccountId uint64 `gorm:"type:int(11);NOT NULL;COMMENT:账号id" json:"account_id"`
	Type      int    `gorm:"type:int(2);NOT NULL;COMMENT:类型1：为总后台用户，2：为合作商用户" json:"type"`
	Username  string `gorm:"type:varchar(50);not null;comment:用户名" json:"username"`
	RoleId    int    `gorm:"type:int(11);not null;comment:角色id" json:"role_id"`
	Ip        string `gorm:"type:varchar(20);NOT NULL;COMMENT:登录ip" json:"ip"`
	Address   string `gorm:"type:varchar(20);comment:登录地址" json:"address"`
	models.BaseModelLast
}
