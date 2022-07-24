package role

import (
	"gorbac/app/models"
	"gorbac/app/response"
	"gorbac/pkg/mysql"
	"gorbac/pkg/utils/logger"
	"strconv"
	"strings"
)

// Create 创建角色，通过 Role.ID 来判断是否创建成功
func (m *Role) Create(a Role) (err error) {
	if err = mysql.DB.Create(&a).Error; err != nil {
		return err
	}
	return nil
}

// GetOne 获取一条数据
func (m *Role) GetOne(where map[string]interface{}) (Role, error) {
	var role Role
	if err := mysql.DB.Where(where).First(&role).Error; err != nil {
		return role, err
	}
	return role, nil
}

// GetValue 获取一条角色详情
func (m *Role) GetValue(Id uint64) (interface{}, error) {
	var result []string
	var routes string
	// 获取表名
	tableName := m.TableName()
	table := mysql.DB.Table(models.Prefix(tableName))
	table = table.Where(models.Prefix("$prefix_role.id =" + strconv.FormatUint(Id, 10)))
	table = table.Select(models.Prefix("$prefix_role.routes"))
	table.Scan(&routes)
	result = strings.Split(routes, ",")
	return result, nil
}

// GetPaginate 获取分页数据，返回错误
func (m *Role) GetPaginate(accountId uint64, orderBy interface{}, lists *models.PageList) {
	var result []response.JsonRole
	// 获取表名
	tableName := m.TableName()
	table := mysql.DB.Table(models.Prefix(tableName))
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
