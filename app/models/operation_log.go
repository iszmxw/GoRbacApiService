package models

import (
	"gorbac/pkg/mysql"
	"gorbac/pkg/utils/logger"
)

// OperationLog 操作日志表
type OperationLog struct {
	BaseModel
	Type      int    `gorm:"type:int(2);not null;default:1;comment:类型1：为总后台用户，2：为合作商用户" json:"type"`
	AccountId uint64 `gorm:"type:int(11);not null;comment:账号id" json:"account_id"`
	Content   string `gorm:"type:varchar(255);not null;comment:操作内容" json:"content"`
	Ip        string `gorm:"type:varchar(20);not null;comment:操作ip" json:"ip"`
	Address   string `gorm:"type:varchar(20);comment:操作地址" json:"address"`
	BaseModelLast
}

// Create 创建登录日志，通过 category.ID 来判断是否创建成功
func (OperationLog *OperationLog) Create() (err error) {
	if err = mysql.DB.Create(&OperationLog).Error; err != nil {
		logger.LogError(err)
		return err
	}
	return nil
}

// GetPaginate 获取分页数据，返回错误
func (OperationLog OperationLog) GetPaginate(where interface{}, orderBy interface{}, lists *PageList) {
	// 设置分页参数
	InitPageList(lists)
	if lists.Total > 0 && lists.LastPage >= lists.CurrentPage {
		// 查询语句
		if err := mysql.DB.Where(where).Order(orderBy).Offset(int(lists.Offset)).Limit(int(lists.PageSize)).Find(lists.Data).Error; err != nil {
			// 记录错误
			logger.LogError(err)
		}
	}
}
