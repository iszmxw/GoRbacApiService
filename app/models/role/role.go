package role

import (
	"gorbac/app/models"
)

// Role 角色表
type Role struct {
	models.BaseModel
	Name   string `gorm:"type:varchar(100);not null;comment:角色名称" json:"name"`
	Routes string `gorm:"type:varchar(255);comment:路由id,该角色所具有的路由" json:"routes"`
	Desc   string `gorm:"type:varchar(255);comment:角色描述" json:"desc"`
	models.BaseModelLast
}
