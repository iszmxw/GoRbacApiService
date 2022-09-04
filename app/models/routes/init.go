package routes

import (
	"gorbac/app/models"
	"gorbac/pkg/config"
)

func (m *Routes) TableName() string {
	prefix := config.GetString("database.mysql.prefix")
	table := "routes"
	return prefix + table
}

// Routes 菜单路由节点
type Routes struct {
	models.BaseModel
	Sort      int    `gorm:"type:int(11);default:0;comment:排序" json:"sort"`
	Type      string `gorm:"type:varchar(10);not null;default:page;comment:page-页面   api-接口" json:"type"`
	IsMenu    int    `gorm:"type:int(1);not null;default:1;comment:是否根菜单1-是 0-否" json:"is_menu"`
	Route     string `gorm:"type:varchar(255);comment:访问路由地址" json:"route" validate:"route"`
	Component string `gorm:"type:varchar(255);comment:页面组件地址" json:"component" validate:"component"`
	Name      string `gorm:"type:varchar(100);not null;comment:路由名称" json:"name" validate:"name"`
	Icon      string `gorm:"type:varchar(255);comment:icon图标" json:"icon"`
	ParentId  int    `gorm:"type:int(11);not null;default:0;comment:上级id" json:"parent_id" validate:"parent_id"`
	CreateBy  int    `gorm:"type:int(11);comment:创建者" json:"create_by"`
	Status    int    `gorm:"type:char(1);comment:1-已启用   0-未启用" json:"status"`
	models.BaseModelLast
}
