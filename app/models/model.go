package models

import (
	"gorbac/pkg/config"
	"gorbac/pkg/utils/helpers"
	"gorm.io/gorm"
	"strings"
)

// BaseModel 模型基类
type BaseModel struct {
	Id uint64 `gorm:"column:id;primaryKey;autoIncrement;not null" json:"id"`
}

// BaseModelLast 模型基类
type BaseModelLast struct {
	CreatedAt helpers.TimeNormal `gorm:"column:created_at;index;comment:创建时间" json:"created_at"`
	UpdatedAt helpers.TimeNormal `gorm:"column:updated_at;comment:更新时间" json:"updated_at"`
	// 支持 gorm 软删除
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;index;comment:删除时间" json:"deleted_at"`
}

// GetStringID 获取 ID 的字符串格式
func (m *BaseModel) GetStringID() string {
	return helpers.Uint64ToString(m.Id)
}

// PageList 分页返回数
type PageList struct {
	CurrentPage int64       `json:"current_page"`
	FirstPage   int64       `json:"first_page"`
	LastPage    int64       `json:"last_page"`
	PageSize    int64       `json:"page_size"`
	Total       int64       `json:"total"`
	Data        interface{} `json:"data"`
	Offset      int64       `json:"-"`
}

// InitPageList 设置分页参数
func InitPageList(lists *PageList) {
	// 当前页数
	if lists.CurrentPage <= 0 {
		lists.CurrentPage = 1
	}
	// 每页取出多少条数据，默认15条
	if lists.PageSize == 0 {
		lists.PageSize = 15
	}
	// 第一页
	if lists.FirstPage == 0 {
		lists.FirstPage = 1
	}

	lists.Offset = (lists.CurrentPage - 1) * lists.PageSize
	if lists.Offset < 0 {
		lists.Offset = 0
	}

	// 最后一页
	if (lists.Total % lists.PageSize) != 0 {
		lists.LastPage = (lists.Total / lists.PageSize) + 1
	} else {
		lists.LastPage = lists.Total / lists.PageSize
	}
}

// Prefix 置换sql中的表前缀
func Prefix(str string) string {
	prefix := config.GetString("database.mysql.prefix")
	return strings.Replace(str, "$prefix_", prefix, -1)
}
