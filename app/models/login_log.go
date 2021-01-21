package models

import (
	"gorbac/pkg/mysql"
	"gorbac/pkg/utils/logger"
)

// LoginLog 登录日志表
type LoginLog struct {
	BaseModel
	AccountId uint64 `gorm:"type:int(11);NOT NULL;COMMENT:账号id" json:"account_id"`
	Type      int    `gorm:"type:int(2);NOT NULL;COMMENT:类型1：为总后台用户，2：为合作商用户" json:"type"`
	Username  string `gorm:"type:varchar(50);not null;comment:用户名" json:"username"`
	RoleId    int    `gorm:"type:int(11);not null;comment:角色id" json:"role_id"`
	Ip        string `gorm:"type:varchar(20);NOT NULL;COMMENT:登录ip" json:"ip"`
	Address   string `gorm:"type:varchar(20);comment:登录地址" json:"address"`
	BaseModelLast
}

// Create 创建登录日志，通过 category.ID 来判断是否创建成功
func (LoginLog *LoginLog) Create() (err error) {
	if err = mysql.DB.Create(&LoginLog).Error; err != nil {
		logger.LogError(err)
		return err
	}

	return nil
}

// GetPaginate 获取分页数据，返回错误
func (LoginLog LoginLog) GetPaginate(where interface{}, orderBy interface{}, lists *PageList) {
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
