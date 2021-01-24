package login_log

import (
	"gorbac/app/models"
	"gorbac/pkg/mysql"
	"gorbac/pkg/utils/logger"
)

// Create 创建登录日志，通过 category.ID 来判断是否创建成功
func (LoginLog *LoginLog) Create() (err error) {
	if err = mysql.DB.Create(&LoginLog).Error; err != nil {
		logger.LogError(err)
		return err
	}

	return nil
}

// GetPaginate 获取分页数据，返回错误
func (LoginLog LoginLog) GetPaginate(where interface{}, orderBy interface{}, lists *models.PageList) {
	// 获取表名
	tableName := LoginLog.TableName()
	table := mysql.DB.Debug().Table(models.Prefix(tableName)).Where(where).Order(orderBy)
	table.Count(&lists.Total)
	// 设置分页参数
	models.InitPageList(lists)
	table = table.Offset(int(lists.Offset)).Limit(int(lists.PageSize)).Find(lists.Data)
	err := table.Error
	if err != nil {
		// 记录错误
		logger.LogError(err)
	}
}
