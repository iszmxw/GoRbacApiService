package models

import (
	"gorbac/pkg/utils/types"
	"gorm.io/gorm"
	"time"
)

// BaseModel 模型基类
type BaseModel struct {
	Id uint64 `gorm:"column:id;primaryKey;autoIncrement;not null" json:"id"`
}

// BaseModelLast 模型基类
type BaseModelLast struct {
	CreatedAt time.Time `gorm:"column:created_at;index;comment:创建时间" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;comment:更新时间" json:"updated_at"`
	// 支持 gorm 软删除
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;index;comment:删除时间" json:"deleted_at"`
}

// GetStringID 获取 ID 的字符串格式
func (a BaseModel) GetStringID() string {
	return types.Uint64ToString(a.Id)
}

// 分页返回数
type PageList struct {
	CurrentPage int64       `json:"current_page"`
	FirstPage   int64       `json:"first_page"`
	LastPage    int64       `json:"last_page"`
	PageSize    int64       `json:"page_size"`
	Total       int64       `json:"total"`
	Data        interface{} `json:"data"`
}
