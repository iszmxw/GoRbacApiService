package operation_log

import (
	"gorbac/app/models"
	"gorbac/pkg/config"
)

func (m *OperationLog) TableName() string {
	prefix := config.GetString("database.mysql.prefix")
	table := "operation_log"
	return prefix + table
}

// OperationLog 操作日志表
type OperationLog struct {
	models.BaseModel
	Type      int    `gorm:"type:int(2);not null;default:1;comment:类型1：为总后台用户，2：为合作商用户" json:"type"`
	AccountId uint64 `gorm:"type:int(11);not null;comment:账号id" json:"account_id"`
	Username  string `gorm:"-" json:"username"`
	Content   string `gorm:"type:varchar(255);not null;comment:操作内容" json:"content"`
	Ip        string `gorm:"type:varchar(20);not null;comment:操作ip" json:"ip"`
	Address   string `gorm:"type:varchar(20);comment:操作地址" json:"address"`
	models.BaseModelLast
}
