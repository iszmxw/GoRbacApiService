package operation_log

import (
	"gorbac/app/models"
	"gorbac/pkg/mysql"
	"gorbac/pkg/utils/logger"
)

// Create 创建登录日志，通过 category.ID 来判断是否创建成功
func (rm *OperationLog) Create() (err error) {
	if err = mysql.DB.Create(&rm).Error; err != nil {
		logger.LogError(err)
		return err
	}
	return nil
}

// GetPaginate 获取分页数据，返回错误
func (rm *OperationLog) GetPaginate(accountId uint64, orderBy interface{}, lists *models.PageList) {
	var result []JsonOperationLog
	// 获取表名
	tableName := rm.TableName()
	table := mysql.DB.Debug().Table(models.Prefix(tableName))
	table = table.Select(models.Prefix("{$prefix}account.username,{$prefix}operation_log.*"))
	table = table.Joins(models.Prefix("left join {$prefix}account on {$prefix}account.id={$prefix}operation_log.account_id"))
	table = table.Where(models.Prefix("{$prefix}operation_log.account_id = ?"), accountId)
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
