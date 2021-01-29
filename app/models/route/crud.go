package route

import (
	"gorbac/app/models"
	"gorbac/pkg/mysql"
	"gorbac/pkg/utils/logger"
)

// Create 创建菜单路由节点
func (Model Route) Create(a Route) (err error) {
	if err = mysql.DB.Create(&a).Error; err != nil {
		return err
	}
	return nil
}

// GetOne 获取一条数据
func (Model Route) GetOne(where map[string]interface{}) (Route, error) {
	var role Route
	if err := mysql.DB.Where(where).First(&role).Error; err != nil {
		return role, err
	}
	return role, nil
}

// GetPaginate 获取分页数据，返回错误
func (Model Route) GetPaginate(accountId uint64, orderBy interface{}, lists *models.PageList) {
	var result []JsonRoute
	// 获取表名
	tableName := Model.TableName()
	table := mysql.DB.Table(models.Prefix(tableName))
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
