package route

import (
	"gorbac/app/models"
	"gorbac/app/models/role"
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
	var route Route
	if err := mysql.DB.Where(where).First(&route).Error; err != nil {
		return route, err
	}
	return route, nil
}

func (Model Route) GetTreeData() ([]role.AllRouteList, error) {
	var result []role.AllRouteList
	tableName := Model.TableName()
	table := mysql.DB.Table(models.Prefix(tableName))
	table = table.Select(models.Prefix("id,is_menu,name,parent_id"))
	if err := table.Scan(&result).Error; err != nil {
		// 记录错误
		logger.LogError(err)
		return result, err
	}
	return result, nil
}

// 递归生成菜单结构
func GetTree(data []role.AllRouteList, parentId int, disabled bool) []role.AllRouteList {
	var listTree []role.AllRouteList
	for _, val := range data {
		val.Disabled = disabled
		if val.ParentId == parentId {
			children := GetTree(data, int(val.Id), disabled)
			if len(children) > 0 {
				val.Children = children
			}
			listTree = append(listTree, val)
		}
	}
	return listTree
}

// GetPaginate 获取分页数据，返回错误
func (Model Route) GetPaginate(_ uint64, orderBy interface{}, lists *models.PageList) {
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
