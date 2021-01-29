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
func (LoginLog LoginLog) GetPaginate(accountId uint64, orderBy interface{}, lists *models.PageList) {
	var result []JsonLoginLog
	// 获取表名
	tableName := LoginLog.TableName()
	table := mysql.DB.Table(models.Prefix(tableName))
	table = table.Select(models.Prefix("$prefix_login_log.*,$prefix_role.name as role_name"))
	table = table.Joins(models.Prefix("left join $prefix_role on $prefix_role.id=$prefix_login_log.role_id"))
	table = table.Where(models.Prefix("$prefix_login_log.account_id = ?"), accountId)
	table.Count(&lists.Total)
	table = table.Order(orderBy)
	table = table.Offset(int(lists.Offset))
	table = table.Limit(int(lists.PageSize))
	if err := table.Scan(&result).Error; err != nil {
		// 记录错误
		logger.LogError(err)
	} else {
		lists.Data = result
	}
	// 设置分页参数
	models.InitPageList(lists)
}
