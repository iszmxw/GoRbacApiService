package routes

import (
	"gorbac/app/models"
	"gorbac/app/response"
	"gorbac/pkg/mysql"
	"gorbac/pkg/utils/logger"
)

// Create 创建菜单路由节点
func (m *Routes) Create(a *Routes) (err error) {
	if err = mysql.DB.Create(a).Error; err != nil {
		return err
	}
	return nil
}

// Delete 删除菜单路由节点
func (m *Routes) Delete(a Routes) (err error) {
	if err = mysql.DB.Debug().Delete(&a).Error; err != nil {
		return err
	}
	return nil
}

// GetDetail 获取一条数据详情
func (m *Routes) GetDetail(where map[string]interface{}, data *Routes) error {
	if err := mysql.DB.Debug().Where(where).First(&data).Error; err != nil {
		return err
	}
	return nil
}

// Updates 更新数据
func (m *Routes) Updates(where map[string]interface{}, data map[string]interface{}) error {
	if err := mysql.DB.Debug().Model(m).Where(where).Updates(data).Error; err != nil {
		return err
	}
	return nil
}

// GetRouteList 获取角色路由数据列表
func (m *Routes) GetRouteList() ([]response.AllRouteList, error) {
	var result []response.AllRouteList
	tableName := m.TableName()
	table := mysql.DB.Table(models.Prefix(tableName))
	table = table.Select(models.Prefix("id,is_menu,name,parent_id"))
	if err := table.Scan(&result).Error; err != nil {
		// 记录错误
		logger.Error(err)
		return result, err
	}
	return result, nil
}

// GetMenuList 获取列表数据，返回错误
func (m *Routes) GetMenuList(where interface{}, orderBy interface{}) ([]response.JsonRouteTree, error) {
	var result []response.JsonRouteTree
	// 获取表名
	tableName := m.TableName()
	table := mysql.DB.Table(models.Prefix(tableName))
	table = table.Where("deleted_at IS NULL") // 排除软删除记录
	table = table.Where(where)
	table = table.Order(orderBy)
	if err := table.Scan(&result).Error; err != nil {
		// 记录错误
		logger.Error(err)
		return nil, err
	}
	return result, nil
}

// GetRoleTree 递归生成角色路由菜单结构
func GetRoleTree(data []response.AllRouteList, parentId int, disabled bool) []response.AllRouteList {
	var listTree []response.AllRouteList
	for _, val := range data {
		val.Disabled = disabled
		if val.ParentId == parentId {
			children := GetRoleTree(data, int(val.Id), disabled)
			if len(children) > 0 {
				val.Children = children
			}
			listTree = append(listTree, val)
		}
	}
	return listTree
}

// GetMenuTree 递归生成菜单结构
func GetMenuTree(data []response.JsonRouteTree, parentId int) []response.JsonRouteTree {
	var listTree []response.JsonRouteTree
	for _, val := range data {
		if val.ParentId == parentId {
			children := GetMenuTree(data, int(val.Id))
			if len(children) > 0 {
				val.Children = children
			}
			listTree = append(listTree, val)
		}
	}
	return listTree
}

// GetPaginate 获取分页数据，返回错误
func (m *Routes) GetPaginate(_ uint64, orderBy interface{}, lists *models.PageList) {
	var result []response.JsonRoute
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
