package menus

import (
	"gorbac/app/models"
	"gorbac/pkg/mysql"
	"gorbac/pkg/utils/logger"
)

// Create 创建菜单路由节点
func (Model Menus) Create(a Menus) (err error) {
	if err = mysql.DB.Create(&a).Error; err != nil {
		return err
	}
	return nil
}

// Delete 删除菜单路由节点
func (Model Menus) Delete(a Menus) (err error) {
	if err = mysql.DB.Debug().Delete(&a).Error; err != nil {
		return err
	}
	return nil
}

// GetOne 获取一条数据
func (Model Menus) GetOne(where map[string]interface{}) (Menus, error) {
	var route Menus
	if err := mysql.DB.Where(where).First(&route).Error; err != nil {
		return route, err
	}
	return route, nil
}

// 递归生成菜单结构
func GetTree(data []JsonTreeMenus, parentId int) []JsonTreeMenus {
	var listTree []JsonTreeMenus
	for _, val := range data {
		if val.ParentId == parentId {
			children := GetTree(data, int(val.Id))
			if len(children) > 0 {
				val.Children = children
			}
			listTree = append(listTree, val)
		}
	}
	return listTree
}

// GetList 获取列表数据，返回错误
func (Model Menus) GetList(orderBy interface{}) ([]JsonTreeMenus, error) {
	var result []JsonTreeMenus
	// 获取表名
	tableName := Model.TableName()
	table := mysql.DB.Table(models.Prefix(tableName))
	table = table.Where("deleted_at IS NULL")	// 排除软删除记录
	table = table.Where("type = ?", "page")
	table = table.Order(orderBy)
	if err := table.Scan(&result).Error; err != nil {
		// 记录错误
		logger.LogError(err)
		return nil, err
	}
	return result, nil
}
