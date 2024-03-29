package account

import (
	"gorbac/app/models"
	"gorbac/app/response"
	"gorbac/pkg/logger"
	"gorbac/pkg/mysql"
)

// Create 创建账户，通过 User.ID 来判断是否创建成功
func Create(m *Account) (err error) {
	if err = mysql.DB.Create(m).Error; err != nil {
		return err
	}
	return nil
}

// Delete 删除账户
func (m *Account) Delete(where map[string]interface{}) (err error) {
	if err = mysql.DB.Debug().Where(where).Delete(&Account{}).Error; err != nil {
		return err
	}
	return nil
}

// GetOne 获取一条数据
func GetOne(where map[string]interface{}) (response.Account, error) {
	var account response.Account
	if err := mysql.DB.Model(Account{}).
		Where(where).
		Joins(models.Prefix("left join $prefix_role on $prefix_role.id=$prefix_account.role_id")).
		Select(models.Prefix("$prefix_account.*,$prefix_role.name as role_name")).
		First(&account).Error; err != nil {
		return account, err
	}
	return account, nil
}

// GetPaginate 获取分页数据，返回错误
func (m Account) GetPaginate(accountId uint64, orderBy interface{}, lists *models.PageList) {
	result := new([]response.Account)
	// 获取表名
	tableName := m.TableName()
	table := mysql.DB.Debug().Table(models.Prefix(tableName))
	table = table.Select(models.Prefix("$prefix_account.*,$prefix_role.name as role_name"))
	table = table.Joins(models.Prefix("left join $prefix_role on $prefix_role.id=$prefix_account.role_id"))
	if accountId != 1 {
		table = table.Where(models.Prefix("$prefix_account.pid = ?"), accountId)
	}
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
