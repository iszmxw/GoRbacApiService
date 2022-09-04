package login_log

import (
	"gorbac/app/models"
	"gorbac/app/response"
	"gorbac/pkg/logger"
	"gorbac/pkg/mysql"
)

// Create 创建登录日志，通过 category.ID 来判断是否创建成功
func (m *LoginLog) Create() (err error) {
	if err = mysql.DB.Create(&m).Error; err != nil {
		logger.Error(err)
		return err
	}
	return nil
}

// GetPaginate 获取分页数据，返回错误
func (m *LoginLog) GetPaginate(accountId uint64, orderBy interface{}, lists *models.PageList) {
	var result []response.JsonLoginLog
	// 获取表名
	tableName := m.TableName()
	table := mysql.DB.Table(models.Prefix(tableName))
	table = table.Select(models.Prefix("$prefix_login_log.*,$prefix_role.name as role_name"))
	table = table.Joins(models.Prefix("left join $prefix_role on $prefix_role.id=$prefix_login_log.role_id"))
	table = table.Where(models.Prefix("$prefix_login_log.account_id = ?"), accountId)
	table.Count(&lists.Total)
	// 设置分页参数
	models.InitPageList(lists)
	table = table.Order(orderBy)
	table = table.Offset(int(lists.Offset))
	table = table.Limit(int(lists.PageSize))
	if err := table.Scan(&result).Error; err != nil {
		// 记录错误
		logger.Error(err)
	} else {
		lists.Data = result
	}
}
