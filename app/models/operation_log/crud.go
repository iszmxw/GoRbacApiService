package operation_log

import (
	"gorbac/app/models"
	"gorbac/pkg/mysql"
	"gorbac/pkg/utils/logger"
)

// Create 创建登录日志，通过 category.ID 来判断是否创建成功
func (OperationLog *OperationLog) Create() (err error) {
	if err = mysql.DB.Create(&OperationLog).Error; err != nil {
		logger.LogError(err)
		return err
	}
	return nil
}

// GetPaginate 获取分页数据，返回错误
func (OperationLog OperationLog) GetPaginate(where interface{}, orderBy interface{}, lists *models.PageList) {
	// 设置分页参数
	models.InitPageList(lists)
	if lists.Total > 0 && lists.LastPage >= lists.CurrentPage {
		// 查询语句
		if err := mysql.DB.Debug().Preload("Username").Where(where).Order(orderBy).Offset(int(lists.Offset)).Limit(int(lists.PageSize)).Find(lists.Data).Error; err != nil {
			// 记录错误
			logger.LogError(err)
		}
	}
}
